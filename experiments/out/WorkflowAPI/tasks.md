# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- Workflow
- ApprovalRecord

## Tasks to Complete

1. Test all CRUD operations for Workflow
2. Add validation for workflow fields
3. Test all CRUD operations for ApprovalRecord
4. Add validation for approvalrecord fields
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

### Workflow

- `POST /workflows` - Create a new workflow
- `GET /workflows` - Get all workflows
- `GET /workflows/{id}` - Get a specific workflow
- `PUT /workflows/{id}` - Update a workflow
- `DELETE /workflows/{id}` - Delete a workflow

### ApprovalRecord

- `POST /approvalrecords` - Create a new approvalrecord
- `GET /approvalrecords` - Get all approvalrecords
- `GET /approvalrecords/{id}` - Get a specific approvalrecord
- `PUT /approvalrecords/{id}` - Update a approvalrecord
- `DELETE /approvalrecords/{id}` - Delete a approvalrecord

