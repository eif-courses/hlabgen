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

// GetRestaurants handles fetching all restaurants.
func GetRestaurants() {
	// Implementation for fetching restaurants
}

// UpdateRestaurant handles updating an existing restaurant.
func UpdateRestaurant() {
	// Implementation for updating a restaurant
}

// DeleteRestaurant handles deleting a restaurant.
func DeleteRestaurant() {
	// Implementation for deleting a restaurant
}
