package handler

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"os"
	"fmt"
)

type Handler interface {
	CreateCar(w http.ResponseWriter, r *http.Request)
	GetCars(w http.ResponseWriter, r *http.Request)
	GetCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
	DeleteCar(w http.ResponseWriter, r *http.Request)
}

type handlerCtx struct {
	nextID int
	carList []Car
}

type Car struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Year  int    `json:"year"`
}

func InitCar() Handler {
	return &handlerCtx{}
}

// CREATE CAR
func (h *handlerCtx) CreateCar(w http.ResponseWriter, r *http.Request) {
	var car Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.ID = h.nextID
	h.nextID++

	h.carList = append(h.carList, car)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

// GET ALL CARS
func (h *handlerCtx) GetCars(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Println("UMV GOT Request from : ", hostname)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.carList)
}

// GET SINGLE CAR
func (h *handlerCtx) GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, car := range h.carList {
		if car.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// UPDATE CAR
func (h *handlerCtx) UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	err = json.NewDecoder(r.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, car := range h.carList {
		if car.ID == id {
			updatedCar.ID = id
			h.carList[index] = updatedCar

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// DELETE CAR
func (h *handlerCtx) DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for index, car := range h.carList {
		if car.ID == id {
			h.carList = append(h.carList[:index], h.carList[index+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Car deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}
