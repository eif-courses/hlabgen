package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCar() {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func GetCars() {
	// Implementation for fetching cars
}
func UpdateCar() {
	// Implementation for updating a car
}
func DeleteCar() {
	// Implementation for deleting a car
}
