package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var store = make([]Person, 0)

func handleGetPerson(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		name := request.URL.Query()["name"][0]
		for _, structs := range store {
			if structs.Name == name {
				err := json.NewEncoder(writer).Encode(&structs)
				if err != nil {
					log.Fatalln("There was an error encoding the initialized struct")
				}
			}
		}
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
	}
}

func handleAddPerson(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method == "POST" {
		writer.WriteHeader(http.StatusOK)
		var human Person
		err := json.NewDecoder(request.Body).Decode(&human)
		if err != nil {
			log.Fatalln("There was an error decoding the request body into the struct")
		}
		store = append(store, human)
		err = json.NewEncoder(writer).Encode(&human)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
	}
}

func handleDeletePerson(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if request.Method == "DELETE" {
		name := request.URL.Query()["name"][0]
		indexChoice := 0
		for index, structs := range store {
			if structs.Name == name {
				indexChoice = index
			}
		}
		store = append(store[:indexChoice], store[indexChoice+1:]...)
		writer.Write([]byte("Deleted It"))
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
	}
}
func main() {

	http.HandleFunc("/api/v1/getPerson", handleGetPerson)
	http.HandleFunc("/api/v1/addPerson", handleAddPerson)
	http.HandleFunc("/api/v1/deletePerson", handleDeletePerson)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}
}
