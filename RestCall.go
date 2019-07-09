package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json: "id,omitempty"`
	Firstname string `json: "firstname,omitempty"`
	Lastname  string `json: "lastname,omitempty"`
}

var people []Person

func GetPerson(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		http.Error(w, "my own error message", http.StatusForbidden)
	}

}

func CreatePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(people)

}
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Murali", Lastname: "Jeyachan"})
	people = append(people, Person{ID: "2", Firstname: "Vinoth", Lastname: "Jeyachan"})
	router.HandleFunc("/people", GetPerson).Methods("GET")
	fmt.Printf("gets int to API ")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8004", router))

}
