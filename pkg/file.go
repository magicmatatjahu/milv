package pkg

import (
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

type Headers []string

type File struct {
	RelPath string `yaml:"path"`
	AbsPath string
	DirPath string
	Content string
	Links   Links `yaml:"links"`
	Headers Headers
	Status  bool
	Config  *FileConfig `yaml:"config"`
	Stats   *FileStats
	parser  *Parser
	valid   *Validation
}

func NewFile(filePath string, fileLinks Links, config *FileConfig) (*File, error) {
	if match, _ := regexp.MatchString(`.md$`, filePath); !match {
		return nil, errors.New("The specified file isn't a markdown file")
	}

	absPath, _ := filepath.Abs(filePath)
	if err := fileExists(absPath); err != nil {
		return nil, err
	}
	content, err := readMarkdown(absPath)
	if err != nil {
		return nil, err
	}

	return &File{
		RelPath: filePath,
		AbsPath: absPath,
		DirPath: filepath.Dir(filePath),
		Content: content,
		Links:   fileLinks,
		Config:  config,
		parser:  &Parser{},
		valid:   &Validation{},
	}, nil
}

func (f *File) Run() {
	f.ExtractLinks().
		ExtractHeaders().
		ValidateLinks().
		ExtractStats()
}

func (f *File) ExtractLinks() *File {
	whiteListExt, whiteListInt := []string{}, []string{}
	if f.Config != nil {
		whiteListExt = f.Config.WhiteListExt
		whiteListInt = f.Config.WhiteListInt
	}

	content := f.Content
	if f.Config != nil && !*f.Config.AllowCodeBlocks {
		content = removeCodeBlocks(content)
	}

	f.Links = f.parser.
		Links(content, f.DirPath).
		AppendConfig(f).
		RemoveWhiteLinks(whiteListExt, whiteListInt).
		Filter(func(link Link) bool {
			if f.Config != nil && f.Config.IgnoreInternal != nil && *f.Config.IgnoreInternal && (link.TypeOf == HashInternalLink || link.TypeOf == InternalLink) {
				return false
			}

			if f.Config != nil && f.Config.IgnoreExternal != nil && *f.Config.IgnoreExternal && link.TypeOf == ExternalLink {
				return false
			}

			return true
		})
	return f
}

func (f *File) ExtractHeaders() *File {
	f.Headers = f.parser.Headers(f.Content)
	return f
}

func (f *File) ValidateLinks() *File {
	f.Links = f.valid.Links(f.Links, f.Headers)
	f.Status = f.Links.CheckStatus()
	return f
}

func (f *File) ExtractStats() *File {
	f.Stats = NewFileStats(f)
	return f
}

func (f *File) WriteStats() *File {
	writeStats(f)
	return f
}

func (f *File) Summary() *File {
	summaryOfFile(f)
	return f
}
