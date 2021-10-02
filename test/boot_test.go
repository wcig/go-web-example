package test

import (
	"fmt"
	"go-app/library/log"
	"go-app/library/util/json"
	"go-app/library/util/yaml"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

// func init() {
// 	wd, _ := os.Getwd()
// 	_ = os.Chdir(filepath.Dir(wd))
// }
//
// func testBootRun() {
// 	boot.Run("dev")
// }
//
// func testAppInit() {
// 	config.Init("dev")
// 	log.Init(config.Get().Logger)
// }

type Config struct {
	Logging log.LogConfig `yaml:"logging" json:"logging"`
}

func Test(t *testing.T) {
	var conf Config
	err := yaml.LoadFromFile(&conf, "../config/config-dev.yml")
	fmt.Println(err)
	json.PrintJsonPretty(conf)
}

func Test2(t *testing.T) {
	conf := viper.New()
	conf.SetConfigType("yaml")
	conf.SetConfigFile("../config/config-dev.yml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := conf.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	json.PrintJsonPretty(cfg)
}

func Test3(t *testing.T) {
	str := `{
  "logging": {
    "enabled": true,
    "path": "/tmp/log",
    "stdout": true,
    "access": true,
    "root": {
      "file_name": "info.log",
      "max_size": 1,
      "max_age": 2,
      "max_backups": 3,
      "compress": true,
      "level": "debug"
    }
  }
}`

	conf := viper.New()
	conf.SetConfigType("json")
	err := conf.ReadConfig(strings.NewReader(str))
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := conf.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	json.PrintJsonPretty(cfg)
}
