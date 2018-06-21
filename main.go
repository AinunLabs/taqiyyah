package main

import (
	"fmt"

	"github.com/taqiyyah/amaliah/reader"
)

func main() {
	dhuha := reader.ReadFromFile("contents/40-doa-shalat-dhuha.json")

	fmt.Println(dhuha)
}
