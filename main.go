package main

import (
	"flag"
	"go-app/library/boot"
	"go-app/library/log"
)

var profile = flag.String("profile", "dev", "profile: dev, test, prod")

func init() {
	boot.Register(&log.LogStarter{})
}

func main() {
	flag.Parse()
	boot.Run(*profile)
}
