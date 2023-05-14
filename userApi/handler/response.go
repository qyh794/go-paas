package handler

import (
	"context"
	"encoding/json"
	"github.com/qyh794/go-paas/userApi/proto/userApi"
	"net/http"
)

func ResponseSucceed(ctx context.Context, msg interface{}, response *userApi.Response) error {
	response.StatusCode = http.StatusOK
	bytes, _ := json.Marshal(msg)
	response.Body = string(bytes)
	return nil
}

func ResponseAuthFailed(ctx context.Context, err error, response *userApi.Response) error {
	response.StatusCode = http.StatusNetworkAuthenticationRequired
	response.Body = err.Error()
	return err
}

func ResponseBadRequest(ctx context.Context, err error, response *userApi.Response) error {
	response.StatusCode = http.StatusBadRequest
	response.Body = err.Error()
	return err
}

func ResponseServerError(ctx context.Context, err error, response *userApi.Response) error {
	response.StatusCode = http.StatusInternalServerError
	response.Body = err.Error()
	return err
}
