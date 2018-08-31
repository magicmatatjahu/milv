package pkg

import (
	"io/ioutil"

	"github.com/magicmatatjahu/milv/cli"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Files          []File   `yaml:"files"`
	WhiteListExt   []string `yaml:"white-list-external"`
	WhiteListInt   []string `yaml:"white-list-internal"`
	BlackList      []string `yaml:"black-list"`
	Timeout 	   int `yaml:"timeout"`
	ReguestTimes   int8 `yaml:"reguest-times"`
	IgnoreInternal bool
	IgnoreExternal bool
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
	var timeout int
	if commands.Timeout != 0 {
		timeout = commands.Timeout
	} else {
		timeout = c.Timeout
	}

	var reguestTimes int8
	if commands.Timeout != 0 {
		reguestTimes = commands.ReguestTimes
	} else {
		reguestTimes = c.ReguestTimes
	}

	return &Config{
		Files:          c.Files,
		WhiteListExt:   unique(append(c.WhiteListExt, commands.WhiteListExt...)),
		WhiteListInt:   unique(append(c.WhiteListInt, commands.WhiteListInt...)),
		BlackList:      unique(append(c.BlackList, commands.BlackList...)),
		Timeout: 		timeout,
		ReguestTimes:   reguestTimes,
		IgnoreInternal: commands.IgnoreInternal,
		IgnoreExternal: commands.IgnoreExternal,
	}
}
