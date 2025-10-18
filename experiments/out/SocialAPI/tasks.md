# Lab Tasks

## Task 1: Implement Social Feed Generation
- **Statement:** Create an endpoint that returns a user's feed based on their follows and posts.
- **Hints:** Use joins to fetch posts from followed users.
- **Expected Output:** JSON array of posts.

## Task 2: Add Follow/Unfollow Functionality
- **Statement:** Implement the follow and unfollow endpoints.
- **Hints:** Ensure to handle duplicate follows and remove entries correctly.
- **Expected Output:** Status 200 for success, 400 for errors.

## Task 3: Implement Notification System
- **Statement:** Create a notification system that alerts users when someone follows them or likes their post.
- **Hints:** Use a queue or a simple database table to store notifications.
- **Expected Output:** JSON array of notifications for a user.