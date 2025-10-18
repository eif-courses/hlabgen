# Lab Tasks

## Task 1: Implement JWT Authentication
- **Statement**: Implement secure authentication using JWT for the customer entity.
- **Hints**: Use the `github.com/dgrijalva/jwt-go` package. Create a login endpoint that returns a JWT.
- **Expected Output**: A working login endpoint that returns a JWT token.

## Task 2: Transactional Order Processing
- **Statement**: Implement transactional order processing when creating an order.
- **Hints**: Use database transactions to ensure that order creation and payment processing are atomic.
- **Expected Output**: An order creation endpoint that processes payments atomically.

## Task 3: Complex Search with Filters
- **Statement**: Implement a search endpoint for products that supports filtering by name and price range.
- **Hints**: Use query parameters to filter results and implement pagination.
- **Expected Output**: A search endpoint that returns filtered and paginated product results.
