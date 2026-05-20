package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"car_rest_api/handler"
)

func printUmv(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Define the data to be returned
	response := map[string]string{"message": "hello world"}

	// Encode and write the response
	json.NewEncoder(w).Encode(response)
}


func main() {
	router := mux.NewRouter()

	h := handler.InitCar()

	router.HandleFunc("/cars", h.CreateCar).Methods("POST")
	router.HandleFunc("/cars", h.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", h.GetCar).Methods("GET")
	router.HandleFunc("/cars/{id}", h.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", h.DeleteCar).Methods("DELETE")
	router.HandleFunc("/umv", printUmv).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
