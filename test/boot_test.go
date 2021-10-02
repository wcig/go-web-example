package test

import (
	"go-app/library/boot"
	"go-app/library/log"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	wd, _ := os.Getwd()
	_ = os.Chdir(filepath.Dir(wd))
}

func initBootStarter() {
	boot.Register(&log.LogStarter{})
}

func testBootRun() {
	initBootStarter()
	boot.Run("dev")
}

func TestLog(t *testing.T) {
	testBootRun()
	log.Info("ok")
}
