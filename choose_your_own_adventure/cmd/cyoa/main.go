package main

import (
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
	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
