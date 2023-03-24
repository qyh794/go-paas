package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volumeApi/form"
	"github.com/qyh794/go-paas/volumeApi/proto/volumeApi"
	"strconv"
)

type VolumeApi struct {
	VolumeService volume.VolumeService
}

const (
	WrongArgs = 409
	Succeed   = 200
)

func (v *VolumeApi) QueryVolumeByID(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.QueryVolumeByID 请求")
	if _, ok := request.Get["volume_id"]; !ok {
		response.StatusCode = WrongArgs
		return errors.New("参数有误")
	}
	volumeID, err := strconv.ParseInt(request.Get["volume_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	volumeInfo, err := v.VolumeService.QueryVolumeByID(ctx, &volume.RequestVolumeID{Id: volumeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(volumeInfo)
	response.Body = string(bytes)
	return nil
}

func (v *VolumeApi) AddVolume(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.AddVolumeByID 请求")
	volumeInfo := &volume.RVolumeInfo{}
	form.FormToSvcStruct(request.Post, volumeInfo)
	volumeID, err := v.VolumeService.AddVolume(ctx, volumeInfo)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(volumeID)
	response.Body = string(bytes)
	return nil
}

func (v *VolumeApi) DeleteVolumeByID(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.DeleteVolumeByID 请求")
	if _, ok := request.Get["volume_id"]; !ok {
		response.StatusCode = WrongArgs
		return errors.New("参数异常")
	}
	volumeID, err := strconv.ParseInt(request.Get["volume_id"].Values[0], 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	res, err := v.VolumeService.DeleteVolume(ctx, &volume.RequestVolumeID{Id: volumeID})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(res)
	response.Body = string(bytes)
	return nil
}

func (v *VolumeApi) Call(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.QueryAllVolume 请求")
	allVolume, err := v.VolumeService.QueryAllVolume(ctx, &volume.RequestQueryAll{})
	if err != nil {
		logger.Error(err)
		return err
	}
	response.StatusCode = Succeed
	bytes, _ := json.Marshal(allVolume)
	response.Body = string(bytes)
	return nil
}
