package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/micro/micro/v3/service/logger"
	"github.com/qyh794/go-paas/route/proto/route"
	"github.com/qyh794/go-paas/routeApi/pkg/jwt"
	"github.com/qyh794/go-paas/routeApi/proto/routeApi"
	"strconv"
	"strings"
)

type RouteApi struct {
	RouteService route.RouteService
}

const (
	Succeed     = 200
	WrongArgs   = 400
	ServiceBusy = 500
)

func (r *RouteApi) QueryRouteByID(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	// @TODO request.Get["route_id"]能否获取到参数
	if _, ok := request.Get["route_id"]; !ok {
		response.StatusCode = WrongArgs
		return errors.New("参数异常")
	}
	// 获取route id
	routeID, err := strconv.ParseInt(request.Get["route_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	routeInfo, err := r.RouteService.QueryRouteByID(ctx, &route.RRouteID{Id: routeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(routeInfo)
	response.Body = string(bytes)
	return nil
}

func (r *RouteApi) AddRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.AddRoute请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	routeInfo := &route.RRouteInfo{}
	err = json.Unmarshal([]byte(request.Body), routeInfo)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
	routeID, err := r.RouteService.AddRoute(ctx, routeInfo)
	if err != nil {
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(routeID)
	response.Body = string(bytes)
	return nil
}

func (r *RouteApi) DeleteRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	if _, ok := request.Get["route_id"]; !ok {
		response.StatusCode = WrongArgs
		return errors.New("参数异常")
	}
	routeID, err := strconv.ParseInt(request.Get["route_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	responseInfo, err := r.RouteService.DeleteRoute(ctx, &route.RRouteID{Id: routeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

func (r *RouteApi) UpdateRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	routeInfo := &route.RRouteInfo{}
	err = json.Unmarshal([]byte(request.Body), routeInfo)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
	routeID, err := r.RouteService.UpdateRoute(ctx, routeInfo)
	if err != nil {
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(routeID)
	response.Body = string(bytes)
	return nil
}

func (r *RouteApi) Call(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	allRoute, err := r.RouteService.QueryAllRoute(ctx, &route.RequestQueryAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = 200
	bytes, _ := json.Marshal(allRoute)
	response.Body = string(bytes)
	return nil
}
