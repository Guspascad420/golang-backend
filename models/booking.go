package models

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	Requirement string    `json:"requirement"`
	RoomID      uint      `json:"room_id"`
	Room        Room      `json:"room"`
}
