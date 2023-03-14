package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/svc/domain/model"
	"github.com/qyh794/go-paas/svc/domain/service"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"strconv"
)

type SvcHandler struct {
	SvcDataService service.IServiceDataService
}

func (s *SvcHandler) AddSvc(ctx context.Context, request *svc.RSvcInfo, response *svc.Response) error {
	log.Info("创建service")
	svcObj := &model.Svc{}
	// 将请求中的数据反序列化到结构体中
	if err := common.SwapTo(request, svcObj); err != nil {
		logger.Error(err)
		return err
	}
	// k8s创建service资源
	// 创建失败,返回
	if err := s.SvcDataService.CreateServiceToK8s(request); err != nil {
		logger.Error(err)
		return err
	} else {
		// 创建成功, 操作数据库添加数据
		serviceID, err := s.SvcDataService.AddService(svcObj)
		if err != nil {
			logger.Error(err)
			return err
		}
		logger.Info("service ID: " + strconv.FormatInt(serviceID, 10) + " 添加成功")
		response.Msg = "service ID: " + strconv.FormatInt(serviceID, 10) + " 添加成功"
	}
	return nil
}

func (s *SvcHandler) DeleteSvcByID(ctx context.Context, request *svc.RequestSvcID, response *svc.Response) error {
	logger.Info("删除service")
	// 通过ID查询数据库中service数据
	serviceInfo, err := s.SvcDataService.QueryServiceByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	// k8s删除service
	if err = s.SvcDataService.DeleteFromK8s(serviceInfo); err != nil {
		// 删除失败直接返回
		logger.Error(err)
		return err
	}
	// 删除成功,操作数据库删除service数据
	if err = s.SvcDataService.DeleteServiceByID(request.Id); err != nil {
		logger.Error(err)
		return err
	}
	response.Msg = "删除service ID: " + strconv.FormatInt(request.Id, 10) + "成功"
	return nil
}

func (s *SvcHandler) UpdateSvc(ctx context.Context, request *svc.RSvcInfo, response *svc.Response) error {
	logger.Info("更新service")
	// @TODO 数据一致性,校验或者分布式事务
	// 删除k8s中的资源,删除失败直接返回
	if err := s.SvcDataService.UpdateServiceToK8s(request); err != nil {
		logger.Error(err)
		return err
	}
	// 删除成功, 通过查询数据库中的旧数据
	serviceInfo, err := s.SvcDataService.QueryServiceByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 将请求中新数据反序列化到结构体中
	if err = common.SwapTo(request, serviceInfo); err != nil {
		logger.Error(err)
		return err
	}
	// 操作数据库执行更新操作
	if err = s.SvcDataService.UpdateService(serviceInfo); err != nil {
		logger.Error(err)
		return err
	}
	response.Msg = "更新service ID : " + strconv.FormatInt(request.Id, 10) + "成功"
	return nil
}

func (s *SvcHandler) QuerySvcByID(ctx context.Context, request *svc.RequestSvcID, response *svc.RSvcInfo) error {
	logger.Info("查询service")
	// 查询数据库
	serviceInfo, err := s.SvcDataService.QueryServiceByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 将查询到的结果反序列化到response中
	if err = common.SwapTo(serviceInfo, response); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *SvcHandler) QueryAll(ctx context.Context, request *svc.RequestQueryAll, response *svc.ResponseAllSvc) error {
	logger.Info("查询所有service")
	allService, err := s.SvcDataService.QueryAllService()
	if err != nil {
		logger.Info(err)
		return err
	}
	// allService := []model.svc
	for i := range allService {
		serviceInfo := &svc.RSvcInfo{}
		if err = common.SwapTo(allService[i], serviceInfo); err != nil {
			logger.Error(err)
			return err
		}
		response.SvcInfo = append(response.SvcInfo, serviceInfo)
	}
	// need return []*RSvcInfo
	return nil
}
