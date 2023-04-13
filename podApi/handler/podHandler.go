package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/podApi/proto/podApi"
	"strconv"
)

const (
	succeed   = 200
	wrongArgs = 500
)

// PodApi 实现PodApiHandler接口
type PodApi struct {
	PodService pod.PodService
}

// QueryPodByID podApi.QueryPodByID 通过Api向外暴露接口为: /podApi/queryPodByID,接受http请求
// /podApi/FindPodById 请求会调用podApi.QueryPodByID 方法
func (p *PodApi) QueryPodByID(ctx context.Context, request *podApi.Request, response *podApi.Response) error {
	fmt.Println("接收到 podApi.QueryPodByID 的请求")
	// 获取请求中的数据
	val, ok := request.Get["pod_id"]
	if !ok {
		response.StatusCode = wrongArgs
		return errors.New("参数异常")
	}
	// 类型转换
	podID, err := strconv.ParseInt(val.Values[0], 10, 64)
	if err != nil {
		return err
	}
	podInfo, err := p.PodService.QueryPodByID(ctx, &pod.RequestPodID{Id: podID})
	if err != nil {
		return err
	}
	// json返回pod信息
	response.StatusCode = succeed
	bytes, _ := json.Marshal(podInfo)
	response.Body = string(bytes)
	return nil
}

// AddPod podApi.AddPod 通过API向外暴露为/podApi/addPod,
// podApi/AddPod 请求会调用go.micro.api.podApi 服务的podApi.AddPod 方法
func (p *PodApi) AddPod(ctx context.Context, request *podApi.Request, response *podApi.Response) error {
	fmt.Println("接受到 podApi.AddPod请求")
	// 请求中带有多个端口号,需要获取然后再添加到podInfo中
	podInfoObj := &pod.RPodInfo{}
	// 从请求体中获取参数,将其反序列化到结构体中
	err := json.Unmarshal([]byte(request.Body), podInfoObj)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
	// 调用pod.AddPod
	responseInfo, err := p.PodService.AddPod(ctx, podInfoObj)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

// DeletePodByID podApi.DeletePodById 通过API向外暴露为/podApi/deletePodById，接收http请求
func (p *PodApi) DeletePodByID(ctx context.Context, request *podApi.Request, response *podApi.Response) error {
	fmt.Println("接收到podApi.DeletePodByID的请求")
	// 从请求url中获取podID
	val, ok := request.Get["pod_id"]
	if !ok {
		return errors.New("参数异常")
	}
	// 将podID进行类型转换
	podID, err := strconv.ParseInt(val.Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 调用deletePod服务
	responseInfo, err := p.PodService.DeletePod(ctx, &pod.RequestPodID{Id: podID})
	if err != nil {
		logger.Error(err)
		return err
	}
	// 返回结果
	response.StatusCode = succeed
	bytes, err := json.Marshal(responseInfo)
	if err != nil {
		return err
	}
	response.Body = string(bytes)
	return nil
}

// UpdatePod podApi.UpdatePod 通过API向外暴露为/podApi/updatePod，接收http请求
func (p *PodApi) UpdatePod(ctx context.Context, request *podApi.Request, response *podApi.Response) error {
	fmt.Println("接收到podApi.UpdatePod请求")
	podObj := &pod.RPodInfo{}
	// 处理请求中port信息
	err := json.Unmarshal([]byte(request.Body), podObj)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
	// 调用UpdatePod服务
	responseInfo, err := p.PodService.UpdatePod(ctx, podObj)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(responseInfo)
	response.Body = string(bytes)
	return nil
}

// Call 默认的方法podApi.Call 通过API向外暴露为/podApi/call，接收http请求
func (p *PodApi) Call(ctx context.Context, request *podApi.Request, response *podApi.Response) error {
	fmt.Println("接收到podApi.QueryAll 请求")
	pods, err := p.PodService.QueryAllPods(ctx, &pod.RequestQueryAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = succeed
	bytes, _ := json.Marshal(pods)
	response.Body = string(bytes)
	return nil
}
