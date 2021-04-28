package initialize

import (
	"fmt"
	"github.com/sawyerwu/swuops/pkg/common"
	"github.com/spf13/viper"
	"os"
)

const (
	configType = "yml"
	configName = "application.yml"
	configPath = "/conf"
)

func InitConfig() {
	v := viper.New()
	workDir, _ := os.Getwd()
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	v.AddConfigPath(workDir + configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Init config file failed: %v", err))
	}

	// convert config file to struct
	if err := v.Unmarshal(&common.Conf); err != nil {
		panic(fmt.Sprintf("Init config file failed: %v", err))
	}

	fmt.Println("Init config file success")
}
