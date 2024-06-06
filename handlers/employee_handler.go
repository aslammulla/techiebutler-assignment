package handlers

import (
	"net/http"
	"strconv"

	"github.com/aslammulla/techiebutler-assignment/models"
	"github.com/gin-gonic/gin"
)

// EmployeeHandler handles HTTP requests related to employees
type EmployeeHandler struct {
	employeeRepository *models.EmployeeRepository
}

// NewEmployeeHandler creates a new instance of EmployeeHandler
func NewEmployeeHandler(employeeRepository *models.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{employeeRepository: employeeRepository}
}

// ListEmployeesHandler handles the listing of employees with pagination
func (eh *EmployeeHandler) ListEmployeesHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	employees, err := eh.employeeRepository.ListEmployees(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// CreateEmployeeHandler handles the creation of a new employee
func (eh *EmployeeHandler) CreateEmployeeHandler(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := eh.employeeRepository.CreateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
}

// GetEmployeeByIDHandler handles the retrieval of an employee by ID
func (eh *EmployeeHandler) GetEmployeeByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, found := eh.employeeRepository.GetEmployeeByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// UpdateEmployeeHandler handles the updating of an existing employee
func (eh *EmployeeHandler) UpdateEmployeeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := eh.employeeRepository.UpdateEmployee(id, employee); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// DeleteEmployeeHandler handles the deletion of an employee by ID
func (eh *EmployeeHandler) DeleteEmployeeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := eh.employeeRepository.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
