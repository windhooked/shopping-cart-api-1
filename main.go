package main

import (
    "encoding/json"
    // "os"
    // "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    // "github.com/vilst3r/golang-boilerplate/config"
)

type Person struct {
	ID string `json:"id, omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    // Address   *Address `json:"address,omitempty"`	
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe"})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe"})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})		
	json.NewEncoder(w).Encode(people)
}

// our main function
func main() {
	// load application configurations
	// if err := app.LoadConfig("./app/config"); err != nil {
	// 	panic(fmt.Errorf("Invalid application configuration: %s", err))
	// }

    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeople).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}