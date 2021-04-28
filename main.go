package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sawyerwu/swuops/initialize"
	"os"
	"os/signal"
)

func main() {
	initialize.InitConfig()
	initialize.InitMysql()
	r := initialize.InitRouters()

	r.Run(":8080")

	quit := make(chan os.Signal)
	// listening on signal and block here
	signal.Notify(quit, os.Interrupt, os.Kill)
	s := <-quit
	fmt.Println("signal: ", s)

}
