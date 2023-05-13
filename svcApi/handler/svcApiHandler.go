package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/pkg/errors"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"github.com/qyh794/go-paas/svcApi/pkg/jwt"
	"github.com/qyh794/go-paas/svcApi/proto/svcApi"
	"strconv"
	"strings"
)

const (
	serverBusy       = 500
	succeed          = 200
	CodeInvalidParam = 400
)

type SvcApi struct {
	SvcService svc.SvcService
}

// QuerySvcByID 对外暴露的接口为:/svcApi/findSvcById，接收http请求
func (s *SvcApi) QuerySvcByID(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("接收到svcApi.QuerySvcByID的请求")
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	addSvcInfo := &svc.RSvcInfo{}

	// 将请求中数据转换到结构体中
	err = json.Unmarshal([]byte(request.Body), addSvcInfo)
	if err != nil {
		response.StatusCode = CodeInvalidParam
		response.Body = err.Error()
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	//处理port
	updateSvcInfo := &svc.RSvcInfo{}
	// 从请求体中获取数据
	err = json.Unmarshal([]byte(request.Body), updateSvcInfo)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}

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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	allSvc, err := s.SvcService.QueryAll(ctx, &svc.RequestQueryAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(allSvc)
	response.Body = string(bytes)
	return nil
}
