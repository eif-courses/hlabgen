# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Order
- OrderItem

## Tasks to Complete

1. Test all CRUD operations for Order
2. Add validation for order fields
3. Test all CRUD operations for OrderItem
4. Add validation for orderitem fields
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

### Order

- `POST /orders` - Create a new order
- `GET /orders` - Get all orders
- `GET /orders/{id}` - Get a specific order
- `PUT /orders/{id}` - Update a order
- `DELETE /orders/{id}` - Delete a order

### OrderItem

- `POST /orderitems` - Create a new orderitem
- `GET /orderitems` - Get all orderitems
- `GET /orderitems/{id}` - Get a specific orderitem
- `PUT /orderitems/{id}` - Update a orderitem
- `DELETE /orderitems/{id}` - Delete a orderitem

