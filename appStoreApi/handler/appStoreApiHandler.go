package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qyh794/go-paas/appStore/proto/appStore"
	"github.com/qyh794/go-paas/appStoreApi/pkg/jwt"
	"github.com/qyh794/go-paas/appStoreApi/proto/appStoreApi"
	"strconv"
)

type AppStoreApi struct {
	AppStoreService appStore.AppStoreService
}

func (a *AppStoreApi) AddApp(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 从请求体中获取数据
	// 1.将请求体中的数据反序列化到结构体中
	appInfo := &appStore.RAppInfo{}
	if err = json.Unmarshal([]byte(request.Body), appInfo); err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	// 2.调用服务层接口
	info, err := a.AppStoreService.AddApp(ctx, appInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, info, response)
}

func (a *AppStoreApi) DeleteAppByID(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	responseInfo, err := a.AppStoreService.DeleteAppByID(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, responseInfo, response)
}

func (a *AppStoreApi) UpdateApp(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.将请求体中的数据反序列化到结构体中
	appInfo := &appStore.RAppInfo{}
	if err := json.Unmarshal([]byte(request.Body), appInfo); err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	// 2.调用服务层接口
	responseInfo, err := a.AppStoreService.UpdateApp(ctx, appInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, responseInfo, response)
}

func (a *AppStoreApi) QueryAppByID(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	appInfo, err := a.AppStoreService.QueryAppByID(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, appInfo, response)
}

func (a *AppStoreApi) QueryAllApps(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.调用服务层接口
	allApps, err := a.AppStoreService.QueryAllApps(ctx, &appStore.RequestAllApp{})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 2.返回响应
	return ResponseSucceed(ctx, allApps, response)
}

func (a *AppStoreApi) AddInstallNum(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	responseInfo, err := a.AppStoreService.AddInstallNum(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, responseInfo, response)
}

func (a *AppStoreApi) QueryInstallNum(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}

	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	installNum, err := a.AppStoreService.QueryInstallNum(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, installNum, response)
}

func (a *AppStoreApi) AddViewNum(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	responseInfo, err := a.AppStoreService.AddViewNum(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, responseInfo, response)
}

func (a *AppStoreApi) QueryViewNum(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从url中获取appID
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	installNum, err := a.AppStoreService.QueryViewNum(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, installNum, response)
}

func (a *AppStoreApi) AddComment(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从请求体中获取参数,反序列化到结构体中
	appComment := &appStore.RAppComment{}
	if err = json.Unmarshal([]byte(request.Body), appComment); err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	// 2.调用服务层接口
	responseInfo, err := a.AppStoreService.AddComment(ctx, appComment)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, responseInfo, response)
}

func (a *AppStoreApi) QueryAllCommentByID(ctx context.Context, request *appStoreApi.Request, response *appStoreApi.Response) error {
	if _, ok := request.Header["Authorization"]; !ok {
		return ResponseEmptyToken(ctx, response)
	}
	err := jwt.CheckToken(request.Header["Authorization"].GetValues()[0])
	if err != nil {
		return ResponseAuthFailed(ctx, err, response)
	}
	// 1.从请求中获取app_id
	appID, err := strconv.ParseInt(request.Get["app_id"].Values[0], 10, 64)
	if err != nil {
		return ResponseBadRequest(ctx, errors.New("wrong args"), response)
	}
	// 2.调用服务层接口
	allAppComment, err := a.AppStoreService.QueryAllCommentByID(ctx, &appStore.RequestID{Id: appID})
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	// 3.返回响应
	return ResponseSucceed(ctx, allAppComment, response)
}
