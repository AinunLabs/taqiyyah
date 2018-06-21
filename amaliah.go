package amaliah

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const contentsDir string = "contents"

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

func (amaliah Amaliah) String() string {
	bytes, err := json.Marshal(amaliah)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(bytes)
}

// All method to get all amaliah content
func All() []Amaliah {
	var results []Amaliah
	contentProperties := getContentFiles(contentsDir)

	for _, contentProperty := range contentProperties {
		amaliah := readFromFile(contentProperty.Path)
		results = append(results, amaliah)
	}

	return results
}

func getContentFiles(baseDir string) map[int]AmaliahContent {
	results := make(map[int]AmaliahContent)

	f, err := os.Open(fmt.Sprintf("./%s", baseDir))
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		sSlice := strings.Split(file.Name(), "-")
		num, _ := strconv.Atoi(sSlice[0])

		results[num] = AmaliahContent{
			Number:   num,
			Filename: file.Name(),
			Path:     fmt.Sprintf("%s/%s", baseDir, file.Name()),
		}
	}

	return results
}

func readFromFile(filepath string) Amaliah {
	var result Amaliah
	// open json file
	jsonFile, err := os.Open(filepath)

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
