package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/appStore/domain/model"
)

type IAppStoreRepository interface {
	// InitTable 初始化表
	InitTable() error

	// CreateApp 通过ID新增一条App记录
	CreateApp(*model.AppStore) (int64, error)

	// DeleteApp 通过ID删除一条App记录
	DeleteApp(int64) error

	// UpdateApp 更新App记录
	UpdateApp(*model.AppStore) error

	// QueryAppByID 通过ID查询App信息
	QueryAppByID(int64) (*model.AppStore, error)

	// QueryAllApp 查询AppStore中所有的App信息
	QueryAllApp() ([]model.AppStore, error)

	// AddAppInstallNumByID 通过ID增加某个App的安装量
	AddAppInstallNumByID(int64) error

	// AddAppViewNumByID 通过ID增加某个App的浏览量
	AddAppViewNumByID(int64) error

	// QueryAppInstallNumByID 通过ID查询某个App的安装量
	QueryAppInstallNumByID(int64) int64

	// QueryAppViewNumByID 通过ID查询某个App的浏览量
	QueryAppViewNumByID(int64) int64

	// AddComment 添加评论
	AddComment(*model.AppComment) error

	// DeleteComment 删除评论
	//DeleteComment(int64) error

	// QueryAllCommentByID 查询app的所有评论
	QueryAllCommentByID(int64) ([]model.AppComment, error)
}

type AppStoreRepository struct {
	mysqlDB *gorm.DB
}

func NewAppStoreRepository(db *gorm.DB) IAppStoreRepository {
	return &AppStoreRepository{mysqlDB: db}
}

func (a *AppStoreRepository) InitTable() error {
	return a.mysqlDB.CreateTable(&model.AppStore{}, &model.AppCategory{},
		&model.AppComment{}, &model.AppIsv{}, &model.AppVolume{},
		&model.AppPod{}, &model.AppMiddle{}).Error
}

func (a *AppStoreRepository) CreateApp(appStore *model.AppStore) (int64, error) {
	return appStore.ID, a.mysqlDB.Create(appStore).Error
}

func (a *AppStoreRepository) DeleteApp(id int64) error {
	tx := a.mysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	// 删除AppStore表
	if err := a.mysqlDB.Where("id = ?", id).Delete(&model.AppStore{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除AppMiddle表
	if err := a.mysqlDB.Where("app_id = ?", id).Delete(&model.AppMiddle{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除AppPod表
	if err := a.mysqlDB.Where("app_id = ?", id).Delete(&model.AppPod{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除AppIsv表
	if err := a.mysqlDB.Where("app_id = ?", id).Delete(&model.AppIsv{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除AppComment表
	if err := a.mysqlDB.Where("app_id = ?", id).Delete(&model.AppComment{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (a *AppStoreRepository) UpdateApp(appStore *model.AppStore) error {
	return a.mysqlDB.Model(&model.AppStore{}).Update(appStore).Error
}

func (a *AppStoreRepository) QueryAppByID(id int64) (*model.AppStore, error) {
	appStore := &model.AppStore{}
	return appStore, a.mysqlDB.Preload("AppPod").
		Preload("AppMiddle").Preload("AppVolume").
		Preload("AppComment").First(appStore, id).Error
}

func (a *AppStoreRepository) QueryAllApp() ([]model.AppStore, error) {
	appStoreArr := make([]model.AppStore, 0)
	// GORM 会自动根据结构体类型名称和 TableName 字段来确定要查询的表。在这个例子中，GORM 会自动查询 AppPod、
	// AppMiddle、AppVolume、AppComment、 表和 appStoreArr 表。
	return appStoreArr, a.mysqlDB.Preload("AppPod").
		Preload("AppMiddle").Preload("AppVolume").
		Preload("AppComment").Find(&appStoreArr).Error
}

func (a *AppStoreRepository) AddAppInstallNumByID(id int64) error {
	return a.mysqlDB.Model(&model.AppStore{}).Where("id = ?", id).UpdateColumn("app_install", gorm.Expr("app_install + ?", 1)).Error
}

func (a *AppStoreRepository) AddAppViewNumByID(id int64) error {
	return a.mysqlDB.Model(&model.AppStore{}).Where("id = ?", id).UpdateColumn("app_views", gorm.Expr("app_views + ?", 1)).Error
}

func (a *AppStoreRepository) QueryAppInstallNumByID(id int64) int64 {
	var installNum int64
	// select app_install from app_store where id = ?
	a.mysqlDB.Where("id = ?", id).Select("app_install").First(&model.AppStore{}).Scan(&installNum)
	return installNum
}

func (a *AppStoreRepository) QueryAppViewNumByID(id int64) int64 {
	var viewNum int64
	a.mysqlDB.Where("id = ?", id).Select("app_views").First(&model.AppStore{}).Scan(&viewNum)
	return viewNum
}

func (a *AppStoreRepository) AddComment(comment *model.AppComment) error {
	return a.mysqlDB.Create(comment).Error
}

//func (a *AppStoreRepository) DeleteComment(id int64) error {
//	return a.mysqlDB.Where("id = ?", id).Delete(&model.AppComment{}).Error
//}

func (a *AppStoreRepository) QueryAllCommentByID(id int64) ([]model.AppComment, error) {
	appComment := make([]model.AppComment, 0)
	return appComment, a.mysqlDB.Where("app_id = ?", id).Find(appComment).Error
}
