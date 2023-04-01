package model

type MiddleEnv struct {
	ID       int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleID int64  `json:"middle_id"` // 关联中间件ID
	EnvKey   string `json:"env_key"`
	EnvValue string `json:"env_value"`
}
