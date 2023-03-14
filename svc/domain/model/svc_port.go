package model

type SvcPort struct {
	ID              int64  `gorm:"primary_key;not_null;auto_increment"`
	SvcID           int64  `json:"svc_id"`
	SvcPort         int32  `json:"svc_port"`        // service端口
	SvcTargetPort   int32  `json:"svc_target_port"` // pod端口
	SvcNodePort     int32  `json:"svc_node_port"`   // 主机端口
	SvcPortProtocol string `json:"svc_port_protocol"`
}
