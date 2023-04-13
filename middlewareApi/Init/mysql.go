package Init

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/common"
)

func InitMySQL(dialect string, mysqlConfig *common.MysqlConfig) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, mysqlConfig.User+":"+mysqlConfig.Password+"@("+mysqlConfig.Host+":3306)/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, err
}
