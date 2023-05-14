package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/route/proto/route"
	"github.com/qyh794/go-paas/routeApi/pkg/jwt"
	"github.com/qyh794/go-paas/routeApi/proto/routeApi"
	"strconv"
)

type RouteApi struct {
	RouteService route.RouteService
}

func (r *RouteApi) QueryRouteByID(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	if _, ok := request.Get["route_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 获取route id
	routeID, err := strconv.ParseInt(request.Get["route_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	routeInfo, err := r.RouteService.QueryRouteByID(ctx, &route.RRouteID{Id: routeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, routeInfo, response)
}

func (r *RouteApi) AddRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.AddRoute请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	routeInfo := &route.RRouteInfo{}
	err = json.Unmarshal([]byte(request.Body), routeInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	routeID, err := r.RouteService.AddRoute(ctx, routeInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, routeID, response)
}

func (r *RouteApi) DeleteRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	if _, ok := request.Get["route_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	routeID, err := strconv.ParseInt(request.Get["route_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	responseInfo, err := r.RouteService.DeleteRoute(ctx, &route.RRouteID{Id: routeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, responseInfo, response)
}

func (r *RouteApi) UpdateRoute(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}

	routeInfo := &route.RRouteInfo{}
	err = json.Unmarshal([]byte(request.Body), routeInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}

	routeID, err := r.RouteService.UpdateRoute(ctx, routeInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}

	return ResponseSucceed(ctx, routeID, response)
}

func (r *RouteApi) Call(ctx context.Context, request *routeApi.Request, response *routeApi.Response) error {
	log.Info("接收到routeApi.QueryRouteByID请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	allRoute, err := r.RouteService.QueryAllRoute(ctx, &route.RequestQueryAll{})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, allRoute, response)
}
