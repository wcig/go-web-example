package router

import "github.com/gin-gonic/gin"

type IRouter interface {
	Router(*gin.RouterGroup)
}

var _Routers []IRouter

func RegisterRouter(router IRouter) {
	_Routers = append(_Routers, router)
}

func GetRouters() []IRouter {
	return _Routers
}
