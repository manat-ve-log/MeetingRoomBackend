package config

import (
	"fmt"
	// "time"

	"meeting/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.MeetingRoom{},
		&entity.CustomerMeetingRoom{},
		&entity.ManageRoom{},
	)

	// GenderMale := entity.Gender{Name: "Male"}
	// GenderFemale := entity.Gender{Name: "Female"}

	// db.FirstOrCreate(&GenderMale, &entity.Gender{Name: "Male"})
	// db.FirstOrCreate(&GenderFemale, &entity.Gender{Name: "Female"})

	// hashedPassword, _ := HashPassword("123456")
	// BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
	room1 := &entity.MeetingRoom{
		RoomName:     "Conference Room A",
		Capacity:     20,
		Detail:       "A medium-sized conference room with a projector and whiteboard.",
		RoomSize:     50.5,
		AirCondition: 1,
		Chair:        20,
		Type:         "Conference",
	}
	db.FirstOrCreate(&room1, entity.MeetingRoom{
		RoomName: "Conference Room A", // condition to find the existing record
	})

}