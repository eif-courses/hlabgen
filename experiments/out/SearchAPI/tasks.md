# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Product

## Tasks to Complete

1. Test all CRUD operations for Product
2. Add validation for product fields
3. Implement database persistence (currently using in-memory storage)
4. Add authentication middleware
5. Implement pagination for list endpoints
6. Add error handling and logging
7. Write integration tests
8. Add API documentation (Swagger/OpenAPI)

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

### Product

- `POST /products` - Create a new product
- `GET /products` - Get all products
- `GET /products/{id}` - Get a specific product
- `PUT /products/{id}` - Update a product
- `DELETE /products/{id}` - Delete a product

