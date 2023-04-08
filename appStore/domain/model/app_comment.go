package model

type AppComment struct {
	ID               int64  `gorm:"primary_key;auto_increment"` // 评论ID
	AppID            int64  `json:"app_id"`                     // 评论的App ID
	AppUserID        int64  `json:"app_user_id"`                // 发布评论者
	AppCommentTitle  string `json:"app_comment_title"`          // 评论标题
	AppCommentDetail string `json:"app_comment_detail"`         // 评论内容
}
