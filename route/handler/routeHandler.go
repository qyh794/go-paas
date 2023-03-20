package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/pod/common"
	"route/domain/model"
	"route/domain/service"
	"route/proto/route"
	"strconv"
)

type RouteHandler struct {
	RouteDateService service.IRouteDataService
}

func (r *RouteHandler) AddRoute(ctx context.Context, request *route.RRouteInfo, response *route.ResponseInfo) error {
	log.Info("接收到route.AddRoute请求")
	routeObj := &model.Route{}
	if err := common.SwapTo(request, routeObj); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	// 在k8s中创建ingress
	if err := r.RouteDateService.CreateRouteToK8s(request); err != nil {
		// 创建ingress失败
		logger.Error(err)
		response.Msg = err.Error()
		return err
	} else { // 创建ingress成功
		// 操作数据库
		routeID, err := r.RouteDateService.AddRoute(routeObj)
		// 插入数据库失败
		if err != nil {
			// @TODO 数据库插入失败,需要回滚,之前在k8s创建的资源需要撤销
			logger.Error(err)
			response.Msg = err.Error()
			return err
		}
		logger.Info("Route ID: " + strconv.FormatInt(routeID, 10))
		response.Msg = "Route ID: " + strconv.FormatInt(routeID, 10)
	}
	return nil
}

func (r *RouteHandler) DeleteRoute(ctx context.Context, request *route.RRouteID, response *route.ResponseInfo) error {
	log.Info("接收到route.DeleteRoute请求")
	routeObj, err := r.RouteDateService.QueryRouteByID(request.Id)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	// 通过资源名称删除k8s中的资源
	if err = r.RouteDateService.DeleteRouteFromK8s(routeObj); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	//@TODO 数据一致性问题, 需要解决使用校验或者分布式事务
	// 通过id删除数据库中对应的数据
	if err = r.RouteDateService.DeleteRoute(routeObj.ID); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	return nil
}

func (r *RouteHandler) UpdateRoute(ctx context.Context, request *route.RRouteInfo, response *route.ResponseInfo) error {
	log.Info("接收到route.UpdateRoute请求")

	// 先更新k8s中的资源
	if err := r.RouteDateService.UpdateRouteToK8s(request); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	routeObj := &model.Route{}
	if err := common.SwapTo(request, routeObj); err != nil {
		logger.Error(err)
		response.Msg = "数据有误" + err.Error()
		return err
	}
	// 更新k8s成功,操作数据库
	if err := r.RouteDateService.UpdateRoute(routeObj); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	return nil
}

func (r *RouteHandler) QueryRouteByID(ctx context.Context, request *route.RRouteID, response *route.RRouteInfo) error {
	log.Info("接收到route.QueryRouteByID请求")
	routeObj, err := r.RouteDateService.QueryRouteByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err = common.SwapTo(routeObj, response); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (r *RouteHandler) QueryAllRoute(ctx context.Context, request *route.RequestQueryAll, response *route.ResponseAllRoute) error {
	log.Info("接收到route.QueryAllRoute请求")
	allRoute, err := r.RouteDateService.QueryAllRoute()
	if err != nil {
		logger.Error(err)
		return err
	}
	for i := 0; i < len(allRoute); i++ {
		routeInfo := &route.RRouteInfo{}
		if err = common.SwapTo(allRoute[i], routeInfo); err != nil {
			logger.Error(err)
			return err
		}
		response.RouteInfo = append(response.RouteInfo, routeInfo)
	}
	return nil
}
