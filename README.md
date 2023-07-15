# todo_API_golang

This is a learning project for building a backend API using Go and Gin. It provides basic CRUD operations for managing todos.
You can find the video I used for learning and creating this project here: https://youtu.be/d_L64KT3SFM

## Endpoints
GET /todos: Retrieves all todos.
POST /todos: Adds a new todo.
GET /todos/:id: Retrieves a specific todo by its ID.
PATCH /todos/:id: Updates the status of a todo.

## Usage
To use the API, you can send HTTP requests to the following endpoints:

GET all todos: http://localhost:8080/todos
POST a new todo: http://localhost:8080/todos
GET a specific todo: http://localhost:8080/todos/:id
PATCH the status of a todo: http://localhost:8080/todos/:id
