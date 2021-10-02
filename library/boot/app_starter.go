package boot

import (
	"go-app/library/util/yaml"

	"github.com/spf13/viper"
)

func init() {
	Register(&AppStarter{})
}

const configKey = "app"

type AppStarter struct {
	Config *AppConfig
}

var ac AppConfig

type AppConfig struct {
	Name string `yaml:"name" json:"name"`
}

func (s *AppStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("app config empty")
	}

	if err := info.Unmarshal(&ac, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}
	s.Config = &ac
}

func (s *AppStarter) Start() {}

func GetAppConfig() *AppConfig {
	return &ac
}
