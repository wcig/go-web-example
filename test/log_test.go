package test

import (
	"go-app/library/boot"
	"go-app/library/log"
	"testing"
)

func initLog() {
	boot.Register(&log.LogStarter{})
	testBootRun()
}

func TestLog(t *testing.T) {
	initLog()
	log.Info("ok")
}
