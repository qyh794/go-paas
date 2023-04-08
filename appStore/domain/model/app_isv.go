package model

type AppIsv struct {
	ID           int64  `gorm:"primary_key;auto_increment"`
	AppIsvName   string `json:"app_isv_name"`
	AppIsvDetail string `json:"app_isv_detail"`
}
