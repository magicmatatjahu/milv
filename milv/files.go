package milv

import (
	"regexp"
	"path/filepath"
	"strings"
)

type Files []*File

func NewFiles(filePaths []string, config *Config) (Files, error) {
	var files Files
	filePaths = removeBlackList(filePaths, config.BlackList)

	for _, filePath := range filePaths {
		if match, _ := regexp.MatchString(".md$", filePath); !match {
			continue
		}

		absPath, _ := filepath.Abs(filePath)
		file := &File{
			RelPath: absPath,
			AbsPath: absPath,
			DirPath: filepath.Dir(filePath),
			Config: NewFileConfig(filePath, config),
			parser: &Parser{},
			valid: &Validation{},
		}
		files = append(files, file)
	}
	return files, nil
}

func (f Files) ExtractStats() []FileStats {
	var fileStats []FileStats
	for _, file := range f {
		fileStats = append(fileStats, file.Stats)
	}
	return fileStats
}

func removeBlackList(filePaths, blackList []string) []string {
	var newFilePaths []string
	for _, file := range filePaths {
		exists := false
		for _, blackFile := range blackList {
			if strings.Contains(file, blackFile) {
				exists = true
				break
			}
		}
		if !exists {
			newFilePaths = append(newFilePaths, file)
		}
	}
	return newFilePaths
}