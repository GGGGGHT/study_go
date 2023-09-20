package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(1)
	name := []string{"Glayds", "Smantha", "Darrin"}
	message, err := greetings.Hellos(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
