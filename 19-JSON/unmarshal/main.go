package main

import (
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
	dog1JSON := `{"name":"Crystal","race":"Poodle","age":6}`

	var dog1 dog

	if err := json.Unmarshal([]byte(dog1JSON), &dog1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog1)

	dog2JSON := `{"name":"Artemis","race":"Pit bull"}`
	dog2 := make(map[string]string)
	if err := json.Unmarshal([]byte(dog2JSON), &dog2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog2)
}
