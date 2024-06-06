package main

import (
	"github.com/aslammulla/techiebutler-assignment/handlers"
	"github.com/aslammulla/techiebutler-assignment/middleware"
	"github.com/aslammulla/techiebutler-assignment/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the employee repository
	employeeRepository := models.NewEmployeeRepository()

	// Initialize the employee handler
	employeeHandler := handlers.NewEmployeeHandler(employeeRepository)

	// as its simple CRUD program so keeping the routes in main. For complex project we can create separate package for it.
	api := r.Group("/api", middleware.AuthMiddleware)
	{
		// Routes
		api.GET("/employees", employeeHandler.ListEmployeesHandler)
		api.POST("/employees", employeeHandler.CreateEmployeeHandler)
		api.GET("/employees/:id", employeeHandler.GetEmployeeByIDHandler)
		api.PUT("/employees/:id", employeeHandler.UpdateEmployeeHandler)
		api.DELETE("/employees/:id", employeeHandler.DeleteEmployeeHandler)
	}

	// Start the server
	r.Run(":8080")
}
