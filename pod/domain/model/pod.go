package model

type Pod struct {
	ID           int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	PodName      string `gorm:"unique_index;not_null" json:"pod_name"`
	PodNamespace string `json:"pod_namespace"`
	// pod所属团队
	PodTeamID string `json:"pod_team_id"`
	// pod 使用cpu最小值
	PodCpuMin float32 `json:"pod_cpu_min"`
	PodCpuMax float32 `json:"pod_cpu_max"`
	// 副本数量
	PodReplicas int32 `json:"pod_replicas"`
	// pod 使用内存最小值
	PodMemoryMin float32 `json:"pod_memory_min"`
	PodMemoryMax float32 `json:"pod_memory_max"`
	// pod 开放端口
	PodPort []PodPort `gorm:"ForeignKey:PodID" json:"pod_port"`
	// pod 使用的环境变量
	PodEnv []PodEnv `gorm:"ForeignKey:PodID" json:"pod_env"`
	// 镜像拉取策略
	PodPullPolicy string `json:"pod_pull_policy"`
	// 重启策略
	PodRestart string `json:"pod_restart"`
	// pod 发布策略
	PodType string `json:"pod_type"`
	// pod 使用镜像
	PodImage string `json:"pod_image"`
	// @TODO 挂盘
	// @TODO 域名设置
}
