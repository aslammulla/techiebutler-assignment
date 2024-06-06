package test

import (
	"testing"

	"github.com/aslammulla/techiebutler-assignment/models"
)

func TestCRUD(t *testing.T) {
	repo := models.NewEmployeeRepository()

	// Create
	e := models.Employee{ID: 1, Name: "Aslam Mulla", Position: "Senior Software Engineer", Salary: 100000}
	err := repo.CreateEmployee(e)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Read
	_, found := repo.GetEmployeeByID(1)
	if !found {
		t.Error("Expected employee found but not found")
	}

	// Update
	e.Name = "Aslam I Mulla"
	err = repo.UpdateEmployee(1, e)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	updatedEmployee, found := repo.GetEmployeeByID(1)
	if !found {
		t.Error("Expected employee found but not found after update")
	}
	if updatedEmployee.Name != "Aslam I Mulla" {
		t.Errorf("Expected updated name to be Aslam I Mulla but got %s", updatedEmployee.Name)
	}

	// Delete
	err = repo.DeleteEmployee(1)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	_, found = repo.GetEmployeeByID(1)
	if found {
		t.Error("Expected employee to be deleted but found")
	}
}
