package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	APIPort     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_ADDRESS  string
	DB_PORT     string
	APIKey      string
	TokenSecret string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	_ = viper.Unmarshal(cfg)
	Cfg = cfg
	fmt.Println(Cfg)
}
