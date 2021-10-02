package web

import (
	"go-app/library/util/yaml"
	"go-app/library/web/router"

	"github.com/spf13/viper"
)

const configKey = "server"

type WebStarter struct {
	Web *Web
}

func (s *WebStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("server config empty")
	}

	var wc WebConfig
	if err := info.Unmarshal(&wc, yaml.YamlDecodeOption()); err != nil {
		panic(err)
	}

	web := New(&wc)
	s.Web = web
}

func (s *WebStarter) Start() {
	s.Web.RegisterRouters(router.GetRouters())
	s.Web.Listen()
}
