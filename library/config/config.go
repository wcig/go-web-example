package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func Init(profile string) *viper.Viper {
	filePath := filepath.Join("config", "config-"+profile+".yml")
	conf := viper.New()
	conf.SetConfigType("yaml")
	conf.SetConfigFile(filePath)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
