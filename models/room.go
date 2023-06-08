package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Number   float32 `json:"number"`
	Building string  `json:"building"`
	Floor    int     `json:"floor"`
}
