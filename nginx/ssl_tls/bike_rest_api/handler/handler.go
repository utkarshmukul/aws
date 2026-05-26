package handler

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

type Handler interface {
	CreateBike(w http.ResponseWriter, r *http.Request)
	GetBikes(w http.ResponseWriter, r *http.Request)
	GetBike(w http.ResponseWriter, r *http.Request)
	UpdateBike(w http.ResponseWriter, r *http.Request)
	DeleteBike(w http.ResponseWriter, r *http.Request)
}

type handlerCtx struct {
	nextID int
	bikeList []Bike
}

type Bike struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Year  int    `json:"year"`
}

func InitBike() Handler {
	return &handlerCtx{}
}

// CREATE CAR
func (h *handlerCtx) CreateBike(w http.ResponseWriter, r *http.Request) {
	var bike Bike

	err := json.NewDecoder(r.Body).Decode(&bike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bike.ID = h.nextID
	h.nextID++

	h.bikeList = append(h.bikeList, bike)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bike)
}

// GET ALL CARS
func (h *handlerCtx) GetBikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.bikeList)
}

// GET SINGLE CAR
func (h *handlerCtx) GetBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, bike := range h.bikeList {
		if bike.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(bike)
			return
		}
	}

	http.Error(w, "Bike not found", http.StatusNotFound)
}

// UPDATE CAR
func (h *handlerCtx) UpdateBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedBike Bike
	err = json.NewDecoder(r.Body).Decode(&updatedBike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, bike := range h.bikeList {
		if bike.ID == id {
			updatedBike.ID = id
			h.bikeList[index] = updatedBike

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBike)
			return
		}
	}

	http.Error(w, "Bike not found", http.StatusNotFound)
}

// DELETE CAR
func (h *handlerCtx) DeleteBike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for index, bike := range h.bikeList {
		if bike.ID == id {
			h.bikeList = append(h.bikeList[:index], h.bikeList[index+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Bike deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Bike not found", http.StatusNotFound)
}
