package main

import (
	"embed"
	"flag"
	_ "go-app/app/controller"
	"go-app/library/boot"
	"go-app/library/log"
	"go-app/library/web"
)

var profile = flag.String("profile", "dev", "profile: dev, test, prod")

//go:embed config
var cfs embed.FS

func init() {
	boot.Register(&log.LogStarter{})
	boot.Register(&web.WebStarter{})
}

func main() {
	flag.Parse()
	boot.Run(cfs, *profile)
}
