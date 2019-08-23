package dict

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config is a structure that stores credentials for any API needed
type Config struct {
	APIKey string `yaml:"merriamWebsterKey"`
}

// ParseConfig parses the yaml config file at the specified file path
func ParseConfig(path string) (*Config, error) {
	conf := Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
