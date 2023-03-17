package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/pkg/errors"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"strconv"
	"svcApi/plugin/form"
	"svcApi/proto/svcApi"
)

const (
	serverBusy = 500
	succeed = 200
	CodeInvalidParam = 400
)

type SvcApi struct {
	SvcService svc.SvcService
}

// QuerySvcByID 对外暴露的接口为:/svcApi/findSvcById，接收http请求
func (s *SvcApi) QuerySvcByID(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("接收到svcApi.QuerySvcByID的请求")
	svcIDStr, ok := request.Get["svc_id"]
	if !ok {
		response.StatusCode = CodeInvalidParam
		return errors.New("参数异常")
	}
	svcID, err := strconv.ParseInt(svcIDStr.Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	svcInfo, err := s.SvcService.QuerySvcByID(ctx, &svc.RequestSvcID{Id: svcID})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(svcInfo)
	response.Body = string(bytes)
	return nil
}

// AddSvc 对外暴露的接口为:/svcApi/AddSvc，接收http请求
func (s *SvcApi) AddSvc(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("添加svc服务")
	addSvcInfo := &svc.RSvcInfo{}
	// 从post请求中获取到所有的svc_type
	svcType, ok := request.Post["svc_type"]
	// 变量 svcType 的类型为 []string
	// 如果请求中带有 svc_type,进行处理
	if ok && len(svcType.Values) > 0{ //svcType.Values 是获取到 svcType 字符串数组中的所有元素
		svcPort := &svc.SvcPort{}
		switch svcType.Values[0] { // 判断[]string第一个, svcType可能包含ClusterIP\NodePort\LoadBalancer\ExternalName
		case "ClusterIP":
			// 获取message SvcPort {}中svc_port的第一个值
			port, err := strconv.ParseInt(request.Post["svc_port"].Values[0], 10, 32)
			if err != nil {
				logger.Error(err)
				return err
			}
			svcPort.SvcPort = int32(port)
			targetPort, err := strconv.ParseInt(request.Post["svc_target_port"].Values[0], 10, 32)
			if err != nil {
				logger.Error(err)
				return err
			}
			svcPort.SvcTargetPort = int32(targetPort)
			svcPort.SvcPortProtocol = request.Post["svc_port_protocol"].Values[0]
			addSvcInfo.SvcPort = append(addSvcInfo.SvcPort, svcPort)
		default:
			return errors.New("类型不支持")
		}
	}
	// 将请求中数据转换到结构体中
	form.FormToSvcStruct(request.Post, addSvcInfo)
	res, err := s.SvcService.AddSvc(ctx, addSvcInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(res)
	response.Body = string(bytes)
	return nil
}

func (s *SvcApi) DeleteSvcByID(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("删除service服务")
	// 从请求中获取要删除的ID
	if _, ok := request.Get["svc_id"]; !ok {
		return errors.New("参数异常")
	}
	// 将ID进行类型转换
	svcID, err := strconv.ParseInt(request.Get["svc_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 调用后端服务
	res, err := s.SvcService.DeleteSvcByID(ctx, &svc.RequestSvcID{Id: svcID})
	if err != nil {
		logger.Error(err)
		return err
	}
	// 返回结果
	response.StatusCode = succeed
	bytes, _ := json.Marshal(res)
	response.Body = string(bytes)
	return nil
}

func (s *SvcApi) UpdateSvc(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("更新service服务")
	//处理port
	updateSvcInfo := &svc.RSvcInfo{}
	// 从post请求中获取到所有的svc_type
	svcType, ok := request.Post["svc_type"]
	// 变量 svcType 的类型为 []string
	// 如果请求中带有 svc_type,进行处理
	if ok && len(svcType.Values) > 0{ //svcType.Values 是获取到 svcType 字符串数组中的所有元素
		svcPort := &svc.SvcPort{}
		switch svcType.Values[0] { // 判断[]string第一个, svcType可能包含ClusterIP\NodePort\LoadBalancer\ExternalName
		case "ClusterIP":
			// 获取message SvcPort {}中svc_port的第一个值
			port, err := strconv.ParseInt(request.Post["svc_port"].Values[0], 10, 32)
			if err != nil {
				logger.Error(err)
				return err
			}
			svcPort.SvcPort = int32(port)
			targetPort, err := strconv.ParseInt(request.Post["svc_target_port"].Values[0], 10, 32)
			if err != nil {
				logger.Error(err)
				return err
			}
			svcPort.SvcTargetPort = int32(targetPort)
			svcPort.SvcPortProtocol = request.Post["svc_port_protocol"].Values[0]
			updateSvcInfo.SvcPort = append(updateSvcInfo.SvcPort, svcPort)
		default:
			return errors.New("类型不支持")
		}
	}
	// 将请求中数据转换到结构体中
	form.FormToSvcStruct(request.Post, updateSvcInfo)
	res, err := s.SvcService.UpdateSvc(ctx, updateSvcInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(res)
	response.Body = string(bytes)
	return nil
}

func (s *SvcApi) Call(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("查询所有service服务")
	allSvc,err:=s.SvcService.QueryAll(ctx,&svc.RequestQueryAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(allSvc)
	response.Body = string(bytes)
	return nil
}

