package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"pod/common"
	"pod/domain/model"
	"pod/domain/service"
	"pod/proto/pod"
	"strconv"
)

type PodHandler struct {
	PodDataService service.IPodDataService
}

func (p *PodHandler) AddPod(ctx context.Context, request *pod.RPodInfo, response *pod.ResponseMsg) error {
	podObj := &model.Pod{}
	err := common.SwapTo(request, podObj)
	if err != nil {
		logger.Error("类型转换失败, ", err)
		response.Msg = err.Error()
		return err
	}
	// 向k8s中添加pod
	err = p.PodDataService.CreateToK8s(request)
	// 添加失败
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	} else { // k8s添加pod成功
		// 操作数据库写入数据
		podID, err := p.PodDataService.AddPod(podObj)
		if err != nil {
			logger.Error(err)
			response.Msg = err.Error()
			return err
		}
		logger.Infof("Pod添加成功,ID为: ", podID)
		response.Msg = "Pod添加成功,ID为: " + strconv.FormatInt(podID, 10)
	}
	return nil
}

func (p *PodHandler) DeletePod(ctx context.Context, request *pod.RequestPodID, response *pod.ResponseMsg) error {
	podObj, err := p.PodDataService.QueryPodByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = p.PodDataService.DeleteFromK8s(podObj)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (p *PodHandler) QueryPodByID(ctx context.Context, request *pod.RequestPodID, response *pod.RPodInfo) error {
	podObj, err := p.PodDataService.QueryPodByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = common.SwapTo(podObj, response)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (p *PodHandler) UpdatePod(ctx context.Context, request *pod.RPodInfo, response *pod.ResponseMsg) error {
	err := p.PodDataService.UpdateToK8s(request)
	if err != nil {
		logger.Error(err)
		return err
	}
	// 存在Pod才能更新
	podObj, err := p.PodDataService.QueryPodByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = common.SwapTo(request, podObj)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = p.PodDataService.UpdatePod(podObj)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (p *PodHandler) QueryAllPods(ctx context.Context, request *pod.RequestQueryAll, response *pod.ResponseAllPod) error {
	pods, err := p.PodDataService.QueryAllPods()
	if err != nil {
		logger.Error(err)
		return err
	}
	for _, v := range pods {
		podInfo := &pod.RPodInfo{}
		err = common.SwapTo(v, podInfo)
		if err != nil {
			logger.Error(err)
			return err
		}
		response.PodInfo = append(response.PodInfo, podInfo)
	}
	return nil
}
