package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type args struct {
	w   http.ResponseWriter
	req *http.Request
}

func setup(t *testing.T) (*httptest.ResponseRecorder, error) {
	dependencies = func() (*http.Response, error) {
		return &http.Response{}, http.ErrShortBody
	}
	return nil, nil
}

func getargs(arg args) args {
	var a args
	a.w = httptest.NewRecorder()
	a.req, _ = http.NewRequest("GET", "localhost:8004/people", nil)
	return a
}

func TestCreatePerson1(t *testing.T) {
	setup(t)
	var a args
	a.w = httptest.NewRecorder()
	a.req, _ = http.NewRequest("GET", "localhost:8004/people", nil)
	GetPerson(a.w, a.req)
}

func TestGetPersonStatusBadGateway(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{name: "Happy path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argssss := getargs(tt.args)
			dependencies = func() (*http.Response, error) {
				return &http.Response{}, http.ErrShortBody
			}
			rec := httptest.NewRecorder()
			GetPerson(rec, argssss.req)
			res := rec.Result()
			if res.StatusCode != http.StatusBadGateway {
				t.Errorf("Error in handling:  got %v", res.Status)
			}
		})
	}
}

func TestGetPersonStatusNotFound(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{name: "Happy path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argssss := getargs(tt.args)
			dependencies = func() (*http.Response, error) {
				return &http.Response{}, nil
			}
			rec := httptest.NewRecorder()
			GetPerson(rec, argssss.req)
			res := rec.Result()
			if res.StatusCode != http.StatusNotFound {
				t.Errorf(" 405 Method Not Allowed:  got %v", res.Status)
			}

		})
	}
}
func TestGetPersonHappyPath(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{name: "Happy path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dependencies = func() (*http.Response, error) {
				return &http.Response{Status: "OK", StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(`OK`))}, nil
			}
			argssss := getargs(tt.args)
			rec := httptest.NewRecorder()
			GetPerson(rec, argssss.req)
			res := rec.Result()
			if res.StatusCode != http.StatusFound {
				t.Errorf(" Expected Ok:  got %v", res.Status)
			}
		})
	}
}

func TestGetPersonHappyPath1(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{name: "Happy path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dependencies = func() (*http.Response, error) {
				return http.Get("https://httpbin.org/get")
			}
			argssss := getargs(tt.args)
			rec := httptest.NewRecorder()
			GetPerson(rec, argssss.req)
			res := rec.Result()
			if res.StatusCode != http.StatusFound {
				t.Errorf(" Expected Ok:  got %v", res.Status)
			}
		})
	}
}
func TestCreatePerson(t *testing.T) {

	tests := []struct {
		name string
		args args
	}{
		{name: "Happy path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argssss := getargs(tt.args)
			CreatePerson(argssss.w, argssss.req)
		})
	}
}
