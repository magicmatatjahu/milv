package pkg

type LinkConfig struct {
	Timeout        *int  `yaml:"timeout"`
	ReguestRepeats *int8 `yaml:"request-repeats"`
	AllowRedirect  *bool `yaml:"allow-redirect"`
}

func NewLinkConfig(link Link, file *File) *LinkConfig {
	if file.Config != nil {
		for _, linkFile := range file.Links {
			if (link.RelPath == linkFile.RelPath || link.AbsPath == linkFile.RelPath) && linkFile.Config != nil {
				var timeout *int
				if linkFile.Config.Timeout != nil {
					timeout = linkFile.Config.Timeout
				} else {
					timeout = file.Config.Timeout
				}

				var requestRepeats *int8
				if linkFile.Config.ReguestRepeats != nil {
					requestRepeats = linkFile.Config.ReguestRepeats
				} else {
					requestRepeats = file.Config.ReguestRepeats
				}

				var allowRedirect *bool
				if linkFile.Config.AllowRedirect != nil {
					allowRedirect = linkFile.Config.AllowRedirect
				} else {
					allowRedirect = file.Config.AllowRedirect
				}

				return &LinkConfig{
					Timeout:        timeout,
					ReguestRepeats: requestRepeats,
					AllowRedirect:  allowRedirect,
				}
			}
		}
		return &LinkConfig{
			Timeout:        file.Config.Timeout,
			ReguestRepeats: file.Config.ReguestRepeats,
			AllowRedirect:	file.Config.AllowRedirect,
		}
	}
	return nil
}
