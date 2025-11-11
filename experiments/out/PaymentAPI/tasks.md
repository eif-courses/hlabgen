# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Payment

## Tasks to Complete

1. Test all CRUD operations for Payment
2. Add validation for payment fields
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

### Payment

- `POST /payments` - Create a new payment
- `GET /payments` - Get all payments
- `GET /payments/{id}` - Get a specific payment
- `PUT /payments/{id}` - Update a payment
- `DELETE /payments/{id}` - Delete a payment

