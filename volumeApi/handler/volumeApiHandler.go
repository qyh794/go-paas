package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volumeApi/pkg/jwt"

	"github.com/qyh794/go-paas/volumeApi/proto/volumeApi"
	"strconv"
	"strings"
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	volumeInfo := &volume.RVolumeInfo{}
	err = json.Unmarshal([]byte(request.Body), volumeInfo)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
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
	token := request.Header["Authorization"].GetValues()[0]
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
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
