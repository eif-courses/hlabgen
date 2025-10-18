# Lab Tasks

## Task 1: Implement Role-Based Access Control
### Statement
Implement role-based access control for the API, ensuring that only users with the appropriate roles can create, update, or delete tasks, projects, users, and teams.
### Hints
- Use middleware to check user roles.
- Define roles in the User model.
### Expected Output
Users with role 'admin' can perform all actions, while 'user' can only read.

## Task 2: Add Task Status Filtering
### Statement
Implement filtering for tasks based on their status (e.g., 'completed', 'in-progress').
### Hints
- Use query parameters to filter tasks.
- Modify the GetTasks handler to accept a status filter.
### Expected Output
GET /tasks?status=completed returns only completed tasks.

## Task 3: Implement Team Collaboration Features
### Statement
Allow users to collaborate on tasks by assigning tasks to teams and allowing team members to update task status.
### Hints
- Modify the Task model to include a TeamID.
- Ensure that only team members can update tasks.
### Expected Output
Team members can update tasks assigned to their team.