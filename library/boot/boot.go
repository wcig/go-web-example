package boot

import (
	"go-app/library/config"
	"go-app/library/starter"

	"github.com/spf13/viper"
)

var bs BootStarter

type BootStarter struct {
	Config   *viper.Viper
	Starters []starter.Starter
}

func (bs *BootStarter) initStarters() {
	for _, s := range bs.Starters {
		s.Init(bs.Config)
	}
}

func (bs *BootStarter) startStarters() {
	for _, s := range bs.Starters {
		s.Start()
	}
}

func Register(starter starter.Starter) {
	if starter == nil {
		panic("starter empty")
	}
	bs.Starters = append(bs.Starters, starter)
}

func Run(profile string) {
	cfg := config.Init(profile)
	bs.Config = cfg
	bs.initStarters()
	bs.startStarters()
}
