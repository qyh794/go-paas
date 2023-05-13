package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/user/domain/model"
)

const secret = "abc.com"

type IUserRepository interface {
	InitTable() error
	CreateUser(*model.User) error
	Login(*model.User) error
	UserIsExist(string) error
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

func (u *UserRepository) CreateUser(user *model.User) error {
	password := user.Password
	user.Password = encryptPassword(password)
	return u.mysqlDB.Create(user).Error
}

func (u *UserRepository) Login(user *model.User) error {
	// 1.查询用户是否存在
	existingUser := &model.User{}
	if err := u.mysqlDB.Where("username = ?", user.Username).First(existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	// 2.将查询出来的用户的密码与用户输入密码进行对比
	if existingUser.Password != encryptPassword(user.Password) {
		return errors.New("wrong password")
	}
	return nil
}

func (u *UserRepository) UserIsExist(username string) error {
	// 1.查询用户是否存在
	existingUser := &model.User{}
	if err := u.mysqlDB.Where("username = ?", username).First(existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return errors.New("user already exist")
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
