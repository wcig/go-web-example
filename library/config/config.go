package config

import (
	"bytes"
	"embed"
	"io/ioutil"
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

func TestInit(profile string) *viper.Viper {
	filePath := filepath.Join("config", "config-"+profile+".yml")
	data, err := ioutil.ReadFile(filePath)
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
