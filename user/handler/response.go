package handler

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/user/proto/user"
)

func ResponseLogin(response *user.ResponseLogin, err error, msg string) error {
	response.Msg = "用户认证为: " + msg
	return err
}

func ResponseSignUp(response *user.ResponseSignUp) error {
	response.Msg = "注册成功"
	return nil
}

func ResponseError(response *user.ResponseLogin, err error) error {
	logger.Error(err)
	return err
}

/*


	if err != nil {
		return err
	}
	// json返回pod信息
	response.StatusCode = succeed
	bytes, _ := json.Marshal(podInfo)
	response.Body = string(bytes)
	return nil

*/
