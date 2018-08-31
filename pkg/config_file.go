package pkg

type FileConfig struct {
	WhiteListExt   []string `yaml:"white-list-external"`
	WhiteListInt   []string `yaml:"white-list-internal"`
	Timeout        *int     `yaml:"timeout"`
	ReguestRepeats *int8    `yaml:"request-repeats"`
	AllowRedirect  *bool	`yaml:"allow-redirect"`
	AllowCodeBlocks *bool 	`yaml:"allow-code-blocks"`
	IgnoreExternal *bool    `yaml:"ignore-external"`
	IgnoreInternal *bool    `yaml:"ignore-internal"`
}

func NewFileConfig(filePath string, config *Config) *FileConfig {
	if config != nil {
		for _, file := range config.Files {
			if filePath == file.RelPath && file.Config != nil {
				var timeout *int
				if file.Config.Timeout != nil {
					timeout = file.Config.Timeout
				} else {
					timeout = &config.Timeout
				}

				var requestRepeats *int8
				if file.Config.Timeout != nil {
					requestRepeats = file.Config.ReguestRepeats
				} else {
					requestRepeats = &config.ReguestRepeats
				}

				var allowRedirect, allowCodeBlocks, ignoreExternal, ignoreInternal *bool
				if file.Config.AllowCodeBlocks != nil {
					allowCodeBlocks = file.Config.AllowCodeBlocks
				} else {
					allowCodeBlocks = &config.AllowCodeBlocks
				}
				if file.Config.AllowRedirect != nil {
					allowRedirect = file.Config.AllowRedirect
				} else {
					allowRedirect = &config.AllowRedirect
				}
				if file.Config.IgnoreExternal != nil {
					ignoreExternal = file.Config.IgnoreExternal
				} else {
					ignoreExternal = &config.IgnoreExternal
				}
				if file.Config.IgnoreInternal != nil {
					ignoreInternal = file.Config.IgnoreInternal
				} else {
					ignoreInternal = &config.IgnoreInternal
				}

				return &FileConfig{
					WhiteListExt:   unique(append(config.WhiteListExt, file.Config.WhiteListExt...)),
					WhiteListInt:   unique(append(config.WhiteListInt, file.Config.WhiteListInt...)),
					Timeout:        timeout,
					ReguestRepeats: requestRepeats,
					AllowRedirect:  allowRedirect,
					AllowCodeBlocks: allowCodeBlocks,
					IgnoreExternal: ignoreExternal,
					IgnoreInternal: ignoreInternal,
				}
			}
		}
		return &FileConfig{
			WhiteListExt:   config.WhiteListExt,
			WhiteListInt:   config.WhiteListInt,
			Timeout:        &config.Timeout,
			ReguestRepeats: &config.ReguestRepeats,
			AllowRedirect:  &config.AllowRedirect,
			AllowCodeBlocks: &config.AllowCodeBlocks,
			IgnoreExternal: &config.IgnoreExternal,
			IgnoreInternal: &config.IgnoreInternal,
		}
	}
	return nil
}
