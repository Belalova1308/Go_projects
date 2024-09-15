package types

type Place struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
