package main

import (
	"flag"
	_ "go-app/app/controller"
	"go-app/library/boot"
	"go-app/library/log"
	"go-app/library/web"
)

var profile = flag.String("profile", "dev", "profile: dev, test, prod")

func init() {
	boot.Register(&log.LogStarter{})
	boot.Register(&web.WebStarter{})
}

func main() {
	flag.Parse()
	boot.Run(*profile)
}
