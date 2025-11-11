# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Role
- UserRole

## Tasks to Complete

1. Test all CRUD operations for Role
2. Add validation for role fields
3. Test all CRUD operations for UserRole
4. Add validation for userrole fields
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

### Role

- `POST /roles` - Create a new role
- `GET /roles` - Get all roles
- `GET /roles/{id}` - Get a specific role
- `PUT /roles/{id}` - Update a role
- `DELETE /roles/{id}` - Delete a role

### UserRole

- `POST /userroles` - Create a new userrole
- `GET /userroles` - Get all userroles
- `GET /userroles/{id}` - Get a specific userrole
- `PUT /userroles/{id}` - Update a userrole
- `DELETE /userroles/{id}` - Delete a userrole

