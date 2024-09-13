package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"meeting/config"
	"meeting/controller"
)

const PORT = "8000"

func main() {

	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	router := r.Group("")
	{

		// User Routes
		router.GET("/meetingRoom", controller.ListMeetingRoom)
		router.GET("/meetingRoom/:id", controller.GetMeetingRoom) // Corrected route name
		router.POST("/meetingRoom", controller.CreateMeetingRoom) // Assuming you want to create a meeting room
		router.PATCH("/meetingRoom", controller.UpdateMeetingRoom)
		router.DELETE("/meetingRoom/:id", controller.DeleteMeetingRoom)

		// customer
		router.GET("/customerMeetingRoom", controller.ListCustomerMeetingRoom)
		router.GET("/customerMeetingRoom/:id", controller.GetCustomerMeetingRoom) // Corrected route name
		router.POST("/customerMeetingRoom", controller.CreateCustomerMeetingRoom) // Assuming you want to create a meeting room
		router.PATCH("/customerMeetingRoom", controller.UpdateCustomerMeetingRoom)
		router.DELETE("/customerMeetingRoom/:id", controller.DeleteCustomerMeetingRoom)

		// customer
		router.GET("/bookingMeetingRoom", controller.ListBookingMeetingRoom)
		router.GET("/bookingMeetingRoom/:id", controller.GetBookingMeetingRoom) // Corrected route name
		router.POST("/bookingMeetingRoom", controller.CreateBookingMeetingRoom) // Assuming you want to create a meeting room
		router.PATCH("/bookingMeetingRoom", controller.UpdateBookingMeetingRoom)
		router.DELETE("/bookingMeetingRoom/:id", controller.DeleteBookingMeetingRoom)

	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server

	r.Run("localhost:" + PORT)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}