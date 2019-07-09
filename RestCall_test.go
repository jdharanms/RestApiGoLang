package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPerson(t *testing.T) {
	// register mock for external API endpoint
	req, err := http.NewRequest("GET", "localhost:8004/people/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	GetPerson(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Errorf("expected status Ok:  got %v", res.Status)
	}

}

func TestCreatePerson(t *testing.T) {

	jsonData := map[string]string{"ID": "1", "Firstname": "Muraliney", "Lastname": "Clooney"}
	jsonValue, _ := json.Marshal(jsonData)
	// bytes.NewBuffer(jsonValue)
	req, err := http.NewRequest("POST", "localhost:8004/people/{3}", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	CreatePerson(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Ok:  got %v", res.Status)
	}
	// // Check the response body is what we expect.
	// expected := `[{"ID":"1","Firstname":"Murali","Lastname":"Jeyachan"},{"ID":"2","Firstname":"Vinoth","Lastname":"Jeyachan"}]`
	// if rec.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rec.Body.String(), expected)
	// }
}
