package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/r3labs/diff"
)

type Ingredient struct {
	IngredientName  string  `json:"ingredient_name" xml:"itemname" diff:"ingedient, identifier"`
	IngredientCount float64 `json:"ingredient_count" xml:"itemcount" diff:"item_count"`
	IngredientUnit  string  `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty" diff:"unit"`
}
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item" diff:"ingredients, indentifier"`
}
type Recipe struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type DBReader interface {
	Read(data []byte) (Recipe, error)
}
type JSONReader Recipe

func (x *JSONReader) Read(data []byte) (Recipe, error) {
	error := json.Unmarshal(data, x)
	if error != nil {
		os.Exit(1)
	}
	return Recipe(*x), error
}

type XMLReader Recipe

func (x *XMLReader) Read(data []byte) (Recipe, error) {
	error := xml.Unmarshal(data, x)
	if error != nil {
		os.Exit(1)
	}
	return Recipe(*x), error
}

func Remake(extension string, data []byte) Recipe {
	var reader DBReader
	if extension == ".json" {
		reader = new(JSONReader)
	} else if extension == ".xml" {
		reader = new(XMLReader)
	} else {
		os.Exit(5)
	}
	result, err := reader.Read(data)
	if err != nil {
		os.Exit(6)
	}
	return result
}

func processChange(changelog []diff.Change, oldRecipe, newRecipe Recipe) {
	for _, change := range changelog {
		path := change.Path

		if len(path) >= 2 {
			if cakeIndex, err := convertToInt(path[1]); err == nil {
				oldCakeName := oldRecipe.Cake[cakeIndex].Name
				newCakeName := newRecipe.Cake[cakeIndex].Name
				if path[2] == "Time" {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%v\" instead of \"%v\"\n", oldCakeName, change.To, change.From)
				}
				if len(path) > 3 && path[2] == "ingredients" {
					if ingredientIndex, err := convertToInt(path[3]); err == nil {
						field := fmt.Sprintf("%v", path[4])
						switch change.Type {
						case "create":
							fmt.Printf("ADDED ingredient \"%v\" for cake \"%s\"\n", change.To, newCakeName)
						case "delete":
							fmt.Printf("REMOVED ingredient \"%v\" for cake \"%s\"\n", change.From, oldCakeName)
						case "update":
							switch field {
							case "ingedient":
								fmt.Printf("CHANGED ingredient name from \"%v\" to \"%v\" for cake \"%s\"\n", change.From, change.To, oldCakeName)
							case "item_count":
								fmt.Printf("CHANGED unit count for ingredient \"%v\" for cake \"%s\" - \"%v\" instead of \"%v\"\n", oldRecipe.Cake[cakeIndex].Ingredients[ingredientIndex].IngredientName, oldCakeName, change.To, change.From)
							case "unit":
								fmt.Printf("CHANGED unit for ingredient \"%v\" for cake \"%s\" - \"%v\" instead of \"%v\"\n", oldRecipe.Cake[cakeIndex].Ingredients[ingredientIndex].IngredientName, oldCakeName, change.To, change.From)
							}
						}
					}
				}
			}
		}
	}
}
func convertToInt(val interface{}) (int, error) {
	switch v := val.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", val)
	}
}

func main() {
	flags := flag.NewFlagSet("flags", flag.ExitOnError)
	oldFile := flags.String("old", "", "path to old")
	newFile := flags.String("new", "", "path to new")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	err := flags.Parse(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}
	infoOld, err := os.Open(*oldFile)
	if err != nil {
		os.Exit(3)
	}
	defer infoOld.Close()
	infoNew, err := os.Open(*newFile)
	if err != nil {
		os.Exit(4)
	}
	defer infoNew.Close()
	byteValueOld, err := io.ReadAll(infoOld)
	if err != nil {
		os.Exit(4)
	}
	byteValueNew, err := io.ReadAll(infoNew)
	if err != nil {
		os.Exit(4)
	}
	extOld := path.Ext(*oldFile)
	extNew := path.Ext(*newFile)
	oldRecipe := Remake(extOld, byteValueOld)
	newRecipe := Remake(extNew, byteValueNew)
	changelog, _ := diff.Diff(oldRecipe, newRecipe)
	processChange(changelog, oldRecipe, newRecipe)
}
