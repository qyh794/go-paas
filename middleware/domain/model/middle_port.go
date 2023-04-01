package model

// MiddlePort 用于表示端口关联中间件的ID、开放的端口协议
// 中间件开放的端口,可能有多个,如果放到一个结构体里面不方便管理
type MiddlePort struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleID       int64  `json:"middle_id"` // 关联中间件的ID
	MiddlePort     int32  `json:"middle_port"`
	MiddleProtocol string `json:"middle_protocol"` // 中间件开放的端口协议
}
