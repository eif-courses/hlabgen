# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- User
- Session

## Tasks to Complete

1. Test all CRUD operations for User
2. Add validation for user fields
3. Test all CRUD operations for Session
4. Add validation for session fields
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

### User

- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/{id}` - Get a specific user
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

### Session

- `POST /sessions` - Create a new session
- `GET /sessions` - Get all sessions
- `GET /sessions/{id}` - Get a specific session
- `PUT /sessions/{id}` - Update a session
- `DELETE /sessions/{id}` - Delete a session

