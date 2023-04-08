package Init

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qyh794/go-paas/common"
)

var DB *gorm.DB

func SetupMysql(dialect string, mysqlConfig *common.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)
	DB, err = gorm.Open(dialect, dsn)
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}

func Close() {
	_ = DB.Close()
}
