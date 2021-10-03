package mongo

import (
	"fmt"
	"go-app/library/util/json"
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

const configKey = "mongo"

type MongoStarter struct {
	Config *MongoConfig
}

func (s *MongoStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("mongo config empty")
	}

	var mc MongoConfig
	if err := info.Unmarshal(&mc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &mc
	fmt.Println("init redis:", json.ToJson(mc))
}

func (s *MongoStarter) Start() {
	Init(s.Config)
}
