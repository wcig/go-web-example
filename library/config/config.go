package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

// type Config struct {
// 	App    *AppCfg        `yaml:"app" json:"app"`
// 	Server *ServerCfg     `yaml:"server" json:"server"`
// 	Logger *log.LogConfig `yaml:"logger" json:"logger"`
// }
//
// type AppCfg struct {
// 	Name    string `yaml:"name" json:"name"`
// 	Profile string `yaml:"profile" json:"profile"`
// }
//
// type ServerCfg struct {
// 	Port        string `yaml:"port" json:"port"`
// 	ContextPath string `yaml:"context_path" json:"context_path"`
// }
//
// var _cfg Config
//
// func Init(profile string) {
// 	filePath := filepath.Join("config", "config-"+profile+".yml")
// 	if err := yaml.LoadFromFile(&_cfg, filePath); err != nil {
// 		panic(err)
// 	}
// 	json.PrintJsonPretty(_cfg)
// }
//
// func Get() *Config {
// 	return &_cfg
// }

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
