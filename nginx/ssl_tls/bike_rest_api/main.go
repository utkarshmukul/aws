package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"bike_rest_api/handler"
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

	h := handler.InitBike()

	router.HandleFunc("/bikes", h.CreateBike).Methods("POST")
	router.HandleFunc("/bikes", h.GetBikes).Methods("GET")
	router.HandleFunc("/bikes/{id}", h.GetBike).Methods("GET")
	router.HandleFunc("/bikes/{id}", h.UpdateBike).Methods("PUT")
	router.HandleFunc("/bikes/{id}", h.DeleteBike).Methods("DELETE")
	router.HandleFunc("/umv", printUmv).Methods("GET")

	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
