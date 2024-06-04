package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Plant struct {
	ID          string    `json:"id"`
	STEP        string    `json:"step"`
	DATE        string    `json:"date"`
	Discription string    `json:"discription"`
	Check       bool      `json:"check"`
	Director    *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var plants []Plant

func main() {
	r := mux.NewRouter()

	plants = append(plants, Plant{ID: "1", STEP: "1", DATE: "30-05-2567", Discription: "เพาะเมล็ดผักสลัด", Check: true, Director: &Director{Firstname: "คันธมาลี", Lastname: "นาอุดม"}})
	plants = append(plants, Plant{ID: "2", STEP: "2", DATE: "31-05-2567", Discription: "ต้นกล้ารับแสง", Check: true, Director: &Director{Firstname: "คันธมาลี", Lastname: "นาอุดม"}})

	r.HandleFunc("/plants", getPlants).Methods("GET")
	r.HandleFunc("/plants/{id}", getPlant).Methods("GET")
	r.HandleFunc("/plants", createPlant).Methods("POST")
	r.HandleFunc("/plants/{id}", updatePlant).Methods("PUT")
	r.HandleFunc("/plants/{id}", deletePlant).Methods("DELETE")

	fmt.Println("Start server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getPlants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plants)
}

func getPlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range plants {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}

}
func createPlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var plant Plant
	_ = json.NewDecoder(r.Body).Decode(&plant)
	plant.ID = strconv.Itoa(rand.Intn(100000000))
	plants = append(plants, plant)
	json.NewEncoder(w).Encode(plant)

}
func updatePlant(w http.ResponseWriter, r *http.Request) {
	//set json Content-Type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies and range
	for index, item := range plants {
		if item.ID == params["id"] {
			//delete item
			plants = append(plants[:index], plants[index+1:]...)
			//add item
			var plant Plant
			_ = json.NewDecoder(r.Body).Decode(&plant)
			plant.ID = strconv.Itoa(rand.Intn(100000000))
			plants = append(plants, plant)
			//call
			json.NewEncoder(w).Encode(plant)
		}
	}

}
func deletePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range plants {
		if item.ID == params["id"] {
			plants = append(plants[:index], plants[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(plants)
}
