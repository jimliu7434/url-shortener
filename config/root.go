package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Root *RootConfig

type RootConfig struct {
	Server      ServerConfig      `mapstructure:"server"`
	Application ApplicationConfig `mapstructure:"app"`
	Redis       RedisConfig       `mapstructure:"redis"`
}

// Setup 初始化設定
func Setup(configFileType string, configFilePath string) {
	// 初始化 viper
	root := viper.New()
	root.SetConfigType(configFileType)
	root.SetConfigFile(configFilePath)

	if err := root.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("讀取 root 設定檔失敗，請檢查 %s 檔案是否存在: %v", configFilePath, err))
	}

	Root = &RootConfig{}
	if err := root.Unmarshal(Root); err != nil {
		panic(fmt.Sprintf("root 設定檔格式有誤，請檢查 %s 檔案規格是否正確: %v", configFilePath, err))
	}
}
