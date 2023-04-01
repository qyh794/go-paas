package model

type Middleware struct {
	ID              int64           `gorm:"primary_key;not_null;auto_increment"`
	MiddleTypeID    int64           `json:"middle_type_id"`
	MiddleVersionID int64           `json:"middle_version_id"`
	MiddleReplicas  int32           `json:"middle_replicas"` // 有多个副本会复制多份配置
	MiddleName      string          `json:"middle_name"`
	MiddleNamespace string          `json:"middle_namespace"`
	MiddleCpu       float32         `json:"middle_cpu"`
	MiddleMemory    float32         `json:"middle_memory"`
	MiddleConfig    MiddleConfig    `gorm:"foreign_key:MiddleID" json:"middle_config"`
	MiddlePort      []MiddlePort    `gorm:"foreign_key:MiddleID" json:"middle_port"`
	MiddleEnv       []MiddleEnv     `gorm:"foreign_key:MiddleID" json:"middle_env"`
	MiddleStorage   []MiddleStorage `gorm:"foreign_key:MiddleID" json:"middle_storage"`
}
