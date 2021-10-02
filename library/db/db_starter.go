package db

import (
	"fmt"
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "db"

type DBStarter struct {
	Config *DBConfig
}

func (s *DBStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("log config empty")
	}

	var dc DBConfig
	if err := info.Unmarshal(&dc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &dc
	fmt.Println("init db:", info.AllSettings())
}

func (s *DBStarter) Start() {
	Init(s.Config)
}
