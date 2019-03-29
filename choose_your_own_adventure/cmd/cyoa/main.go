package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyoa "github.com/samueldaviddelacruz/golang-exercises/choose_your_own_adventure"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with the Choose your own adventure story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	file, err := os.Open(*fileName)

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	var story cyoa.Story

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
