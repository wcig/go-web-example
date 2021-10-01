package web

import (
	"go-app/library/web/middleware"
	"go-app/library/web/router"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const configKey = "server"

type WebStarter struct {
	Web *Web
}

type WebConfig struct {
	Port        string `yaml:"port" json:"port"`
	ContextPath string `yaml:"context_path" json:"context_path"`
}

func (s *WebStarter) Init(cfg *viper.Viper) {
	info := cfg.Sub(configKey)
	if info == nil {
		panic("server config empty")
	}

	var wc WebConfig
	if err := info.Unmarshal(&wc); err != nil {
		panic(err)
	}

	web := New(&wc)
	s.Web = web
}

func (s *WebStarter) Start() {
	s.Web.RegisterRouters(router.GetRouters())
	s.Web.Listen()
}

type Web struct {
	Engine    *gin.Engine
	BaseGroup *gin.RouterGroup
	Config    *WebConfig
}

func New(wc *WebConfig) *Web {
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.NoRoute(middleware.NoRoute)
	engine.NoMethod(middleware.NoMethod)
	engine.Use(middleware.AccessLog)
	return &Web{
		Engine:    engine,
		BaseGroup: engine.Group(wc.ContextPath),
		Config:    wc,
	}
}

func (web *Web) RegisterRouters(routers []router.IRouter) {
	for _, v := range routers {
		v.Router(web.BaseGroup)
	}
}

func (web *Web) Listen() {
	addr := ":" + web.Config.Port
	err := gracehttp.Serve(&http.Server{Addr: addr, Handler: web.Engine})
	if err != nil {
		panic(err)
	}
}
