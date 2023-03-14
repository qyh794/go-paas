package repository

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/jinzhu/gorm"
	"svc/domain/model"
)

type ISvcRepository interface {
	InitTable() error
	QueryServiceByID(int64) (*model.Svc, error)
	CreateService(svc *model.Svc) (int64, error)
	DeleteServiceByID(int64) error
	UpdateService(svc *model.Svc) error
	QueryAllService() ([]model.Svc, error)
}

type SvcRepository struct {
	mysqlDB *gorm.DB
}

func NewSvcRepository(db *gorm.DB) ISvcRepository {
	return &SvcRepository{mysqlDB: db}
}

func (s *SvcRepository) InitTable() error {
	return s.mysqlDB.CreateTable(&model.Svc{}, &model.SvcPort{}).Error
}

func (s *SvcRepository) QueryServiceByID(id int64) (*model.Svc, error) {
	svcObj := &model.Svc{}
	// select * from Svc where svc.id = id
	return svcObj, s.mysqlDB.First(svcObj, id).Error
}

func (s *SvcRepository) CreateService(svc *model.Svc) (int64, error) {
	return svc.ID, s.mysqlDB.Create(svc).Error
}

func (s *SvcRepository) DeleteServiceByID(id int64) error {
	tx := s.mysqlDB.Begin()
	defer func() {
		if a := recover(); a != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		logger.Error(tx.Error)
		return tx.Error
	}
	// 多表删除先删从表
	// delete from svc where id = ?
	if err := s.mysqlDB.Where("id = ?", id).Delete(&model.Svc{}).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}
	// delete from svc.port where svc_id = ?
	if err := s.mysqlDB.Where("svc_id = ?", id).Delete(&model.SvcPort{}).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}
	return tx.Commit().Error
}

func (s *SvcRepository) UpdateService(svc *model.Svc) error {
	return s.mysqlDB.Model(svc).Update(svc).Error
}

func (s *SvcRepository) QueryAllService() (allSvc []model.Svc, err error) {
	return allSvc, s.mysqlDB.Find(&allSvc).Error
}
