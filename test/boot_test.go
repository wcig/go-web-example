package test

import (
	"go-app/boot"
	"go-app/library/config"
	"go-app/library/log"
	"os"
	"path/filepath"
)

func init() {
	wd, _ := os.Getwd()
	_ = os.Chdir(filepath.Dir(wd))
}

func testBootRun() {
	boot.Run("dev")
}

func testAppInit() {
	config.Init("dev")
	log.Init(config.Get().Logger)
}
