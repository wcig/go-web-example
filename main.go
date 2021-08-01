package main

import (
	"flag"
	"go-app/boot"
)

var profile = flag.String("profile", "dev", "profile: dev, test, prod")

func main() {
	flag.Parse()
	boot.Run(*profile)
}
