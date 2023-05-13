package service

import (
	"github.com/qyh794/go-paas/user/domain/model"
	"github.com/qyh794/go-paas/user/domain/repository"
	"github.com/qyh794/go-paas/user/pkg/jwt"
	"github.com/qyh794/go-paas/user/pkg/snowflake"
)

type IUserDataService interface {
	SignUp(user *model.User) error
	Login(user *model.User) (string, error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

func (u *UserDataService) SignUp(user *model.User) error {
	// 判断用户是否存在
	if err := u.UserRepository.UserIsExist(user.Username); err != nil {
		return err
	}
	// 不存在就创建用户
	// 生成用户ID
	user.ID = snowflake.GenID()
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) Login(user *model.User) (string, error) {
	if err := u.UserRepository.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.ID, user.Username)
}
