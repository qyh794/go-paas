package model

// MiddleType 中间件类型
type MiddleType struct {
	ID             int64           `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleTypeName string          `json:"middle_type_name"`
	MiddleVersion  []MiddleVersion `gorm:"foreign_key:MiddleTypeID" json:"middle_version"` // 中间件版本,同一个中间件有多个版本
	// 这里的外键是MiddleVersion表中的MiddleTypeID,关联到了MiddleType表中的主键
}
