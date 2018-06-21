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

// AmaliahFile struct
type AmaliahFile struct {
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

func getContentFiles(baseDir string) map[int]AmaliahFile {
	results := make(map[int]AmaliahFile)

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

		results[num] = AmaliahFile{
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
