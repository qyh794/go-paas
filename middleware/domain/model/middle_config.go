package model

// MiddleConfig 用于存储重要信息,例如初始化的账号\密码,在mysql\redis中间件中比较常见
type MiddleConfig struct {
	ID                       int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleID                 int64  `json:"middle_id"`                   // 关联的中间件ID
	MiddleConfigRootUser     string `json:"middle_config_root_user"`     // root用户名
	MiddleConfigRootPassword string `json:"middle_config_root_password"` // root密码
	MiddleConfigUser         string `json:"middle_config_user"`          // 普通用户名
	MiddleConfigPassword     string `json:"middle_config_password"`      // 普通用户密码
	MiddleConfigDatabase     string `json:"middle_config_database"`      // 使用数据库名
}
