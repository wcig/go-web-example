package rocketmq

import (
	"fmt"
	"go-app/library/util/json"
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "rocketmq"

type RocketMQStarter struct {
	Config *RocketMQConfig
}

func (s *RocketMQStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("rocketmq config empty")
	}

	var rc RocketMQConfig
	if err := info.Unmarshal(&rc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &rc
	fmt.Println("init rocketmq:", json.ToJson(rc))
}

func (s *RocketMQStarter) Start() {
	Init(s.Config)
}
