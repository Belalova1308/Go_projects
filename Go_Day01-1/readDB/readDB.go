package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

type Ingredient struct {
	IngredientName  string  `json:"ingredient_name" xml:"itemname"`
	IngredientCount float64 `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string  `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
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
func Remake(x DBReader, data []byte, extension string) []byte {
	var res []byte
	if extension == ".json" {
		x = new(JSONReader)
	} else if extension == ".xml" {
		x = new(XMLReader)
	} else {
		os.Exit(5)
	}
	cakeData, error := x.Read(data)
	if error != nil {
		os.Exit(6)
	}
	if extension == ".json" {
		res, error = xml.MarshalIndent(cakeData, "", "    ")
		if error != nil {
			os.Exit(6)
		}
	} else if extension == ".xml" {
		res, error = json.MarshalIndent(cakeData, "", "    ")
		if error != nil {
			os.Exit(6)
		}
	} else {
		os.Exit(5)
	}
	return res
}
func main() {
	fileName := flag.String("f", "", "Path to the file")
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	infoFile, err := os.Open(*fileName)
	if err != nil {
		os.Exit(3)
	}
	byteValue, err := io.ReadAll(infoFile)
	if err != nil {
		os.Exit(4)
	}
	var reader DBReader
	extension := path.Ext(*fileName)
	var output []byte
	output = Remake(reader, byteValue, extension)
	fmt.Println(string(output))
	defer infoFile.Close()
}
