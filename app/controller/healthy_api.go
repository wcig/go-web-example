package controller

import (
	"go-app/library/log"
	"go-app/library/response"
	"go-app/library/web/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.RegisterRouter(&HealthyController{})
}

type HealthyController struct{}

func (ctrl *HealthyController) Router(rg *gin.RouterGroup) {
	rg.GET("/ready", ctrl.ready)
}

func (ctrl *HealthyController) ready(c *gin.Context) {
	log.Info("ready")
	log.Debug("ready")
	log.Error("ready")
	response.Ok(c, nil)
}
