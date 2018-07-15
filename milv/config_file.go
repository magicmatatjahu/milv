package milv

import (
	"strings"
)

type FileConfig struct {
	WhiteListExt	[]string	`yaml:"white-list-external"`
	WhiteListInt	[]string	`yaml:"white-list-internal"`
	AllowDup		bool		`yaml:"allow-duplicate"`
	AllowSSL		bool		`yaml:"allow-ssl"`
	AllowRed		bool		`yaml:"allow-redirect"`
	AllowStatus		[]string	`yaml:"allow-status"`
}

func NewFileConfig(filePath string, config *Config) FileConfig {
	for _, file := range config.Files {
		if strings.Replace(filePath, "mds/", "", -1) == file.RelPath {
			return FileConfig{
				WhiteListExt: append(config.WhiteListExt, file.Config.WhiteListExt...),
				WhiteListInt: append(config.WhiteListInt, file.Config.WhiteListInt...),
			}
		}
	}
	return FileConfig{
		WhiteListExt: config.WhiteListExt,
		WhiteListInt: config.WhiteListInt,
	}
}