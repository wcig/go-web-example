package redis

import (
	"fmt"
	"go-app/library/util/json"
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "redis"

type RedisStarter struct {
	Config *RedisConfig
}

func (s *RedisStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("redis config empty")
	}

	var rc RedisConfig
	if err := info.Unmarshal(&rc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &rc
	fmt.Println("init redis:", json.ToJson(rc))
}

func (s *RedisStarter) Start() {
	Init(s.Config)
}
