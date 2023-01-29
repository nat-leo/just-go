package main

import (
	"fmt"
	"log"

	"natleo/greetings"
)

func main() {
	// set the log entry prefix
	log.SetPrefix("greetings: ")
	// disable printing time, line number, source file
	log.SetFlags(0)

	message, err := greetings.Hello("your cousin Throckmorton")
	//message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// list of names
	names := []string{"Gladys", "Willy", "Cucumber Boy"}
	messages, error := greetings.Hellos(names)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}

