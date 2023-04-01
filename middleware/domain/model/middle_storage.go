package model

// MiddleStorage 中间件存储盘
type MiddleStorage struct {
	ID                      int64   `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleID                int64   `json:"middle_id"` // 关联中间件ID
	MiddleStorageSize       float32 `json:"middle_storage_size"`
	MiddleStorageName       string  `json:"middle_storage_name"`  // 存储名称
	MiddleStoragePath       string  `json:"middle_storage_path"`  // 存储创建后需要挂载到中间件的目录
	MiddleStorageClass      string  `json:"middle_storage_class"` // 存储的类型,例如固态还是机械硬盘
	MiddleStorageAccessMode string  `json:"middle_storage_access_mode"`
}
