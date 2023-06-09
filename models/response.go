package models

type LoginResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}
