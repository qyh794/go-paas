package model

type AppPod struct {
	ID       int64 `gorm:"primary_key;auto_increment"`
	AppID    int64 `json:"app_id"`
	AppPodID int64 `json:"app_pod_id"`
}
