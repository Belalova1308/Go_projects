// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"text/template"
	"tidy/db"
	"tidy/types"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang-jwt/jwt/v5"
)

const limit = 10

var secretKey = []byte("your-secret-key")

type TokenResponse struct {
	Token string `json:"token"`
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	store := db.NewElasticsearchStore(es)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlePlaces(w, r, store)
	})

	http.HandleFunc("/api/places", func(w http.ResponseWriter, r *http.Request) {
		handlePlacesAPI(w, r, store)
	})

	http.HandleFunc("/api/recommend", JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleRecommendationsAPI(w, r, store)
	}))

	http.HandleFunc("/api/get_token", HandleGetToken)
	log.Println("Server is running on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func getPlacesData(r *http.Request, store db.Store) ([]types.Place, int, int, int, error) {
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit
	places, total, err := store.GetPlaces(limit, offset)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	lastPage := (total + limit - 1) / limit

	return places, total, page, lastPage, nil
}

func handlePlaces(w http.ResponseWriter, r *http.Request, store db.Store) {
	places, total, page, lastPage, err := getPlacesData(r, store)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	data := struct {
		Total    int
		Places   []types.Place
		Page     int
		LastPage int
		PrevPage int
		NextPage int
	}{
		Total:    total,
		Places:   places,
		Page:     page,
		LastPage: lastPage,
		PrevPage: page - 1,
		NextPage: page + 1,
	}

	tmpl := template.Must(template.ParseFiles("templates/places.html"))
	tmpl.Execute(w, data)
}

func handlePlacesAPI(w http.ResponseWriter, r *http.Request, store db.Store) {
	w.Header().Set("Content-Type", "application/json")

	places, total, page, lastPage, err := getPlacesData(r, store)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	response := struct {
		Total    int           `json:"total"`
		Page     int           `json:"page"`
		LastPage int           `json:"last_page"`
		Places   []types.Place `json:"places"`
	}{
		Total:    total,
		Page:     page,
		LastPage: lastPage,
		Places:   places,
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func handleRecommendationsAPI(w http.ResponseWriter, r *http.Request, store db.Store) {
	latParam := r.URL.Query().Get("lat")
	lonParam := r.URL.Query().Get("lon")

	if latParam == "" || lonParam == "" {
		http.Error(w, "Missing 'lat' or 'lon' parameter", http.StatusBadRequest)
		return
	}

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'lat' value: '%s'", latParam), http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonParam, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'lon' value: '%s'", lonParam), http.StatusBadRequest)
		return
	}

	places, _, err := store.GetPlaces(limit, 0)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}
	type PlaceWithDistance struct {
		Place    types.Place
		Distance float64
	}

	var placesWithDistance []PlaceWithDistance
	for _, p := range places {
		latFloat, _ := strconv.ParseFloat(p.Latitude, 64)
		lonFloat, _ := strconv.ParseFloat(p.Longitude, 64)
		distance := calculateDistance(lat, lon, latFloat, lonFloat)
		placesWithDistance = append(placesWithDistance, PlaceWithDistance{
			Place:    p,
			Distance: distance,
		})
	}

	sort.Slice(placesWithDistance, func(i, j int) bool {
		return placesWithDistance[i].Distance < placesWithDistance[j].Distance
	})
	if len(placesWithDistance) > 3 {
		placesWithDistance = placesWithDistance[:3]
	}
	response := struct {
		Name   string        `json:"name"`
		Places []types.Place `json:"places"`
	}{
		Name:   "Recommendation",
		Places: places,
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return distance
}

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": username,
		"iat":  time.Now().Unix(),
	})
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}
func HandleGetToken(w http.ResponseWriter, r *http.Request) {
	username := "user"
	token, err := createToken(username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: token})
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
