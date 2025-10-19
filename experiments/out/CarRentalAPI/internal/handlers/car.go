package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

var cars []models.Car

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cars = append(cars, car)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}
