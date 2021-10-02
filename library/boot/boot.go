package boot

import (
	"embed"
	"go-app/library/config"
	"go-app/library/starter"

	"github.com/spf13/viper"
)

var boot Boot

type Boot struct {
	Config   *viper.Viper
	Starters []starter.Starter
}

func (b *Boot) initStarters() {
	for _, s := range b.Starters {
		s.Init(b.Config)
	}
}

func (b *Boot) startStarters() {
	for _, s := range b.Starters {
		s.Start()
	}
}

func Register(starter starter.Starter) {
	if starter == nil {
		panic("starter empty")
	}
	boot.Starters = append(boot.Starters, starter)
}

func Run(cfs embed.FS, profile string) {
	cfg := config.Init(cfs, profile)
	boot.Config = cfg
	boot.initStarters()
	boot.startStarters()
}

func TestRun(profile string) {
	cfg := config.TestInit(profile)
	boot.Config = cfg
	boot.initStarters()
	boot.startStarters()
}
