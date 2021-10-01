package starter

import "github.com/spf13/viper"

type Starter interface {
	Init(cfg *viper.Viper)
	Start(cfg *viper.Viper)
}
