package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Name        string         `json:"name"`
	Date        datatypes.Date `json:"date"`
	Time        string         `json:"time"`
	Requirement string         `json:"requirement"`
	RoomID      uint           `json:"room_id"`
	Room        Room           `json:"room"`
}
