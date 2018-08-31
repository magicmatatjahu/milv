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
	Timeout        int      `yaml:"timeout"`
	ReguestRepeats int8     `yaml:"request-repeats"`
	AllowRedirect  bool		`yaml:"allow-redirect"`
	AllowCodeBlocks bool 	`yaml:"allow-code-blocks"`
	IgnoreExternal bool     `yaml:"ignore-external"`
	IgnoreInternal bool     `yaml:"ignore-internal"`
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
	if commands.FlagsSet["timeout"] {
		timeout = commands.Timeout
	} else {
		timeout = c.Timeout
	}

	var requestRepeats int8
	if commands.FlagsSet["request-repeats"] {
		requestRepeats = commands.ReguestRepeats
	} else {
		requestRepeats = c.ReguestRepeats
	}

	var allowRedirect, allowCodeBlocks, ignoreExternal, ignoreInternal bool
	if commands.FlagsSet["allow-redirect"] {
		allowRedirect = commands.AllowRedirect
	} else {
		allowRedirect = c.AllowRedirect
	}
	if commands.FlagsSet["allow-code-blocks"] {
		allowCodeBlocks = commands.AllowCodeBlocks
	} else {
		allowCodeBlocks = c.AllowCodeBlocks
	}
	if commands.FlagsSet["ignore-external"] {
		ignoreExternal = commands.IgnoreExternal
	} else {
		ignoreExternal = c.IgnoreExternal
	}
	if commands.FlagsSet["ignore-internal"] {
		ignoreInternal = commands.IgnoreInternal
	} else {
		ignoreInternal = c.IgnoreInternal
	}

	return &Config{
		Files:          c.Files,
		WhiteListExt:   unique(append(c.WhiteListExt, commands.WhiteListExt...)),
		WhiteListInt:   unique(append(c.WhiteListInt, commands.WhiteListInt...)),
		BlackList:      unique(append(c.BlackList, commands.BlackList...)),
		Timeout:        timeout,
		ReguestRepeats: requestRepeats,
		AllowRedirect:  allowRedirect,
		AllowCodeBlocks: allowCodeBlocks,
		IgnoreExternal: ignoreExternal,
		IgnoreInternal: ignoreInternal,
	}
}
