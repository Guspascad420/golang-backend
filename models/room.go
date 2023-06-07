package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Number      float32 `json:"number"`
	Building    string  `json:"building"`
	IsAvailable bool    `json:"is_available"`
	Floor       int     `json:"floor"`
}
