package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// AmaliahItem struct
type AmaliahItem struct {
	Title    []Content `json:"title"`
	Contents []Content `json:"contents"`
	Number   int       `json:"number"`
	Type     int       `json:"type"`
	Source   string    `json:"source"`
}

// Content struct
type Content struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

// ReadFromFile to read source from json file
func ReadFromFile(s string) AmaliahItem {
	var result AmaliahItem
	// open json file
	jsonFile, err := os.Open(s)

	if err != nil {
		fmt.Println(err.Error())
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &result)

	return result
}
