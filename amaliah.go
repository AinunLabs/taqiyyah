package amaliah

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/taqiyyah/amaliah/resources"
)

const contentsDir string = "contents"

var contents map[int]AmaliahContent

// AmaliahContent struct
type AmaliahContent struct {
	Number   int
	Filename string
	Path     string
}

// Amaliah struct
type Amaliah struct {
	Titles   []Content `json:"titles"`
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

func init() {
	contents = make(map[int]AmaliahContent)
	data, err := resources.AssetDir(contentsDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, filename := range data {
		fSlice := strings.Split(filename, "-")
		num, _ := strconv.Atoi(fSlice[0])
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
		var amaliah Amaliah
		byteValue, _ := resources.Asset(content.Path)
		json.Unmarshal(byteValue, &amaliah)

		results = append(results, amaliah)
	}

	return results
}
