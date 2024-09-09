package controller

import (
	"meeting/config"
	"meeting/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /ManageRoom
func CreateBooking(c *gin.Context) {
	var manageRoom entity.ManageRoom

	// bind เข้าตัวแปร manageRoom
	if err := c.ShouldBindJSON(&manageRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()
	// ค้นหา Customer ด้วย id
	var customer entity.Customer
	db.First(&customer, manageRoom.CustomerID)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "gender not found"})
		return
	}

	var meetingRoom entity.MeetingRoom
	db.First(&meetingRoom, manageRoom.MeetingRoomID)
	if meetingRoom.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "gender not found"})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล

	// สร้าง User
	b := entity.ManageRoom{
		StartTime:     manageRoom.StartTime,  // ตั้งค่าฟิลด์ StartTime
		EndTime:       manageRoom.EndTime,    // ตั้งค่าฟิลด์ EndTime
		CustomerID: manageRoom.CustomerID,
		Customer:    customer,    // ตั้งค่าฟิลด์ CustomerID
		MeetingRoomID: manageRoom.MeetingRoomID, // ตั้งค่าฟิลด์ MeetingRoomID
		MeetingRoom:   meetingRoom, 
	}

	// บันทึก
	if err := db.Create(&b).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": b})
}


// GET /Booking/:id
func GetBooking(c *gin.Context) {
	ID := c.Param("id")
	var booking entity.ManageRoom

	db := config.DB()
	results := db.Preload("Gender").First(&booking, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if booking.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// GET /Bookings
func ListBooking(c *gin.Context) {

	var bookings []entity.ManageRoom

	db := config.DB()
	results := db.Preload("Gender").Find(&bookings)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}


// DELETE /Booking/:id
func DeleteBooking(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM ManagRoom WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

// PATCH /Booking
func UpdateBooking(c *gin.Context) {
	var booking entity.ManageRoom

	BookingID := c.Param("id")

	db := config.DB()
	result := db.First(&booking, BookingID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&booking)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}