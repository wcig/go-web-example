package yaml

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func LoadFromFile(v interface{}, filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, v)
}
