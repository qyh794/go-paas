package model

type AppMiddle struct {
	ID          int64 `gorm:"primary_key;auto_increment"`
	AppID       int64 `json:"app_id"`
	AppMiddleID int64 `json:"app_middle_id"`
}
