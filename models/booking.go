package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	Name        string `json:"name"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Requirement string `json:"requirement"`
	RoomID      uint   `json:"room_id"`
}
