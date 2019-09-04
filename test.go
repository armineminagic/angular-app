package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID string 'json:"id,omitempty"'
	Firstname string 'json:"firstname,omitempty"'
	Lastname string 'json:"lastname,omitempty"'
	Adress *Address 'json:"address,omitempty"'
}

type Address struct {
	City string 'json:"city,omitempty"'
	State string 'json:"state,omitempty"'
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var person Person
	_ = json.NevDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = appedn(people, person)
	json.NewEncoder(w).Encode(people)
}