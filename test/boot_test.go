package test

import (
	"go-app/library/boot"
	"os"
	"path/filepath"
)

func init() {
	wd, _ := os.Getwd()
	_ = os.Chdir(filepath.Dir(wd))
}

func testBootRun() {
	boot.TestRun("dev")
}
