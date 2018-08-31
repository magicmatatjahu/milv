package pkg

import "fmt"

type FileConfig struct {
	WhiteListExt   []string `yaml:"white-list-external"`
	WhiteListInt   []string `yaml:"white-list-internal"`
	Timeout 	   int `yaml:"timeout"`
	ReguestTimes   int8 `yaml:"reguest-times"`
	IgnoreExternal bool `yaml:"ignore-external"`
	IgnoreInternal bool `yaml:"ignore-internal"`
}

func NewFileConfig(filePath string, config *Config) *FileConfig {
	for _, file := range config.Files {
		if filePath == file.RelPath {
			var timeout int
			if file.Config.Timeout != 0 {
				timeout = file.Config.Timeout
			} else {
				timeout = config.Timeout
			}

			var reguestTimes int8
			if file.Config.Timeout != 0 {
				reguestTimes = file.Config.ReguestTimes
			} else {
				reguestTimes = config.ReguestTimes
			}

			var ignoreExternal, ignoreInternal bool
			if file.Config.IgnoreExternal {
				ignoreExternal = file.Config.IgnoreExternal
			} else {
				ignoreExternal = config.IgnoreExternal
			}
			if file.Config.IgnoreInternal {
				ignoreInternal = file.Config.IgnoreInternal
			} else {
				ignoreInternal = config.IgnoreInternal
			}

			fmt.Println(ignoreExternal, ignoreInternal)

			return &FileConfig{
				WhiteListExt: unique(append(config.WhiteListExt, file.Config.WhiteListExt...)),
				WhiteListInt: unique(append(config.WhiteListInt, file.Config.WhiteListInt...)),
				Timeout: 		timeout,
				ReguestTimes:   reguestTimes,
				IgnoreExternal: ignoreExternal,
				IgnoreInternal: ignoreInternal,
			}
		}
	}
	return &FileConfig{
		WhiteListExt:   config.WhiteListExt,
		WhiteListInt:   config.WhiteListInt,
		Timeout: 		config.Timeout,
		ReguestTimes:   config.ReguestTimes,
		IgnoreExternal: config.IgnoreExternal,
		IgnoreInternal: config.IgnoreInternal,
	}
}
