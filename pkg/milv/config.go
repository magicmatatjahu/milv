package milv

import (
	"github.com/magicmatatjahu/milv/cli"
	"regexp"
	"time"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DocumentRoot	string	 			`yaml:"document-root"`
	ExternalPath	string 				`yaml:"external-path"`
	WhiteList		[]*regexp.Regexp 	`yaml:"white-list"`
	BlackList		[]*regexp.Regexp 	`yaml:"black-list"`
	Timeout        	time.Duration 		`yaml:"timeout"`
	Repeats 		int8     			`yaml:"repeats"`
	IgnoreExternal 	bool     			`yaml:"ignore-external"`
	IgnoreInternal 	bool     			`yaml:"ignore-internal"`
}

func newConfig(commands cli.Commands, config *Config) (*Config, error) {
	c := &Config{}

	err := c.readConfigFiles(commands.ConfigFile)
	if err != nil {
		return nil, err
	}

	return c.combine(commands), nil
}

func (c *Config) readConfigFiles(configFile string) error {
	configFiles := []string{"./milv.yaml", "./milv.yml"}

	if configFile != "" {
		configFiles = append([]string{configFile}, configFiles...)
	}

	for _, file := range configFiles {
		exists, err := c.readConfigFile(file)
		if err != nil {
			return err
		}
		if exists {
			break
		}
	}

	return nil
}

func (c *Config) readConfigFile(file string) (bool, error) {
	exists, err := fileExists(file)
	if err != nil {
		return exists, err
	}
	if !exists {
		return false, nil
	}

	data, err := readFile(file)
	if err != nil {
		return exists, err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return exists, err
	}

	return true, nil
}

func (c *Config) combine(commands cli.Commands) *Config {
	return c
}