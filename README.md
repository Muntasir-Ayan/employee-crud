# Employee Management System in Go

This project is a simple CRUD API built with Go to manage employee records. It allows you to create, read, update, and delete employee data via HTTP requests.


# Features
- Create Employee: Add a new employee.
- Read Employee: Retrieve details of all employees or a single employee by ID.
- Update Employee: Modify an existing employee's details.
- Delete Employee: Remove an employee record from the system.

# Technologies Used
- Go: Programming language for backend.
- Gorilla Mux: HTTP router for handling API routes.
- JSON: Data format for requests and responses.
- In-memory storage: Employee data is stored in a Go map for simplicity.

## Getting Started
# Prerequisites

- Install Go (version 1.18 or higher).
- Clone the repository to your local machine.
    ``` bash 
        https://github.com/Muntasir-Ayan/employee-crud.git
        cd employee-crud

    ```
# Installation
1. **Install dependencies:** This project uses the gorilla/mux package to handle HTTP routing. You can install it using:

    ``` bash 
        go get -u github.com/gorilla/mux

    ```
2. **Run the application:** Start the Go server by running:
    ``` bash 
       go run main.go

    ```
    The server will start on http://localhost:8080.

# API Endpoints
1. **Create Employee (POST):**
    - URL: /employees
    - Method: POST
2. **Get All Employees (GET):**
    - URL: /employees
    - Method: GET

3. **Get Employee by ID (GET):**
    - URL: /employees/{id}
    - Method: GET

4. **Update Employee (PUT):**
    - URL: /employees/{id}
    - Method: PUT

5. **Delete Employee (DELETE):**
    - URL: /employees/{id}
    - Method: DELETE
    - Response: No content, HTTP status 204.

# Testing API Endpoints

- **Create Employee:**
    ```bash
        curl -X POST http://localhost:8080/employees \
        -H "Content-Type: application/json" \
        -d '{"name": "Jhon", "position": "Software Engineer", "department": "IT", "salary": 50000}'
    ```
- **Get All Employees:**
    ```bash
        curl -X GET http://localhost:8080/employees
    ```
- **Get Employee by ID:**
    ```bash
        curl -X GET http://localhost:8080/employees/<id>
    ```
- **Update Employee:**
    ```bash
        curl -X PUT http://localhost:8080/employees/<id> \
        -H "Content-Type: application/json" \
        -d '{"name": "Jhon Doe", "position": "Project Manager", "department": "IT", "salary": 50000}'
    ```
- **Delete Employee:**
    ```bash
            curl -X DELETE http://localhost:8080/employees/<id>
    ```


# google slide link: https://docs.google.com/presentation/d/1RZejC58BmxR6qkps-YkvWufwOyas3mimjlQbEiKJJA0/edit#slide=id.g2a4cc96df6c_0_26