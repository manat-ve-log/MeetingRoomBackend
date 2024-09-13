package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"meeting/config"
	"meeting/entity"
)

// GET /meetingroom
func ListCustomer(c *gin.Context) {
	var customer []entity.Customer

	db := config.DB()

	db.Find(&customer)

	c.JSON(http.StatusOK, &customer)
}

func GetCustomer(c*gin.Context){
	ID := c.Param("id")
	var customer entity.Customer
	db := config.DB()
	results := db.Preload("customer").First(&customer, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if customer.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	var room entity.Customer

	CustomerID := c.Param("id")

	db := config.DB()
	result := db.First(&room, CustomerID)
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

// GET /CreateCutomer
func CreateCustomer(c *gin.Context) {
	var user entity.Customer

	// Bind JSON to the user variable
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// สร้าง Customer
	customer := entity.Customer{
		First_Name: user.First_Name, // ตั้งค่าฟิลด์ FirstName
		Last_Name:  user.Last_Name,  // ตั้งค่าฟิลด์ LastName
		Email:     user.Email,     // ตั้งค่าฟิลด์ Email
		Tel:       user.Tel,       // ตั้งค่าฟิลด์ Tel
	}

	// บันทึก
	if err := db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": customer})
}

func DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

