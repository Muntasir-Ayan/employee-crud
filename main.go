package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee represents an employee record
type Employee struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	Department string `json:"department"`
	Salary    float64 `json:"salary"`
}

var employees = make(map[int]Employee)
var idCounter = 1



// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(employees)
}

// Get a single employee
func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	employee, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

// Create a new employee
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	employee.ID = idCounter
	idCounter++
	employees[employee.ID] = employee
	fmt.Println("Employee added")
	json.NewEncoder(w).Encode(employee)
}

// Update an employee
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	var updatedEmployee Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	_, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	updatedEmployee.ID = id
	employees[id] = updatedEmployee
	fmt.Println("Employee record upadated: ",id)
	json.NewEncoder(w).Encode(updatedEmployee)
}

// Delete an employee
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	_, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	fmt.Println("Dleted Employee id: ",id)
	delete(employees, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", createEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8080", r)
}
