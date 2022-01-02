package web

import (
	"go-app/library/web/middleware"
	"go-app/library/web/router"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

type Web struct {
	Engine    *gin.Engine
	BaseGroup *gin.RouterGroup
	Config    *WebConfig
}

type WebConfig struct {
	Port        string `yaml:"port" json:"port"`
	ContextPath string `yaml:"context_path" json:"context_path"`
}

func New(wc *WebConfig) *Web {
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.NoRoute(middleware.NoRoute)
	engine.NoMethod(middleware.NoMethod)
	engine.Use(middleware.AccessLog)
	// engine.Use(middleware.LimitMiddleware(100)) // limit concurrent request num
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
