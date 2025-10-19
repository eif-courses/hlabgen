package handlers

import (
	"FoodDeliveryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var restaurants []models.Restaurant

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	restaurants = append(restaurants, restaurant)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}
