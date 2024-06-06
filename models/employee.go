package models

import (
	"fmt"
	"sync"
)

// Employee struct represents an employee
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// EmployeeRepository represents a store for managing employees
type EmployeeRepository struct {
	employees map[int]Employee
	mu        sync.RWMutex
}

// NewEmployeeRepository creates a new instance of EmployeeRepository
func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		employees: make(map[int]Employee),
	}
}

// CreateEmployee adds a new employee to the repository
func (er *EmployeeRepository) CreateEmployee(e Employee) error {
	er.mu.Lock()
	defer er.mu.Unlock()

	if _, exists := er.employees[e.ID]; exists {
		return fmt.Errorf("Employee with ID %d already exists", e.ID)
	}

	er.employees[e.ID] = e
	return nil
}

// GetEmployeeByID retrieves an employee from the repository by ID
func (er *EmployeeRepository) GetEmployeeByID(id int) (Employee, bool) {
	er.mu.RLock()
	defer er.mu.RUnlock()
	e, ok := er.employees[id]
	return e, ok
}

// UpdateEmployee updates an existing employee in the repository
func (er *EmployeeRepository) UpdateEmployee(id int, e Employee) error {
	er.mu.Lock()
	defer er.mu.Unlock()

	if _, exists := er.employees[id]; !exists {
		return fmt.Errorf("Employee with ID %d not found", id)
	}

	er.employees[id] = e
	return nil
}

// DeleteEmployee deletes an employee from the repository by ID
func (er *EmployeeRepository) DeleteEmployee(id int) error {
	er.mu.Lock()
	defer er.mu.Unlock()

	if _, exists := er.employees[id]; !exists {
		return fmt.Errorf("Employee with ID %d not found", id)
	}

	delete(er.employees, id)
	return nil
}

// ListEmployees lists employees with pagination
func (er *EmployeeRepository) ListEmployees(page, pageSize int) ([]Employee, error) {
	er.mu.RLock()
	defer er.mu.RUnlock()

	var employees []Employee
	start := (page - 1) * pageSize
	end := start + pageSize

	for _, e := range er.employees {
		employees = append(employees, e)
	}

	if start < 0 || start >= len(employees) {
		return nil, fmt.Errorf("Invalid page number")
	}

	if end > len(employees) {
		end = len(employees)
	}

	return employees[start:end], nil
}
