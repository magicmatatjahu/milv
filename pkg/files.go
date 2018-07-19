package pkg

type Files []*File

func NewFiles(filePaths []string, config *Config) (Files, error) {
	var files Files

	filePaths = removeBlackList(filePaths, config.BlackList)
	for _, filePath := range filePaths {
		file, err := NewFile(filePath, NewFileConfig(filePath, config))
		if err != nil {
			return Files{}, err
		}
		files = append(files, file)
	}
	return files, nil
}

func (f Files) Run() {
	for _, file := range f {
		file.Run()
		file.WriteStats()
	}
}

func (f Files) Summary() bool {
	return summaryOfFiles(f)
}