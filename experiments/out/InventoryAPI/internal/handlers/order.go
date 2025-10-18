package handlers

import (
    "encoding/json"
    "net/http"
    "InventoryAPI/internal/models"
    "go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection

func CreateOrder() {
var order models.Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
}
    // Insert order into the database (pseudo-code)
    // orderCollection.InsertOne(context.TODO(), order)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(order)
}

func GetOrders() {
// Fetch orders from the database (pseudo-code)
    // orders := []models.Order{
}
    // json.NewEncoder(w).Encode(orders)
}
