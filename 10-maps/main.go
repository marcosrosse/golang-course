package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Marcos",
		"nickname": "Rosse",
		"email":    "marcosrossem@gmail.com",
	}
	fmt.Println(user["email"])
	fmt.Println(user)

	user2 := map[string]map[string]string{
		"person": {
			"first_name": "Marcos",
			"last_name":  "Rosse",
		},
		"course": {
			"name": "Golang",
			"lang": "Go",
		},
	}
	fmt.Println(user2)

	delete(user2, "course")

	fmt.Println(user2)

	user2["job"] = map[string]string{
		"company": "Avanade",
		"title":   "DevOps Engineer",
	}
	fmt.Println(user2)
}
