package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Flight struct to define the flight data model
type Flight struct {
	Id          string  `json:"id"`
	Number      string  `json:"number"`
	AirlineName string  `json:"airline_name"` // Fixed typo "ArilineName"
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Capacity    int     `json:"capacity"`
	Price       float32 `json:"price"`
}

// In-memory store for flights (for testing purposes)
var flights = []Flight{
	{Id: "1001", Number: "AI 845", AirlineName: "Air India", Source: "Mumbai", Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0},
	{Id: "1002", Number: "AI 846", AirlineName: "Air India", Source: "Abu Dhabi", Destination: "Mumbai", Capacity: 180, Price: 15000.0},
}

// Function to get all flights
func readAllFlights(c *gin.Context) {
	c.JSON(http.StatusOK, flights)
}

// Function to get flight by ID
func readFlightById(c *gin.Context) {
	id := c.Param("id")
	for _, flight := range flights {
		if flight.Id == id {
			c.JSON(http.StatusOK, flight)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
}

// Function to create a new flight
func createFlight(c *gin.Context) {
	var newFlight Flight
	err := c.BindJSON(&newFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ID (this could be handled by a DB in a real application)
	newFlight.Id = fmt.Sprintf("%d", len(flights)+1)

	// Add to the in-memory flight list
	flights = append(flights, newFlight)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Flight created successfully",
		"flight":  newFlight,
	})
}

// Function to update an existing flight by ID
func updateFlight(c *gin.Context) {
	id := c.Param("id")
	var updatedFlight Flight
	err := c.BindJSON(&updatedFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, flight := range flights {
		if flight.Id == id {
			// Update the flight details
			flights[i] = updatedFlight
			flights[i].Id = id // Retain the original ID
			c.JSON(http.StatusOK, gin.H{
				"message": "Flight updated successfully",
				"flight":  flights[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
}

// Function to delete a flight by ID
func deleteFlight(c *gin.Context) {
	id := c.Param("id")
	for i, flight := range flights {
		if flight.Id == id {
			// Remove the flight from the list
			flights = append(flights[:i], flights[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Flight deleted successfully",
				"id":      id,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define the routes (API Endpoints)
	r.GET("/flights", readAllFlights)      // Get all flights
	r.GET("/flights/:id", readFlightById)  // Get flight by ID
	r.POST("/flights", createFlight)       // Create a new flight
	r.PUT("/flights/:id", updateFlight)    // Update flight by ID
	r.DELETE("/flights/:id", deleteFlight) // Delete flight by ID

	// Run the server (on port 8080 by default)
	r.Run(":8080")
}
