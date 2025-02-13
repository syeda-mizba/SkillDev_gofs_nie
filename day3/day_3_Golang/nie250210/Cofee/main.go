package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Coffee struct defines the coffee data model
type Coffee struct {
	Id         string  `json:"id"`
	Category   string  `json:"category"` // Small, Medium, Large
	Type       string  `json:"type"`     // Filter, Instant
	Price      float32 `json:"price"`
	SugarLevel string  `json:"sugar_level"` // No, Less, Normal
}

// In-memory store for coffee items (for testing purposes)
var coffees = []Coffee{
	{Id: "1001", Category: "Medium", Type: "Filter", Price: 100.0, SugarLevel: "Normal"},
	{Id: "1002", Category: "Large", Type: "Instant", Price: 150.0, SugarLevel: "Less"},
}

// Function to get all coffee items
func readAllCoffees(c *gin.Context) {
	c.JSON(http.StatusOK, coffees)
}

// Function to get coffee by ID
func readCoffeeById(c *gin.Context) {
	id := c.Param("id")
	for _, coffee := range coffees {
		if coffee.Id == id {
			c.JSON(http.StatusOK, coffee)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
}

// Function to create a new coffee
func createCoffee(c *gin.Context) {
	var newCoffee Coffee
	err := c.BindJSON(&newCoffee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ID (this could be handled by a DB in a real application)
	newCoffee.Id = fmt.Sprintf("%d", len(coffees)+1)

	// Add to the in-memory coffee list
	coffees = append(coffees, newCoffee)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Coffee created successfully",
		"coffee":  newCoffee,
	})
}

// Function to update an existing coffee by ID
func updateCoffee(c *gin.Context) {
	id := c.Param("id")
	var updatedCoffee Coffee
	err := c.BindJSON(&updatedCoffee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, coffee := range coffees {
		if coffee.Id == id {
			// Update the coffee details
			coffees[i] = updatedCoffee
			coffees[i].Id = id // Retain the original ID
			c.JSON(http.StatusOK, gin.H{
				"message": "Coffee updated successfully",
				"coffee":  coffees[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
}

// Function to delete a coffee by ID
func deleteCoffee(c *gin.Context) {
	id := c.Param("id")
	for i, coffee := range coffees {
		if coffee.Id == id {
			// Remove the coffee from the list
			coffees = append(coffees[:i], coffees[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Coffee deleted successfully",
				"id":      id,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define the routes (API Endpoints)
	r.GET("/coffees", readAllCoffees)      // Get all coffee items
	r.GET("/coffees/:id", readCoffeeById)  // Get coffee by ID
	r.POST("/coffees", createCoffee)       // Create a new coffee
	r.PUT("/coffees/:id", updateCoffee)    // Update coffee by ID
	r.DELETE("/coffees/:id", deleteCoffee) // Delete coffee by ID

	// Run the server (on port 8080 by default)
	r.Run(":8083")
}
