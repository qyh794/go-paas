package model

type Volume struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	// 存储名
	VolumeName string `json:"volume_name"`
	// 存储所属命名空间
	VolumeNamespace string `json:"volume_namespace"`
	// 存储的访问方式
	VolumeAccessMode       string `json:"volume_access_mode"`
	VolumeStorageClassName string `json:"volume_storage_class_name"`
	// 请求资源大小
	VolumeRequest float32 `json:"volume_request"`
	// 存储类型
	VolumePersistentVolumeMode string `json:"volume_persistent_volume_mode"`
}
