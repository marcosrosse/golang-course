package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marcosrosse/golang-course/22-CRUD/db"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Create users in mysql database
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read the request body!"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error to convert json to struct"))
		return
	}
	db, err := db.Conn()
	if err != nil {
		w.Write([]byte("Error to connect in the db!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error to create statement!"))
		return
	}

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Error to execute the statement!"))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Error getting ID!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User %s inserted with success! Id: %d", user.Name, idInserted)))

}
