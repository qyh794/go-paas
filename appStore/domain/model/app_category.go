package model

type AppCategory struct {
	ID           int64  `gorm:"primary_key;auto_increment"`
	CategoryName string `json:"category_name"`
}
