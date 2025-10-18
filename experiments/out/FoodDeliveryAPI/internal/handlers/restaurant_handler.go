package handlers

import (
	"FoodDeliveryAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateRestaurant handles the creation of a new restaurant.
func CreateRestaurant() {
	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}
