package pkg

import (
	"github.com/magicmatatjahu/milv/cli"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Files			[]File		`yaml:"files"`
	WhiteListExt	[]string	`yaml:"white-list-external"`
	WhiteListInt	[]string	`yaml:"white-list-internal"`
	BlackList		[]string	`yaml:"black-list"`
}

func NewConfig(commands cli.Commands) (*Config, error) {
	config := &Config{}

	err := fileExists(commands.ConfigFile)
	if commands.ConfigFile != "milv.config.yaml" && err != nil {
		return nil, err
	}
	if err == nil {
		yamlFile, err := ioutil.ReadFile(commands.ConfigFile)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(yamlFile, config)
		if err != nil {
			return nil, err
		}
	}
	return config.combine(commands), nil
}

func (c *Config) combine(commands cli.Commands) *Config {
	return &Config{
		Files: c.Files,
		WhiteListExt: unique(append(c.WhiteListExt, commands.WhiteListExt...)),
		WhiteListInt: unique(append(c.WhiteListInt, commands.WhiteListInt...)),
		BlackList: unique(append(c.BlackList, commands.BlackList...)),
	}
}