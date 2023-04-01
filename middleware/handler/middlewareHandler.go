package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/middleware/domain/model"
	"github.com/qyh794/go-paas/middleware/domain/service"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"strconv"
)

type MiddlewareHandler struct {
	MiddlewareDataService     service.IMiddlewareDataService
	MiddlewareTypeDataService service.IMiddlewareTypeDataService
}

func (m *MiddlewareHandler) AddMiddleware(ctx context.Context, requestMiddlewareInfo *middleware.RMiddlewareInfo, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.AddMiddleware请求")
	middlewareObj := &model.Middleware{}
	// 将请求中的数据序列化到结构体中
	if err := common.SwapTo(requestMiddlewareInfo, middlewareObj); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	// 根据请求中的中间件版本ID middle_version_id 获取平台预先提供的中间件镜像版本
	image, err := m.MiddlewareTypeDataService.QueryImageVersionByID(requestMiddlewareInfo.MiddleVersionId)
	if err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	requestMiddlewareInfo.MiddleDockerImageVersion = image
	// k8s中创建资源
	if err = m.MiddlewareDataService.CreateToK8s(requestMiddlewareInfo); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	} else { // 创建成功
		// 数据库记录数据
		middlewareID, err := m.MiddlewareDataService.AddMiddleware(middlewareObj)
		if err != nil {
			logger.Error(err)
			responseInfo.Msg = err.Error()
			return err
		}
		// 数据库记录数据成功
		responseInfo.Msg = "中间件ID: " + strconv.FormatInt(middlewareID, 10) + "添加成功"
		logger.Info(responseInfo.Msg)
	}
	return nil
}

func (m *MiddlewareHandler) DeleteMiddlewareByID(ctx context.Context, requestMiddlewareID *middleware.RequestMiddlewareID, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.DeleteMiddlewareByID请求")
	middlewareObj, err := m.MiddlewareDataService.QueryMiddlewareByID(requestMiddlewareID.Id)
	if err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	if err = m.MiddlewareDataService.DeleteFromK8s(middlewareObj); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	// K8s删除资源成功
	// 删除数据库中的记录
	if err = m.MiddlewareDataService.DeleteMiddlewareByID(requestMiddlewareID.Id); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	return nil
}

func (m *MiddlewareHandler) UpdateMiddleware(ctx context.Context, requestMiddlewareInfo *middleware.RMiddlewareInfo, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.UpdateMiddleware请求")
	if err := m.MiddlewareDataService.UpdateToK8s(requestMiddlewareInfo); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	middlewareInfo, err := m.MiddlewareDataService.QueryMiddlewareByID(requestMiddlewareInfo.Id)
	if err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	if err = common.SwapTo(requestMiddlewareInfo, middlewareInfo); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	// 更新数据库中的记录
	if err = m.MiddlewareDataService.UpdateMiddleware(middlewareInfo); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	return nil
}

