package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sawyerwu/swuops/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitMysql()
}
