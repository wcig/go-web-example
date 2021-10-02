package log

import (
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "logging"

type LogStarter struct {
	LogConfig *LogConfig
}

func (s *LogStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("log config empty")
	}

	var lc LogConfig
	if err := info.Unmarshal(&lc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.LogConfig = &lc
}

func (s *LogStarter) Start() {
	Init(s.LogConfig)
}
