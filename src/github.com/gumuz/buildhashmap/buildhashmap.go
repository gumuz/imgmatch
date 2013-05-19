package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gumuz/imghash"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	hashmap := map[string]uint64{}

	sourceDir := os.Args[1]
	destFile := os.Args[2]

	fmt.Print("Loading images from ", sourceDir, "...")

	// scan source directory for png/jpg image files
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		fmt.Println("Error while reading files from source directory:", err)
		os.Exit(1)
	}
	fmt.Println(len(files), "files found")

	fmt.Print("Building hashmap")
	for idx, file := range files {
		path := sourceDir + file.Name()

		if !strings.Contains(path, ".png") && !strings.Contains(path, ".jpg") {
			continue
		}

		hash, _ := imghash.MeanHash(path)
		hashmap[path] = hash

		// fmt.Println(path, hash)

		// progres bar
		if idx%(len(files)/100) == 0 {
			fmt.Print(".")
		}
	}

	fmt.Println("Saving to", destFile)
	file, _ := os.Create(destFile)
	encoder := gob.NewEncoder(file)
	encoder.Encode(hashmap)
}
