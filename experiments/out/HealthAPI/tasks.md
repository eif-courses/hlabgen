# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Patient
- Doctor
- Record

## Tasks to Complete

1. Test all CRUD operations for Patient
2. Add validation for patient fields
3. Test all CRUD operations for Doctor
4. Add validation for doctor fields
5. Test all CRUD operations for Record
6. Add validation for record fields
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

### Patient

- `POST /patients` - Create a new patient
- `GET /patients` - Get all patients
- `GET /patients/{id}` - Get a specific patient
- `PUT /patients/{id}` - Update a patient
- `DELETE /patients/{id}` - Delete a patient

### Doctor

- `POST /doctors` - Create a new doctor
- `GET /doctors` - Get all doctors
- `GET /doctors/{id}` - Get a specific doctor
- `PUT /doctors/{id}` - Update a doctor
- `DELETE /doctors/{id}` - Delete a doctor

### Record

- `POST /records` - Create a new record
- `GET /records` - Get all records
- `GET /records/{id}` - Get a specific record
- `PUT /records/{id}` - Update a record
- `DELETE /records/{id}` - Delete a record

