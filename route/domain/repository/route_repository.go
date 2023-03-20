package repository

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/route/domain/model"
)

type IRouteRepository interface {
	InitTable() error
	QueryRouteByID(int64) (*model.Route, error)
	CreateRoute(*model.Route) (int64, error)
	DeleteRouteByID(int64) error
	UpdateRoute(*model.Route) error
	QueryAllRoute() ([]model.Route, error)
}

type RouteRepository struct {
	mysqlDB *gorm.DB
}

func NewRouteRepository(db *gorm.DB) IRouteRepository {
	return &RouteRepository{mysqlDB: db}
}

func (r *RouteRepository) InitTable() error {
	return r.mysqlDB.CreateTable(&model.Route{}, &model.RoutePath{}).Error
}

func (r *RouteRepository) QueryRouteByID(id int64) (*model.Route, error) {
	route := &model.Route{}
	return route, r.mysqlDB.Preload("RoutePath").First(route, id).Error
}

func (r *RouteRepository) CreateRoute(route *model.Route) (int64, error) {
	return route.ID, r.mysqlDB.Create(route).Error
}

func (r *RouteRepository) DeleteRouteByID(id int64) error {
	tx := r.mysqlDB.Begin()
	defer func() {
		if a := recover(); a != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		logger.Error(tx.Error)
	}
	// 先删从表
	if err := r.mysqlDB.Where("id = ?", id).Delete(&model.Route{}).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}
	if err := r.mysqlDB.Where("route_id = ?", id).Delete(&model.RoutePath{}).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}
	return tx.Commit().Error
}

func (r *RouteRepository) UpdateRoute(route *model.Route) error {
	// ...
	return r.mysqlDB.Model(&model.Route{}).Update(route).Error
}

func (r *RouteRepository) QueryAllRoute() ([]model.Route, error) {
	var allRoute []model.Route
	return allRoute, r.mysqlDB.Preload("RoutePath").Find(&allRoute).Error
}
