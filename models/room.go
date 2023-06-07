package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Number      int    `json:"number"`
	Building    string `json:"building"`
	IsAvailable bool   `json:"is_available"`
}
