package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/samueldaviddelacruz/golang-exercises/choose_your_own_adventure"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the choose your own adventure web application on")
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

	handler := cyoa.NewHandler(story)

	fmt.Printf("Starting the server at port: %d\n", *port)
	address := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(address, handler))

}
