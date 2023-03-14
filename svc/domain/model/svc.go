package model

type Svc struct {
	ID              int64     `gorm:"primary_key;not_null;auto_increment"`
	SvcName         string    `gorm:"unique_index;not_null" json:"svc_name"`
	SvcNamespace    string    `gorm:"not_null" json:"svc_namespace"`
	SvcPodName      string    `gorm:"not_null" json:"svc_pod_name"`
	SvcType         string    `json:"svc_type"`
	SvcExternalName string    `json:"svc_external_name"`
	SvcTeamID       string    `json:"svc_team_id"`
	SvcPort         []SvcPort `gorm:"foreignKey:SvcID" json:"svc_port"`
}
