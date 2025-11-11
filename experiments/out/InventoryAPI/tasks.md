# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- InventoryItem
- Reservation

## Tasks to Complete

1. Test all CRUD operations for InventoryItem
2. Add validation for inventoryitem fields
3. Test all CRUD operations for Reservation
4. Add validation for reservation fields
5. Implement database persistence (currently using in-memory storage)
6. Add authentication middleware
7. Implement pagination for list endpoints
8. Add error handling and logging
9. Write integration tests
10. Add API documentation (Swagger/OpenAPI)

## Running the Application

```bash
# Install dependencies
go mod tidy

# Run the server
go run cmd/main.go

# Run tests
go test ./...

# Check coverage
go test -cover ./...
```

## API Endpoints

### InventoryItem

- `POST /inventoryitems` - Create a new inventoryitem
- `GET /inventoryitems` - Get all inventoryitems
- `GET /inventoryitems/{id}` - Get a specific inventoryitem
- `PUT /inventoryitems/{id}` - Update a inventoryitem
- `DELETE /inventoryitems/{id}` - Delete a inventoryitem

### Reservation

- `POST /reservations` - Create a new reservation
- `GET /reservations` - Get all reservations
- `GET /reservations/{id}` - Get a specific reservation
- `PUT /reservations/{id}` - Update a reservation
- `DELETE /reservations/{id}` - Delete a reservation

