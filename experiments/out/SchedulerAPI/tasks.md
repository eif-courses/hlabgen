# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- ScheduledJob
- ExecutionLog

## Tasks to Complete

1. Test all CRUD operations for ScheduledJob
2. Add validation for scheduledjob fields
3. Test all CRUD operations for ExecutionLog
4. Add validation for executionlog fields
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

### ScheduledJob

- `POST /scheduledjobs` - Create a new scheduledjob
- `GET /scheduledjobs` - Get all scheduledjobs
- `GET /scheduledjobs/{id}` - Get a specific scheduledjob
- `PUT /scheduledjobs/{id}` - Update a scheduledjob
- `DELETE /scheduledjobs/{id}` - Delete a scheduledjob

### ExecutionLog

- `POST /executionlogs` - Create a new executionlog
- `GET /executionlogs` - Get all executionlogs
- `GET /executionlogs/{id}` - Get a specific executionlog
- `PUT /executionlogs/{id}` - Update a executionlog
- `DELETE /executionlogs/{id}` - Delete a executionlog

