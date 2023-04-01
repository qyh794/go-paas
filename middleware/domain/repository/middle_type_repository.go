package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/middleware/domain/model"
)

type IMiddleTypeRepository interface {
	InitTable() error
	QueryTypeByID(int64) (*model.MiddleType, error)
	CreateMiddleType(*model.MiddleType) (int64, error)
	DeleteMiddleTypeByID(int64) error
	UpdateMiddleType(*model.MiddleType) error
	QueryAllMiddleType() ([]model.MiddleType, error)
	QueryVersionByID(int64) (*model.MiddleVersion, error)
	// QueryAllVersionByTypeID 查询中间件的所有镜像版本
	QueryAllVersionByTypeID(int64) ([]model.MiddleVersion, error)
}

type MiddleTypeRepository struct {
	mysqlDB *gorm.DB
}

func NewMiddleTypeRepository(db *gorm.DB) IMiddleTypeRepository {
	return &MiddleTypeRepository{mysqlDB: db}
}

func (m *MiddleTypeRepository) InitTable() error {
	return m.mysqlDB.CreateTable(&model.MiddleType{}, &model.MiddleVersion{}).Error
}

func (m *MiddleTypeRepository) QueryTypeByID(middleTypeID int64) (*model.MiddleType, error) {
	middleType := &model.MiddleType{}
	return middleType, m.mysqlDB.Preload("middle_version").First(middleType, middleTypeID).Error
}

func (m *MiddleTypeRepository) CreateMiddleType(middleType *model.MiddleType) (int64, error) {
	return middleType.ID, m.mysqlDB.Create(middleType).Error
}

func (m *MiddleTypeRepository) DeleteMiddleTypeByID(middleTypeID int64) error {
	tx := m.mysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	if err := m.mysqlDB.Where("id = ?", middleTypeID).Delete(&model.MiddleType{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := m.mysqlDB.Where("middle_type_id = ?", middleTypeID).Delete(&model.MiddleVersion{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Rollback().Error
}

func (m *MiddleTypeRepository) UpdateMiddleType(middleType *model.MiddleType) error {
	return m.mysqlDB.Model(&model.MiddleType{}).Update(middleType).Error
}

func (m *MiddleTypeRepository) QueryAllMiddleType() ([]model.MiddleType, error) {
	var middleTypes []model.MiddleType
	return middleTypes, m.mysqlDB.Preload("middle_version").Find(&middleTypes).Error
}

func (m *MiddleTypeRepository) QueryVersionByID(middleTypeID int64) (*model.MiddleVersion, error) {
	var middleVersion *model.MiddleVersion
	return middleVersion, m.mysqlDB.First(middleVersion, middleTypeID).Error
}

func (m *MiddleTypeRepository) QueryAllVersionByTypeID(middleTypeID int64) ([]model.MiddleVersion, error) {
	var middleVersions []model.MiddleVersion
	return middleVersions, m.mysqlDB.Where("middle_type_id = ?", middleTypeID).Find(&middleVersions).Error
}
