package common

import (
	"github.com/asim/go-micro/v3/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

func GetMysqlConfigFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}

func InitMySQL(dialect string, mysqlConfig *MysqlConfig) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, mysqlConfig.User+":"+mysqlConfig.Password+"@("+mysqlConfig.Host+":3306)/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, err
}
