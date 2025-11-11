# Lab Tasks

## Overview

This project implements a REST API with the following entities:

- ApprovalRequest
- Approval

## Tasks to Complete

1. Test all CRUD operations for ApprovalRequest
2. Add validation for approvalrequest fields
3. Test all CRUD operations for Approval
4. Add validation for approval fields
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

### ApprovalRequest

- `POST /approvalrequests` - Create a new approvalrequest
- `GET /approvalrequests` - Get all approvalrequests
- `GET /approvalrequests/{id}` - Get a specific approvalrequest
- `PUT /approvalrequests/{id}` - Update a approvalrequest
- `DELETE /approvalrequests/{id}` - Delete a approvalrequest

### Approval

- `POST /approvals` - Create a new approval
- `GET /approvals` - Get all approvals
- `GET /approvals/{id}` - Get a specific approval
- `PUT /approvals/{id}` - Update a approval
- `DELETE /approvals/{id}` - Delete a approval

