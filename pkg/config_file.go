package pkg

type FileConfig struct {
	WhiteListExt   []string `yaml:"white-list-external"`
	WhiteListInt   []string `yaml:"white-list-internal"`
	IgnoreExternal bool
	IgnoreInternal bool
}

func NewFileConfig(filePath string, config *Config) *FileConfig {
	for _, file := range config.Files {
		if filePath == file.RelPath {
			return &FileConfig{
				WhiteListExt: unique(append(config.WhiteListExt, file.Config.WhiteListExt...)),
				WhiteListInt: unique(append(config.WhiteListInt, file.Config.WhiteListInt...)),
			}
		}
	}
	return &FileConfig{
		WhiteListExt:   config.WhiteListExt,
		WhiteListInt:   config.WhiteListInt,
		IgnoreExternal: config.IgnoreExternal,
		IgnoreInternal: config.IgnoreInternal,
	}
}
