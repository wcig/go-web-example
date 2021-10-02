package yaml

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func JsonDecodeOption() viper.DecoderConfigOption {
	return func(v *mapstructure.DecoderConfig) {
		v.TagName = "json"
	}
}

func YamlDecodeOption() viper.DecoderConfigOption {
	return func(v *mapstructure.DecoderConfig) {
		v.TagName = "yaml"
	}
}
