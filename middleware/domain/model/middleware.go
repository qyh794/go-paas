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

/*
type Middleware struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	//中间件的名称
	MiddleName string `json:"middle_name"`
	//中间件创建的命名空间
	MiddleNamespace string `json:"middle_namespace"`
	//中间件的类型
	MiddleTypeID int64 `json:"middle_type_id"`
	//中间件的版本
	MiddleVersionID int64 `json:"middle_version_id"`
	//中间件的端口
	MiddlePort[] MiddlePort `gorm:"ForeignKey:MiddleID" json:"middle_port"`
	//默认生成的账号密码
	MiddleConfig MiddleConfig `gorm:"ForeignKey:MiddleID" json:"middle_config"`
	//环境变量
	MiddleEnv[] MiddleEnv `gorm:"ForeignKey:MiddleID" json:"middle_env"`
	//中间件的CPU
	MiddleCpu float32 `json:"middle_cpu"`
	//中间件内存
	MiddleMemory float32 `json:"middle_memory"`
	//中间件存储
	MiddleStorage[] MiddleStorage `gorm:"ForeignKey:MiddleID" json:"middle_storage"`
	//中间件副本
	MiddleReplicas int32 `json:"middle_replicas"`
}
*/
