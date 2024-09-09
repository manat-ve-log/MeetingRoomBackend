package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"meeting/config"
	"meeting/entity"
)

// GET /MeetingRoom
func ListMeetingRoom(c *gin.Context) {
	var meetingroom []entity.MeetingRoom

	db := config.DB()

	db.Find(&meetingroom)

	c.JSON(http.StatusOK, &meetingroom)
}
// GET /MeetingRoom/:id
func GetMeetingRoom(c *gin.Context) {
	ID := c.Param("id")
	var room entity.MeetingRoom

	db := config.DB()
	results := db.Preload("Gender").First(&room, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if room.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, room)
}

//POST /UpdateMeetingRoom
func UpdateMeetingRoom(c *gin.Context) {
	var room entity.MeetingRoom

	MeetingRoomID := c.Param("id")

	db := config.DB()
	result := db.First(&room, MeetingRoomID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&room)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

// POST /users
func CreateMeetingRoom(c *gin.Context) {
	var room entity.MeetingRoom

	// bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ค้นหา gender ด้วย id
	
	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล

	// สร้าง User
	r := entity.MeetingRoom{
		RoomName:     room.RoomName,
		Capacity:     room.Capacity,
		Detail:       room.Detail,
		RoomSize:     room.RoomSize,
		AirCondition: room.AirCondition,
		Chair:        room.Chair,
		Type:         room.Type,
	}

	// บันทึก
	if err := db.Create(&r).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": r})
}

func DeleteMeetingRoom(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM ManagRoom WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}