package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/pkg/errors"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"github.com/qyh794/go-paas/svcApi/pkg/jwt"
	"github.com/qyh794/go-paas/svcApi/proto/svcApi"
	"strconv"
)

type SvcApi struct {
	SvcService svc.SvcService
}

// QuerySvcByID 对外暴露的接口为:/svcApi/findSvcById，接收http请求
func (s *SvcApi) QuerySvcByID(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("接收到svcApi.QuerySvcByID的请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	svcIDStr, ok := request.Get["svc_id"]
	if !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	svcID, err := strconv.ParseInt(svcIDStr.Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	svcInfo, err := s.SvcService.QuerySvcByID(ctx, &svc.RequestSvcID{Id: svcID})
	if err != nil {
		return ResponseServerError(ctx, err, response)

	}
	return ResponseSucceed(ctx, svcInfo, response)
}

// AddSvc 对外暴露的接口为:/svcApi/AddSvc，接收http请求
func (s *SvcApi) AddSvc(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("添加svc服务")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	addSvcInfo := &svc.RSvcInfo{}

	// 将请求中数据转换到结构体中
	err = json.Unmarshal([]byte(request.Body), addSvcInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	res, err := s.SvcService.AddSvc(ctx, addSvcInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}

	return ResponseSucceed(ctx, res, response)
}

func (s *SvcApi) DeleteSvcByID(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("删除service服务")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 从请求中获取要删除的ID
	if _, ok := request.Get["svc_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 将ID进行类型转换
	svcID, err := strconv.ParseInt(request.Get["svc_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 调用后端服务
	res, err := s.SvcService.DeleteSvcByID(ctx, &svc.RequestSvcID{Id: svcID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 返回结果
	return ResponseSucceed(ctx, res, response)

}

func (s *SvcApi) UpdateSvc(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("更新service服务")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	//处理port
	updateSvcInfo := &svc.RSvcInfo{}
	// 从请求体中获取数据
	err = json.Unmarshal([]byte(request.Body), updateSvcInfo)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	res, err := s.SvcService.UpdateSvc(ctx, updateSvcInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)

	}
	return ResponseSucceed(ctx, res, response)
}

func (s *SvcApi) Call(ctx context.Context, request *svcApi.Request, response *svcApi.Response) error {
	log.Info("查询所有service服务")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	allSvc, err := s.SvcService.QueryAll(ctx, &svc.RequestQueryAll{})
	if err != nil {
		return ResponseServerError(ctx, err, response)

	}
	return ResponseSucceed(ctx, allSvc, response)
}
