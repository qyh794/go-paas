package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"github.com/qyh794/go-paas/middlewareApi/proto/middlewareApi"
	"net/http"
	"strconv"
)

type MiddlewareApi struct {
	MiddlewareService middleware.MiddlewareService
}

func (m *MiddlewareApi) AddMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.AddMiddleware 请求")
	middlewareInfo := &middleware.RMiddlewareInfo{}
	// 从请求表单中获取middle_port的所有值
	err := json.Unmarshal([]byte(request.Body), middlewareInfo)
	if err != nil {
		logger.Error(err)
		response.Body = errors.New("参数有误").Error()
		return err
	}

	responseInfo, err := m.MiddlewareService.AddMiddleware(ctx, middlewareInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) DeleteMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.DeleteMiddleware 请求")
	id, err := strconv.ParseInt(request.Get["middle_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	responseInfo, err := m.MiddlewareService.DeleteMiddlewareByID(ctx, &middleware.RequestMiddlewareID{Id: id})
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	response.StatusCode = 200
	marshal, _ := json.Marshal(responseInfo)
	response.Body = string(marshal)
	return nil
}

func (m *MiddlewareApi) UpdateMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middlewareApi.UpdateMiddleware 请求")
	// 1.创建服务层需要的对象
	middlewareInfo := &middleware.RMiddlewareInfo{}
	err := json.Unmarshal([]byte(request.Body), middlewareInfo)
	if err != nil {
		logger.Error(err)
		response.Body = errors.New("参数有误").Error()
		return err
	}

	// 2.调用后端服务
	responseInfo, err := m.MiddlewareService.UpdateMiddleware(ctx, middlewareInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 3.返回结果
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) QueryMiddlewareByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryMiddlewareByID 请求")
	// 1.从请求url中获取middlewareID
	if _, ok := request.Get["middle_id"]; !ok {
		// 请求中不存在数据
		response.StatusCode = 500
		return errors.New("参数有误")
	}
	middlewareID, err := strconv.ParseInt(request.Get["middle_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return errors.New("参数有误")
	}
	// 2.调用服务层接口
	middlewareInfo, err := m.MiddlewareService.QueryMiddlewareByID(ctx, &middleware.RequestMiddlewareID{Id: middlewareID})
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	// 3.查询结果序列化
	bytes, _ := json.Marshal(middlewareInfo)
	// 4.返回结果
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) QueryAllMiddleware(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryAllMiddleware 请求")
	// 1.调用服务层接口
	allMiddleware, err := m.MiddlewareService.QueryAllMiddleware(ctx, &middleware.RequestAll{})
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	// 2.查询结果序列化
	bytes, _ := json.Marshal(allMiddleware)
	// 3.返回结果
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) QueryAllMiddlewareByTypeID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.QueryAllMiddlewareByTypeID 请求")
	// 0.判断请求中是否存在数据
	if _, ok := request.Get["type_id"]; !ok {
		logger.Error("请求参数有误")
		response.StatusCode = 500
		return nil
	}
	// 1.从请求url中获取typeID
	// 2.将typeID转换为int64
	typeID, err := strconv.ParseInt(request.Get["type_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	// 3.调用服务层接口
	allMiddlewareType, err := m.MiddlewareService.QueryAllMiddlewareByTypeID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		logger.Error(err)
		response.StatusCode = 500
		return err
	}
	// 4.将查询结果序列化
	bytes, err := json.Marshal(allMiddlewareType)
	// 5.返回结果
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) AddMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	log.Info("接收到 middleware.AddMiddlewareType 请求")
	middleTypeInfo := &middleware.RMiddleTypeInfo{}

	if err := json.Unmarshal([]byte(request.Body), middleTypeInfo); err != nil {
		// 解析出错，返回错误信息
		response.StatusCode = http.StatusBadRequest
		response.Body = err.Error()
		return nil
	}

	responseInfo, err := m.MiddlewareService.AddMiddlewareType(ctx, middleTypeInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) DeleteMiddlewareTypeByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	// 1.判断请求中是否存在type_id
	if _, ok := request.Get["type_id"]; !ok {
		return errors.New("参数异常")
	}
	// 2.获取请求中的type_id
	typeID, err := strconv.ParseInt(request.Get["type"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 3.调用服务层接口
	responseInfo, err := m.MiddlewareService.DeleteMiddleTypeByID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	// 4.返回结果
	bytes, _ := json.Marshal(responseInfo)
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) UpdateMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	// 1.创建服务层对象
	typeInfo := &middleware.RMiddleTypeInfo{}
	// 2.将请求体中的json数据反序列化到结构体中
	err := json.Unmarshal([]byte(request.Body), typeInfo)
	if err != nil {
		logger.Error(err)
		response.Body = err.Error()
		return errors.New("参数有误")
	}
	// 3.调用服务层接口
	responseInfo, err := m.MiddlewareService.UpdateMiddleType(ctx, typeInfo)
	if err != nil {
		response.StatusCode = 500
		return err
	}
	// 4.返回相应
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) QueryMiddlewareTypeByID(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	// 1.判断请求中是否存在type_id
	if _, ok := request.Get["type_id"]; !ok {
		return errors.New("参数异常")
	}
	// 2.获取请求中的type_id
	typeID, err := strconv.ParseInt(request.Get["type_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 3.调用服务层接口
	typeInfo, err := m.MiddlewareService.QueryMiddleTypeByID(ctx, &middleware.RequestMiddleTypeID{TypeId: typeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	// 4.返回结果
	bytes, _ := json.Marshal(typeInfo)
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}

func (m *MiddlewareApi) QueryAllMiddlewareType(ctx context.Context, request *middlewareApi.Request, response *middlewareApi.Response) error {
	// 1.调用服务层接口
	middleType, err := m.MiddlewareService.QueryAllMiddleType(ctx, &middleware.RequestAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	// 2.返回结果
	bytes, _ := json.Marshal(middleType)
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}
