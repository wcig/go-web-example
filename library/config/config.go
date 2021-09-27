package config

import (
	"go-app/library/log"
	"go-app/library/util/json"
	"go-app/library/util/yaml"
	"path/filepath"
)

type Config struct {
	App    *AppCfg        `yaml:"app" json:"app"`
	Server *ServerCfg     `yaml:"server" json:"server"`
	Logger *log.LoggerCfg `yaml:"logger" json:"logger"`
}

type AppCfg struct {
	Name    string `yaml:"name" json:"name"`
	Profile string `yaml:"profile" json:"profile"`
}

type ServerCfg struct {
	Port        string `yaml:"port" json:"port"`
	ContextPath string `yaml:"context_path" json:"context_path"`
}

var _cfg Config

func Init(profile string) {
	filePath := filepath.Join("config", "config-"+profile+".yml")
	if err := yaml.LoadFromFile(&_cfg, filePath); err != nil {
		panic(err)
	}
	json.PrintJsonPretty(_cfg)
}

func Get() *Config {
	return &_cfg
}
