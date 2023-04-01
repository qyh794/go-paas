package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/middleware/domain/model"
)

type IMiddlewareRepository interface {
	InitTable() error
	QueryMiddlewareByID(int64) (*model.Middleware, error)
	CreateMiddleware(*model.Middleware) (int64, error)
	DeleteMiddlewareByID(int64) error
	UpdateMiddleware(*model.Middleware) error
	QueryAllMiddleware() ([]model.Middleware, error)
	QueryAllMiddlewareByType(int64) ([]model.Middleware, error)
}

type MiddlewareRepository struct {
	mysqlDB *gorm.DB
}

func NewMiddlewareRepository(db *gorm.DB) IMiddlewareRepository {
	return &MiddlewareRepository{mysqlDB: db}
}

func (m *MiddlewareRepository) InitTable() error {
	return m.mysqlDB.CreateTable(&model.Middleware{}, &model.MiddleEnv{},
		&model.MiddleConfig{}, &model.MiddleStorage{}, &model.MiddlePort{}).Error
}

func (m *MiddlewareRepository) QueryMiddlewareByID(id int64) (*model.Middleware, error) {
	middleware := &model.Middleware{}
	return middleware, m.mysqlDB.
		Preload("middle_config").
		Preload("middle_port").
		Preload("middle_env").
		Preload("middle_storage").
		First(middleware, id).Error
}

func (m *MiddlewareRepository) CreateMiddleware(middleware *model.Middleware) (int64, error) {
	return middleware.ID, m.mysqlDB.Create(middleware).Error
}

func (m *MiddlewareRepository) DeleteMiddlewareByID(id int64) error {
	tx := m.mysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	if err := m.mysqlDB.Where("id = ?", id).Delete(&model.Middleware{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := m.mysqlDB.Where("middle_id = ?", id).Delete(&model.MiddleConfig{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := m.mysqlDB.Where("middle_id = ?", id).Delete(&model.MiddlePort{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := m.mysqlDB.Where("middle_id = ?", id).Delete(&model.MiddleEnv{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := m.mysqlDB.Where("middle_id = ?", id).Delete(&model.MiddleStorage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *MiddlewareRepository) UpdateMiddleware(middleware *model.Middleware) error {
	return m.mysqlDB.Model(&model.Middleware{}).Update(middleware).Error
}

func (m *MiddlewareRepository) QueryAllMiddleware() ([]model.Middleware, error) {
	var middlewares []model.Middleware
	return middlewares, m.mysqlDB.
		Model(&model.Middleware{}).
		Preload("middle_config").
		Preload("middle_port").
		Preload("middle_env").
		Preload("middle_storage").
		Find(&middlewares).Error
}

func (m *MiddlewareRepository) QueryAllMiddlewareByType(id int64) ([]model.Middleware, error) {
	var middlewares []model.Middleware
	return middlewares, m.mysqlDB.
		Preload("middle_config").
		Preload("middle_port").
		Preload("middle_env").
		Preload("middle_storage").
		Where("middle_type_id = ?", id).
		Find(&middlewares).Error

}
