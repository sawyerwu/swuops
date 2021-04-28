package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sawyerwu/swuops/pkg/common"
)

func InitMysql() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		common.Conf.Mysql.Username,
		common.Conf.Mysql.Password,
		common.Conf.Mysql.Host,
		common.Conf.Mysql.Port,
		common.Conf.Mysql.Database,
		common.Conf.Mysql.Query)
	fmt.Println(args)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// db.SingularTable(true)
	common.Mysql = db

	fmt.Println("Init db connection success")
}
