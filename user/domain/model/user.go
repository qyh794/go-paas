package model

type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique_index" json:"username"`
	Password string `json:"password"`
}
