package model

type AppVolume struct {
	ID          int64 `gorm:"primary_key;auto_increment"`
	AppID       int64 `json:"app_id"`
	AppVolumeID int64 `json:"app_volume_id"`
}
