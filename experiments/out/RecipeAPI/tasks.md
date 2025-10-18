# Lab Tasks

## Task 1: Implement Create Recipe
- **Statement**: Implement the `CreateRecipe` function to add a new recipe to the database.
- **Hints**: Use `json.NewDecoder(r.Body).Decode(&recipe)` to decode the request body.
- **Expected Output**: A new recipe should be added to the database and return a 201 status code.

## Task 2: Implement Ingredient Search
- **Statement**: Implement a search feature for ingredients based on name.
- **Hints**: Use query parameters to filter ingredients by name.
- **Expected Output**: A list of ingredients matching the search criteria should be returned.

## Task 3: Write CRUD Tests
- **Statement**: Write tests for all CRUD operations for recipes.
- **Hints**: Use `httptest.NewRequest` to create requests for testing.
- **Expected Output**: All tests should pass, confirming that CRUD operations work as expected.