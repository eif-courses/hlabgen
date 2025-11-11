# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Notification

## Tasks to Complete

1. Test all CRUD operations for Notification
2. Add validation for notification fields
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

### Notification

- `POST /notifications` - Create a new notification
- `GET /notifications` - Get all notifications
- `GET /notifications/{id}` - Get a specific notification
- `PUT /notifications/{id}` - Update a notification
- `DELETE /notifications/{id}` - Delete a notification

