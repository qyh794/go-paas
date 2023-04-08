package model

type AppStore struct {
	ID            int64        `gorm:"primary_key;auto_increment"`
	AppInstall    int64        `json:"app_install"`
	AppViews      int64        `json:"app_views"`
	AppCategoryID int64        `json:"app_category_id"`
	AppIsvID      int64        `json:"app_isv_id"` //还可以包含AppIsvID服务商的公司信息资质等字段
	AppPrice      float32      `json:"app_price"`
	AppSku        string       `gorm:"unique_index;not_null" json:"app_sku"`
	AppTitle      string       `json:"app_title"`
	AppDetail     string       `json:"app_detail"`
	AppCheck      bool         `json:"app_check"`
	AppPod        []AppPod     `gorm:"foreignKey:AppID" json:"app_pod"`
	AppMiddle     []AppMiddle  `gorm:"foreignKey:AppID" json:"app_middle"`
	AppVolume     []AppVolume  `gorm:"foreignKey:AppID" json:"app_volume"`
	AppComment    []AppComment `gorm:"foreignKey:AppID" json:"app_comment"`
}