func (m *MiddlewareHandler) QueryMiddlewareByID(ctx context.Context, requestMiddlewareID *middleware.RequestMiddlewareID, responseInfo *middleware.RMiddlewareInfo) error {
	log.Info("接收到middleware.QueryMiddlewareByID请求")
	// 1. 数据库查询中间件信息
	middlewareInfo, err := m.MiddlewareDataService.QueryMiddlewareByID(requestMiddlewareID.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 2. 将查询结果序列化到响应中
	if err = common.SwapTo(middlewareInfo, responseInfo); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (m *MiddlewareHandler) QueryAllMiddleware(ctx context.Context, requestAll *middleware.RequestAll, responseAllMiddleware *middleware.ResponseAllMiddleware) error {
	log.Info("接收到middleware.QueryAllMiddleware请求")
	allMiddleware, err := m.MiddlewareDataService.QueryAllMiddleware()
	if err != nil {
		logger.Error()
		return err
	}
	for i := 0; i < len(allMiddleware); i++ {
		middlewareInfo := &middleware.RMiddlewareInfo{}
		if err = common.SwapTo(allMiddleware[i], middlewareInfo); err != nil {
			logger.Error(err)
			return err
		}
		responseAllMiddleware.MiddlewareInfo = append(responseAllMiddleware.MiddlewareInfo, middlewareInfo)
	}
	return nil
}

// QueryAllMiddlewareByTypeID 通过中间件类型查找所有的中间件
func (m *MiddlewareHandler) QueryAllMiddlewareByTypeID(ctx context.Context, requestMiddlewareTypeID *middleware.RequestMiddleTypeID, responseAllMiddleware *middleware.ResponseAllMiddleware) error {
	log.Info("接收到middleware.QueryAllMiddlewareByTypeID请求")
	middleType, err := m.MiddlewareDataService.QueryAllMiddlewareByTypeID(requestMiddlewareTypeID.TypeId)
	if err != nil {
		logger.Error(err)
		return err
	}
	for i := 0; i < len(middleType); i++ {
		middlewareInfo := &middleware.RMiddlewareInfo{}
		if err = common.SwapTo(middleType[i], middlewareInfo); err != nil {
			logger.Error(err)
			return err
		}
		responseAllMiddleware.MiddlewareInfo = append(responseAllMiddleware.MiddlewareInfo, middlewareInfo)
	}
	return nil
}

func (m *MiddlewareHandler) QueryMiddleTypeByID(ctx context.Context, requestTypeID *middleware.RequestMiddleTypeID, responseTypeInfo *middleware.RMiddleTypeInfo) error {
	log.Info("接收到middleware.QueryMiddleTypeByID请求")
	middleType, err := m.MiddlewareTypeDataService.QueryMiddlewareTypeByID(requestTypeID.TypeId)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err = common.SwapTo(middleType, responseTypeInfo); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (m *MiddlewareHandler) AddMiddlewareType(ctx context.Context, requestMiddleTypeInfo *middleware.RMiddleTypeInfo, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.AddMiddlewareType请求")
	middlewareType := &model.MiddleType{}
	if err := common.SwapTo(requestMiddleTypeInfo, middlewareType); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	middlewareTypeID, err := m.MiddlewareTypeDataService.AddMiddlewareType(middlewareType)
	if err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	responseInfo.Msg = "中间件类型ID: " + strconv.FormatInt(middlewareTypeID, 10) + "添加成功"
	logger.Info(responseInfo.Msg)
	return nil
}

func (m *MiddlewareHandler) DeleteMiddleTypeByID(ctx context.Context, requestMiddleTypeID *middleware.RequestMiddleTypeID, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.DeleteMiddleTypeByID请求")
	if err := m.MiddlewareTypeDataService.DeleteMiddlewareTypeByID(requestMiddleTypeID.TypeId); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	return nil
}

func (m *MiddlewareHandler) UpdateMiddleType(ctx context.Context, requestMiddlewareTypeInfo *middleware.RMiddleTypeInfo, responseInfo *middleware.ResponseInfo) error {
	log.Info("接收到middleware.UpdateMiddleType请求")
	middlewareType, err := m.MiddlewareTypeDataService.QueryMiddlewareTypeByID(requestMiddlewareTypeInfo.Id)
	if err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	if err = common.SwapTo(requestMiddlewareTypeInfo, middlewareType); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	if err = m.MiddlewareTypeDataService.UpdateMiddlewareType(middlewareType); err != nil {
		logger.Error(err)
		responseInfo.Msg = err.Error()
		return err
	}
	responseInfo.Msg = "中间件类型ID: " + strconv.FormatInt(requestMiddlewareTypeInfo.Id, 10) + "更新成功"
	return nil
}

func (m *MiddlewareHandler) QueryAllMiddleType(ctx context.Context, requestAll *middleware.RequestAll, responseAllMiddleType *middleware.ResponseAllMiddleType) error {
	log.Info("接收到middleware.AddMiddleware请求")
	allMiddlewareType, err := m.MiddlewareTypeDataService.QueryAllMiddlewareType()
	if err != nil {
		logger.Error(err)
		return err
	}
	for i := 0; i < len(allMiddlewareType); i++ {
		middlewareType := &middleware.RMiddleTypeInfo{}
		if err = common.SwapTo(allMiddlewareType[i], middlewareType); err != nil {
			logger.Error(err)
			return err
		}
		responseAllMiddleType.MiddlewareTypeInfo = append(responseAllMiddleType.MiddlewareTypeInfo, middlewareType)
	}
	return nil
}
