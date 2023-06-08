package models

type Booking struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	RoomID uint   `json:"room_id"`
	Room   Room   `json:"room"`
}
