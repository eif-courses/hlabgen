package handlers

import (
    "encoding/json"
    "net/http"
    "InventoryAPI/internal/models"
    "go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

func CreateProduct() {
var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
}
    // Insert product into the database (pseudo-code)
    // productCollection.InsertOne(context.TODO(), product)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

func GetProducts() {
// Fetch products from the database (pseudo-code)
    // products := []models.Product{
}
    // json.NewEncoder(w).Encode(products)
}
