package models

import (
	"gorm.io/datatypes"
)

type Booking struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `json:"name"`
	Date        datatypes.Date `json:"date"`
	Time        string         `json:"time"`
	Requirement string         `json:"requirement"`
	RoomID      uint           `json:"room_id"`
	UserID      uint           `json:"user_id"`
	Room        Room           `json:"room"`
}
