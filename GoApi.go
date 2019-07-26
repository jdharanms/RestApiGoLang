package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Args struct {
	Arg    string `json: "arg"`
	Origin string `json: "origin"`
	URL    string `json: "url"`
	ID     string `json: "id"`
}

var argument []Args

type deps func() (*http.Response, error)

// variable assigned to anonymous function which return reponse & error of 3rd Party API end point
var dependencies = func() (*http.Response, error) {
	return http.Get("https://httpbin.org/get")
}

// GetPerson end point which hits 3rd party API and gets response from it /
func GetPerson(w http.ResponseWriter, req *http.Request) {

	// reponse, err := http.Get("https://httpbin.org/get")
	reponse, err := dependencies()
	var datastring string
	dataPeople := Args{}
	if err != nil {
		fmt.Printf("The HTTP Request failed with error %s\n", err)
		http.Error(w, "Error in handling", http.StatusBadGateway)
	} else if reponse.StatusCode != http.StatusOK {
		fmt.Printf("The HTTP Request failed with error ")
		http.Error(w, "Third party API Internal error", http.StatusNotFound)

	} else {
		// defer reponse.Body.Close()
		data, _ := ioutil.ReadAll(reponse.Body)
		datastring = string(data)
		json.Unmarshal([]byte(datastring), &dataPeople)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(dataPeople)
	}

}

// CreatePerson end point which hits 3rd party API to post person details into /
func CreatePerson(w http.ResponseWriter, req *http.Request) {

	jsonData := map[string]string{"url": "httsp://myurl", "origin": "192.00.00.8, 192.00.00.80"}
	jsonValue, _ := json.Marshal(jsonData)
	reponse, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP Request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(reponse.Body)
		fmt.Println(string(data))
		fmt.Println(string(reponse.Status))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

}
func DeletePerson(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8004", router))

}
