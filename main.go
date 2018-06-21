package main

import (
	"fmt"
	"os"
)

func main() {
	// open json file
	jsonFile, err := os.Open("contents/40-doa-shalat-dhuha.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened json file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
