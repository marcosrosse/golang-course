package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  int    `json:"age"`
}

func main() {
	dog1 := dog{"Crystal", "Poodle", 6}

	dog1JSON, err := json.Marshal(dog1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewBuffer(dog1JSON)) // you need to convert the byte slice to a string with the bytes.NewBuffer() function

	// with map
	dog2 := map[string]string{
		"name": "Artemis",
		"race": "Pit bull",
	}
	dog2JSON, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewBuffer(dog2JSON)) // you need to convert the byte slice to a string with the bytes.NewBuffer() function
}
