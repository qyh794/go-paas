package service

import (
	"github.com/qyh794/go-paas/middleware/domain/model"
	"github.com/qyh794/go-paas/middleware/domain/repository"
)

type IMiddlewareTypeDataService interface {
	AddMiddlewareType(*model.MiddleType) (int64, error)
	DeleteMiddlewareTypeByID(int64) error
	UpdateMiddlewareType(*model.MiddleType) error
	QueryMiddlewareTypeByID(int64) (*model.MiddleType, error)
	QueryAllMiddlewareType() ([]model.MiddleType, error)
	QueryImageVersionByID(int64) (string, error)
	QueryVersionByID(int64) (*model.MiddleVersion, error)
	QueryAllVersionByTypeID(int64) ([]model.MiddleVersion, error)
}

type MiddlewareTypeDataService struct {
	MiddlewareTypeRepository repository.IMiddleTypeRepository
}

func NewMiddlewareTypeDataService(middlewareTypeRepository repository.IMiddleTypeRepository) IMiddlewareTypeDataService {
	return &MiddlewareTypeDataService{
		MiddlewareTypeRepository: middlewareTypeRepository,
	}
}

func (m *MiddlewareTypeDataService) AddMiddlewareType(middleType *model.MiddleType) (int64, error) {
	return m.MiddlewareTypeRepository.CreateMiddleType(middleType)
}

func (m *MiddlewareTypeDataService) DeleteMiddlewareTypeByID(id int64) error {
	return m.MiddlewareTypeRepository.DeleteMiddleTypeByID(id)
}

func (m *MiddlewareTypeDataService) UpdateMiddlewareType(middleType *model.MiddleType) error {
	return m.MiddlewareTypeRepository.UpdateMiddleType(middleType)
}

func (m *MiddlewareTypeDataService) QueryMiddlewareTypeByID(id int64) (*model.MiddleType, error) {
	return m.MiddlewareTypeRepository.QueryTypeByID(id)
}

func (m *MiddlewareTypeDataService) QueryAllMiddlewareType() ([]model.MiddleType, error) {
	return m.MiddlewareTypeRepository.QueryAllMiddleType()
}

func (m *MiddlewareTypeDataService) QueryImageVersionByID(id int64) (string, error) {
	middlewareVersion, err := m.QueryVersionByID(id)
	if err != nil {
		return "", err
	}
	return middlewareVersion.MiddleDockerImage + ":" + middlewareVersion.MiddleDockerImageVersion, err
}

// QueryVersionByID 根据version ID查询单个镜像
func (m *MiddlewareTypeDataService) QueryVersionByID(id int64) (*model.MiddleVersion, error) {
	return m.MiddlewareTypeRepository.QueryVersionByID(id)
}

// QueryAllVersionByTypeID 根据中间件类型MiddleTypeID查找该中间件所有的版本
func (m *MiddlewareTypeDataService) QueryAllVersionByTypeID(middleTypeID int64) ([]model.MiddleVersion, error) {
	return m.MiddlewareTypeRepository.QueryAllVersionByTypeID(middleTypeID)
}
