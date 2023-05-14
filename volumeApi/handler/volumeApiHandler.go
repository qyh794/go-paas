package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volumeApi/pkg/jwt"

	"github.com/qyh794/go-paas/volumeApi/proto/volumeApi"
	"strconv"
)

type VolumeApi struct {
	VolumeService volume.VolumeService
}

func (v *VolumeApi) QueryVolumeByID(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.QueryVolumeByID 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	if _, ok := request.Get["volume_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	volumeID, err := strconv.ParseInt(request.Get["volume_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	volumeInfo, err := v.VolumeService.QueryVolumeByID(ctx, &volume.RequestVolumeID{Id: volumeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, volumeInfo, response)
}

func (v *VolumeApi) AddVolume(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.AddVolumeByID 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	volumeInfo := &volume.RVolumeInfo{}
	err = json.Unmarshal([]byte(request.Body), volumeInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	volumeID, err := v.VolumeService.AddVolume(ctx, volumeInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, volumeID, response)
}

func (v *VolumeApi) DeleteVolumeByID(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.DeleteVolumeByID 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	if _, ok := request.Get["volume_id"]; !ok {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	volumeID, err := strconv.ParseInt(request.Get["volume_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	res, err := v.VolumeService.DeleteVolume(ctx, &volume.RequestVolumeID{Id: volumeID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, res, response)
}

func (v *VolumeApi) Call(ctx context.Context, request *volumeApi.Request, response *volumeApi.Response) error {
	log.Info("接收到 volumeApi.QueryAllVolume 请求")
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}

	allVolume, err := v.VolumeService.QueryAllVolume(ctx, &volume.RequestQueryAll{})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}

	return ResponseSucceed(ctx, allVolume, response)
}
