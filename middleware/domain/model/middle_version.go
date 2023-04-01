package model

// MiddleVersion 镜像版本
type MiddleVersion struct {
	ID                       int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	MiddleTypeID             int64  `json:"middle_type_id"`
	MiddleDockerImage        string `json:"middle_docker_image"`
	MiddleDockerImageVersion string `json:"middle_docker_image_version"`
	// 通过MiddleDockerImage:MiddleDockerImageVersion访问镜像
}
