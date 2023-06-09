package models

type LoginResponse struct {
	Token string `json:"token"`
}

type UserDataResponse struct {
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}
