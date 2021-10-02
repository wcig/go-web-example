package config

import (
	"bytes"
	"embed"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init(cfs embed.FS, profile string) *viper.Viper {
	filePath := filepath.Join("config", "config-"+profile+".yml")
	data, err := cfs.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	conf := viper.New()
	conf.SetConfigType("yaml")
	if err = conf.ReadConfig(bytes.NewReader(data)); err != nil {
		panic(err)
	}
	return conf
}
