package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qyh794/go-paas/podApi/proto/podApi"
	"net/http"
)

func ResponseSucceed(ctx context.Context, msg interface{}, response *podApi.Response) error {
	response.StatusCode = http.StatusOK
	bytes, _ := json.Marshal(msg)
	response.Body = string(bytes)
	return nil
}

func ResponseAuthFailed(ctx context.Context, err error, response *podApi.Response) error {
	response.StatusCode = http.StatusNetworkAuthenticationRequired
	response.Body = err.Error()
	return err
}

func ResponseBadRequest(ctx context.Context, err error, response *podApi.Response) error {
	response.StatusCode = http.StatusBadRequest
	response.Body = err.Error()
	return err
}

func ResponseServerError(ctx context.Context, err error, response *podApi.Response) error {
	response.StatusCode = http.StatusInternalServerError
	response.Body = err.Error()
	return err
}

func ResponseEmptyToken(ctx context.Context, response *podApi.Response) error {
	response.StatusCode = http.StatusNetworkAuthenticationRequired
	response.Body = errors.New("empty token").Error()
	return errors.New("empty token")
}
