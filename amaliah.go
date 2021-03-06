package amaliah

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/taqiyyah/amaliah/resources"
)

const contentsDir string = "contents"

var contents map[float64]AmaliahContent

// AmaliahContent struct
type AmaliahContent struct {
	Number   float64
	Filename string
	Path     string
}

// Amaliah struct
type Amaliah struct {
	Titles   []Content `json:"titles"`
	Contents []Content `json:"contents"`
	Number   float64   `json:"number"`
	Type     int       `json:"type"`
	Source   string    `json:"source"`
}

// Content struct
type Content struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

// Data method on AmaliahContent types
func (content AmaliahContent) Data() Amaliah {
	var result Amaliah
	byteValue, _ := resources.Asset(content.Path)

	json.Unmarshal(byteValue, &result)

	return result
}

// Make json string with Stringers method
func (amaliah Amaliah) String() string {
	bytes, err := json.Marshal(amaliah)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(bytes)
}

func init() {
	contents = make(map[float64]AmaliahContent)
	data, err := resources.AssetDir(contentsDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, filename := range data {
		fSlice := strings.Split(filename, "-")
		num, _ := strconv.ParseFloat(fSlice[0], 64)
		path := fmt.Sprintf("%s/%s", contentsDir, filename)

		contents[num] = AmaliahContent{
			Number:   num,
			Filename: filename,
			Path:     path,
		}
	}
}

// All method to get all amaliah content
func All() []Amaliah {
	var results []Amaliah

	for _, content := range contents {
		results = append(results, content.Data())
	}

	return results
}

// Get one amaliah content by number
func Get(num float64) Amaliah {
	result := contents[num]

	return result.Data()
}
