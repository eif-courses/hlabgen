# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Book
- Member
- Loan

## Tasks to Complete

1. Test all CRUD operations for Book
2. Add validation for book fields
3. Test all CRUD operations for Member
4. Add validation for member fields
5. Test all CRUD operations for Loan
6. Add validation for loan fields
7. Implement database persistence (currently using in-memory storage)
8. Add authentication middleware
9. Implement pagination for list endpoints
10. Add error handling and logging
11. Write integration tests
12. Add API documentation (Swagger/OpenAPI)

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

### Book

- `POST /books` - Create a new book
- `GET /books` - Get all books
- `GET /books/{id}` - Get a specific book
- `PUT /books/{id}` - Update a book
- `DELETE /books/{id}` - Delete a book

### Member

- `POST /members` - Create a new member
- `GET /members` - Get all members
- `GET /members/{id}` - Get a specific member
- `PUT /members/{id}` - Update a member
- `DELETE /members/{id}` - Delete a member

### Loan

- `POST /loans` - Create a new loan
- `GET /loans` - Get all loans
- `GET /loans/{id}` - Get a specific loan
- `PUT /loans/{id}` - Update a loan
- `DELETE /loans/{id}` - Delete a loan

