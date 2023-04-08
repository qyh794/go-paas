package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/appStore/domain/model"
	"github.com/qyh794/go-paas/appStore/domain/service"
	"github.com/qyh794/go-paas/appStore/proto/appStore"
	"github.com/qyh794/go-paas/common"
	"strconv"
)

type AppStoreHandler struct {
	AppStoreDataService service.IAppStoreDataService
}

func (a *AppStoreHandler) AddApp(ctx context.Context, requestAppInfo *appStore.RAppInfo, response *appStore.Response) error {
	log.Info("接收到appStore.AddApp请求")
	appInfo := &model.AppStore{}
	err := common.SwapTo(requestAppInfo, appInfo)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	appID, err := a.AppStoreDataService.AddApp(appInfo)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	response.Msg = "添加应用成功,应用ID: " + strconv.FormatInt(appID, 10)
	logger.Info(response.Msg)
	return nil
}

func (a *AppStoreHandler) DeleteAppByID(ctx context.Context, requestID *appStore.RequestID, response *appStore.Response) error {
	log.Info("接收到appStore.DeleteAppByID请求")
	err := a.AppStoreDataService.DeleteApp(requestID.Id)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	return nil
}

func (a *AppStoreHandler) UpdateApp(ctx context.Context, requestAppInfo *appStore.RAppInfo, response *appStore.Response) error {
	log.Info("接收到 appStore.UpdateApp 请求")
	// 先做一次查询,再将新数据和旧数据结合,最后操作数据库更新,因为允许用户更新时之填部分字段
	appInfo, err := a.AppStoreDataService.QueryAppByID(requestAppInfo.Id)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	if err = common.SwapTo(requestAppInfo, appInfo); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	if err = a.AppStoreDataService.UpdateApp(appInfo); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	return nil
}

func (a *AppStoreHandler) QueryAppByID(ctx context.Context, requestID *appStore.RequestID, responseAppInfo *appStore.RAppInfo) error {
	log.Info("接收到 appStore.QueryAppByID 请求")
	appInfo, err := a.AppStoreDataService.QueryAppByID(requestID.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err = common.SwapTo(appInfo, responseAppInfo); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (a *AppStoreHandler) QueryAllApps(ctx context.Context, requestAllApp *appStore.RequestAllApp, responseAllApp *appStore.ResponseAllApp) error {
	log.Info("接收到 appStore.QueryAllApps 请求")
	allApp, err := a.AppStoreDataService.QueryAllApp()
	if err != nil {
		logger.Error(err)
		return err
	}
	for i, _ := range allApp {
		appInfo := &appStore.RAppInfo{}
		if err = common.SwapTo(allApp[i], appInfo); err != nil {
			logger.Error(err)
			return err
		}
		responseAllApp.AppInfo = append(responseAllApp.AppInfo, appInfo)
	}
	return nil
}

func (a *AppStoreHandler) AddInstallNum(ctx context.Context, requestID *appStore.RequestID, response *appStore.Response) error {
	log.Info("接收到 appStore.AddInstallNum 请求")
	if err := a.AppStoreDataService.AddAppInstallNumByID(requestID.Id); err != nil {
		logger.Error(err)
		return err
	}
	response.Msg = "统计成功"
	return nil
}

func (a *AppStoreHandler) QueryInstallNum(ctx context.Context, requestID *appStore.RequestID, responseNum *appStore.ResponseNum) error {
	log.Info("接收到 appStore.QueryInstallNum 请求")
	installNum := a.AppStoreDataService.QueryAppInstallNumByID(requestID.Id)
	responseNum.Num = installNum
	return nil
}

func (a *AppStoreHandler) AddViewNum(ctx context.Context, requestID *appStore.RequestID, response *appStore.Response) error {
	log.Info("接收到 appStore.AddViewNum 请求")
	if err := a.AppStoreDataService.AddAppViewNumByID(requestID.Id); err != nil {
		logger.Error(err)
		return err
	}
	response.Msg = "统计成功"
	return nil
}

func (a *AppStoreHandler) QueryViewNum(ctx context.Context, requestID *appStore.RequestID, responseNum *appStore.ResponseNum) error {
	log.Info("接收到 appStore.QueryViewNum 请求")
	viewNum := a.AppStoreDataService.QueryAppViewNumByID(requestID.Id)
	responseNum.Num = viewNum
	return nil
}

func (a *AppStoreHandler) AddComment(ctx context.Context, comment *appStore.RAppComment, response *appStore.Response) error {
	log.Info("接收到 appStore.AddComment 请求")
	appComment := &model.AppComment{}
	if err := common.SwapTo(comment, appComment); err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	err := a.AppStoreDataService.AddComment(appComment)
	if err != nil {
		logger.Error(err)
		response.Msg = err.Error()
		return err
	}
	response.Msg = "评论成功"
	return nil
}

func (a *AppStoreHandler) QueryAllCommentByID(ctx context.Context, requestID *appStore.RequestID, allAppComment *appStore.ResponseAllAppComment) error {
	allComment, err := a.AppStoreDataService.QueryAllCommentByID(requestID.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	for i, _ := range allComment {
		commentInfo := &appStore.RAppComment{}
		if err = common.SwapTo(allComment[i], commentInfo); err != nil {
			logger.Error(err)
			return err
		}
		allAppComment.AppComment = append(allAppComment.AppComment, commentInfo)
	}
	return nil
}
