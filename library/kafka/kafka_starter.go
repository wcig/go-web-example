package kafka

import (
	"fmt"
	"go-app/library/util/json"
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "kafka"

type KafkaStarter struct {
	Config *KafkaConfig
}

func (s *KafkaStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("kafka config empty")
	}

	var dc KafkaConfig
	if err := info.Unmarshal(&dc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &dc
	fmt.Println("init kafka:", json.ToJson(dc))
}

func (s *KafkaStarter) Start() {
	Init(s.Config)
}
