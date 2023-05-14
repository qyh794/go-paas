package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"github.com/qyh794/go-paas/middlewareApi/pkg/jwt"
	"github.com/qyh794/go-paas/middlewareApi/proto/middlewareApi"
	"strconv"
)

type MiddlewareApi struct {
	MiddlewareService middleware.MiddlewareService
}

func (m *MiddlewareApi) AddMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.AddMiddleware 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	middlewareInfo := &middleware.RMiddlewareInfo{}
	// 从请求表单中获取middle_port的所有值
	err = json.Unmarshal([]byte(request.Body), middlewareInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)

	}
	responseInfo, err := m.MiddlewareService.AddMiddleware(ctx, middlewareInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)

	}
	return ResponseSucceed(ctx, responseInfo, response)

}

func (m *MiddlewareApi) DeleteMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.DeleteMiddleware 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	id, err := strconv.ParseInt(request.Get["middle_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	responseInfo, err := m.MiddlewareService.DeleteMiddlewareByID(ctx, &middleware.RequestMiddlewareID{Id: id})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, responseInfo, response)
}

func (m *MiddlewareApi) UpdateMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.UpdateMiddleware 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.创建服务层需要的对象
	middlewareInfo := &middleware.RMiddlewareInfo{}
	err = json.Unmarshal([]byte(request.Body), middlewareInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}

	// 2.调用后端服务
	responseInfo, err := m.MiddlewareService.UpdateMiddleware(ctx, middlewareInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}

	return ResponseSucceed(ctx, responseInfo, response)
}

func (m *MiddlewareApi) QueryMiddlewareByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryMiddlewareByID 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从请求url中获取middlewareID
	if _, ok := request.Get["middle_id"]; !ok {
		// 请求中不存在数据
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	middlewareID, err := strconv.ParseInt(request.Get["middle_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	middlewareInfo, err := m.MiddlewareService.QueryMiddlewareByID(ctx, &middleware.RequestMiddlewareID{Id: middlewareID})
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	return ResponseSucceed(ctx, middlewareInfo, response)
}

func (m *MiddlewareApi) QueryAllMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryAllMiddleware 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.调用服务层接口
	allMiddleware, err := m.MiddlewareService.QueryAllMiddleware(ctx, &middleware.RequestAll{})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, allMiddleware, response)
}

func (m *MiddlewareApi) QueryAllMiddlewareByTypeID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryAllMiddlewareByTypeID 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 0.判断请求中是否存在数据
	if _, ok := request.Get["type_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 1.从请求url中获取typeID
	// 2.将typeID转换为int64
	typeID, err := strconv.ParseInt(request.Get["type_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 3.调用服务层接口
	allMiddlewareType, err := m.MiddlewareService.QueryAllMiddlewareByTypeID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, allMiddlewareType, response)
}

func (m *MiddlewareApi) AddMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.AddMiddlewareType 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}

	middleTypeInfo := &middleware.RMiddleTypeInfo{}

	if err = json.Unmarshal([]byte(request.Body), middleTypeInfo); err != nil {
		return ResponseBadRequest(ctx, err, response)
	}

	responseInfo, err := m.MiddlewareService.AddMiddlewareType(ctx, middleTypeInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, responseInfo, response)
}

func (m *MiddlewareApi) DeleteMiddlewareTypeByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	// 1.判断请求中是否存在type_id
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	if _, ok := request.Get["type_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.获取请求中的type_id
	typeID, err := strconv.ParseInt(request.Get["type"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 3.调用服务层接口
	responseInfo, err := m.MiddlewareService.DeleteMiddleTypeByID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 4.返回结果
	return ResponseSucceed(ctx, responseInfo, response)
}

func (m *MiddlewareApi) UpdateMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.创建服务层对象
	typeInfo := &middleware.RMiddleTypeInfo{}
	// 2.将请求体中的json数据反序列化到结构体中
	err = json.Unmarshal([]byte(request.Body), typeInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	// 3.调用服务层接口
	responseInfo, err := m.MiddlewareService.UpdateMiddleType(ctx, typeInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 4.返回相应
	return ResponseSucceed(ctx, responseInfo, response)
}

func (m *MiddlewareApi) QueryMiddlewareTypeByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.判断请求中是否存在type_id
	if _, ok := request.Get["type_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.获取请求中的type_id
	typeID, err := strconv.ParseInt(request.Get["type_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 3.调用服务层接口
	typeInfo, err := m.MiddlewareService.QueryMiddleTypeByID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 4.返回结果
	return ResponseSucceed(ctx, typeInfo, response)
}

func (m *MiddlewareApi) QueryAllMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.调用服务层接口
	middleType, err := m.MiddlewareService.QueryAllMiddleType(ctx, &middleware.RequestAll{})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 2.返回结果
	return ResponseSucceed(ctx, middleType, response)
}
