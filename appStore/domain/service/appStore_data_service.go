package service

import (
	"appStore/domain/model"
	"appStore/domain/repository"
)

type IAppStoreDataService interface {
	// AddApp 通过ID新增一条App记录
	AddApp(*model.AppStore) (int64, error)

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

	// QueryAllCommentByID 查询app的所有评论
	QueryAllCommentByID(int64) ([]model.AppComment, error)
}

type AppStoreDataService struct {
	AppStoreRepository repository.IAppStoreRepository
}


func NewAppStoreDataService(appStoreRepository repository.IAppStoreRepository) IAppStoreDataService {
	return &AppStoreDataService{AppStoreRepository: appStoreRepository}
}

func (a *AppStoreDataService) AddApp(store *model.AppStore) (int64, error) {
	return a.AppStoreRepository.CreateApp(store)
}

func (a *AppStoreDataService) DeleteApp(id int64) error {
	return a.AppStoreRepository.DeleteApp(id)
}

func (a *AppStoreDataService) UpdateApp(store *model.AppStore) error {
	return a.AppStoreRepository.UpdateApp(store)
}

func (a *AppStoreDataService) QueryAppByID(id int64) (*model.AppStore, error) {
	return a.AppStoreRepository.QueryAppByID(id)
}

func (a *AppStoreDataService) QueryAllApp() ([]model.AppStore, error) {
	return a.AppStoreRepository.QueryAllApp()
}

func (a *AppStoreDataService) AddAppInstallNumByID(id int64) error {
	return a.AppStoreRepository.AddAppInstallNumByID(id)
}

func (a *AppStoreDataService) AddAppViewNumByID(id int64) error {
	return a.AppStoreRepository.AddAppViewNumByID(id)
}

func (a *AppStoreDataService) QueryAppInstallNumByID(id int64) int64 {
	return a.AppStoreRepository.QueryAppInstallNumByID(id)
}

func (a *AppStoreDataService) QueryAppViewNumByID(id int64) int64 {
	return a.AppStoreRepository.QueryAppViewNumByID(id)
}

func (a *AppStoreDataService) AddComment(comment *model.AppComment) error {
	return a.AppStoreRepository.AddComment(comment)
}

func (a *AppStoreDataService) QueryAllCommentByID(id int64) ([]model.AppComment, error) {
	return a.AppStoreRepository.QueryAllCommentByID(id)
}
