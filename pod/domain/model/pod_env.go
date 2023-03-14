package model

type PodEnv struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 通过PodID进行关联, PodID与POD表中的ID(主键)一致
	PodID int64 `json:"pod_id"`
	// 环境变量key
	EnvKey   string `json:"env_key"`
	EnvValue string `json:"env_value"`
}
