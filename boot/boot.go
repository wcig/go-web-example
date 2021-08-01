package boot

import (
	"go-app/library/config"
	"go-app/library/log"
	"go-app/library/web"

	_ "go-app/app/controller"
)

func Run(profile string) {
	config.Init(profile)
	log.Init(config.Get().Logger)
	web.Start()
}
