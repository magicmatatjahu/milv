package milv

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Files			[]File		`yaml:"files"`
	WhiteListExt	[]string	`yaml:"white-list-external"`
	WhiteListInt	[]string	`yaml:"white-list-internal"`
	BlackList		[]string	`yaml:"black-list"`
	AllowDup		bool		`yaml:"allow-duplicate"`
	AllowSSL		bool		`yaml:"allow-ssl"`
	AllowRed		bool		`yaml:"allow-redirect"`
	AllowStatus		[]string	`yaml:"allow-status"`
}

func NewConfig(filePath string) (*Config, error) {
	config := &Config{}
	if err := fileExists(filePath); err != nil {
		return nil, err
	}

	return config.initConfig(filePath)
}

func (c *Config) initConfig(filePath string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}