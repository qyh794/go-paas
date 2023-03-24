package handler

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/micro/micro/v3/service/logger"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/volume/domain/model"
	"github.com/qyh794/go-paas/volume/domain/service"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"strconv"
)

type VolumeHandler struct {
	VolumeDataService service.IVolumeDataService
}

func (v *VolumeHandler) AddVolume(ctx context.Context, request *volume.RVolumeInfo, response *volume.Response) error {
	log.Info("接收到 volume.AddVolume请求")
	volumeObj := &model.Volume{}
	if err := common.SwapTo(request, volumeObj); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	if err := v.VolumeDataService.CreateVolumeToK8s(request); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	} else {
		volumeID, err := v.VolumeDataService.AddVolume(volumeObj)
		if err != nil {
			logger.Error(err)
			response.Msg = err.Error()
			return err
		}
		response.Msg = "volume ID: " + strconv.FormatInt(volumeID, 10) + "添加成功"
		return nil
	}
}

func (v *VolumeHandler) QueryVolumeByID(ctx context.Context, request *volume.RequestVolumeID, response *volume.RVolumeInfo) error {
	log.Info("接收到 volume.QueryVolumeByID请求")
	volumeInfo, err := v.VolumeDataService.QueryVolumeByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err = common.SwapTo(volumeInfo, request); err != nil {
		logger.Error(err)
		return err
	}
	return err
}

func (v *VolumeHandler) DeleteVolume(ctx context.Context, request *volume.RequestVolumeID, response *volume.Response) error {
	log.Info("接收到 volume.DeleteVolume请求")
	volumeInfo, err := v.VolumeDataService.QueryVolumeByID(request.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	// k8s
	// @TODO 可以优化,删除k8s中的volume只用到了两个参数
	if err = v.VolumeDataService.DeleteVolumeFromK8s(volumeInfo); err != nil {
		logger.Error(err)
		return err
	}
	// 数据库
	// @TODO 数据一致性,分布式事务或者代码校验
	if err = v.VolumeDataService.DeleteVolumeByID(request.Id); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (v *VolumeHandler) UpdateVolume(ctx context.Context, request *volume.RVolumeInfo, response *volume.Response) error {
	log.Info("接收到 volume.UpdateVolume请求")
	//TODO implement me pvc更新待做
	panic("implement me")
}

func (v *VolumeHandler) QueryAllVolume(ctx context.Context, request *volume.RequestQueryAll, response *volume.ResponseAllVolume) error {
	log.Info("接收到 volume.QueryAllVolume请求")
	allVolume, err := v.VolumeDataService.QueryAllVolume()
	if err != nil {
		logger.Error(err)
		return err
	}
	for i := 0; i < len(allVolume); i++ {
		volumeInfo := &volume.RVolumeInfo{}
		err = common.SwapTo(allVolume[i], volumeInfo)
		if err != nil {
			logger.Error(err)
			return err
		}
		response.VolumeInfo = append(response.VolumeInfo, volumeInfo)
	}
	return nil
}
