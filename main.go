package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee represents an employee record
type Employee struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Department string `json:"department"`
	Salary    float64 `json:"salary"`
}

var filePath = "employees.json"

// Load employees from JSON file
func loadEmployeesFromFile() (map[int]Employee, error) {
	employees := make(map[int]Employee)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return employees, nil
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &employees)
	return employees, err
}

// Save employees to JSON file
func saveEmployeesToFile(employees map[int]Employee) error {
	data, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := loadEmployeesFromFile()
	if err != nil {
		http.Error(w, "Failed to load employees", http.StatusInternalServerError)
		return
	}
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
	employees, err := loadEmployeesFromFile()
	if err != nil {
		http.Error(w, "Failed to load employees", http.StatusInternalServerError)
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
	employees, err := loadEmployeesFromFile()
	if err != nil {
		http.Error(w, "Failed to load employees", http.StatusInternalServerError)
		return
	}
	idCounter := len(employees) + 1
	employee.ID = idCounter
	employees[employee.ID] = employee
	if err := saveEmployeesToFile(employees); err != nil {
		http.Error(w, "Failed to save employee", http.StatusInternalServerError)
		return
	}
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
	employees, err := loadEmployeesFromFile()
	if err != nil {
		http.Error(w, "Failed to load employees", http.StatusInternalServerError)
		return
	}
	_, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	updatedEmployee.ID = id
	employees[id] = updatedEmployee
	if err := saveEmployeesToFile(employees); err != nil {
		http.Error(w, "Failed to save employee", http.StatusInternalServerError)
		return
	}
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
	employees, err := loadEmployeesFromFile()
	if err != nil {
		http.Error(w, "Failed to load employees", http.StatusInternalServerError)
		return
	}
	_, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	delete(employees, id)
	if err := saveEmployeesToFile(employees); err != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}
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
