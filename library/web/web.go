package web

import (

	// "go-app/app"

	"go-app/library/config"
	"go-app/library/web/middleware"
	"go-app/library/web/router"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	// "github.com/vskit-tv/vcomm-go/util/client"
	// "github.com/vskit-tv/vcomm-go/base/verr"
	// "github.com/vskit-tv/vcomm-go/model"
	// "github.com/facebookgo/grace/gracehttp"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-contrib/gzip"
	// "github.com/gin-gonic/gin"
	// "github.com/vskit-tv/vcomm-go/framework/gin/web"
	// "github.com/vskit-tv/vcomm-go/framework/web/auth"
)

type Web struct {
	Config    *config.ServerCfg
	Engine    *gin.Engine
	BaseGroup *gin.RouterGroup
}

func Start() {
	web := New()
	web.RegisterRouters(router.GetRouters())
	web.Listen()
}

func New() *Web {
	cfg := config.Get()
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.NoRoute(middleware.NoRoute)
	engine.NoMethod(middleware.NoMethod)
	engine.Use(middleware.AccessLog)
	return &Web{
		Config:    cfg.Server,
		Engine:    engine,
		BaseGroup: engine.Group(cfg.Server.ContextPath),
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
