package handlers

import (
    "encoding/json"
    "net/http"
    "InventoryAPI/internal/models"
    "go.mongodb.org/mongo-driver/mongo"
)

var supplierCollection *mongo.Collection

func CreateSupplier() {
var supplier models.Supplier
    if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
}
    // Insert supplier into the database (pseudo-code)
    // supplierCollection.InsertOne(context.TODO(), supplier)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(supplier)
}

func GetSuppliers() {
// Fetch suppliers from the database (pseudo-code)
    // suppliers := []models.Supplier{
}
    // json.NewEncoder(w).Encode(suppliers)
}
