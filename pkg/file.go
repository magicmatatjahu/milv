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

func NewFile(filePath string, config *FileConfig) (*File, error) {
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
		Config:  config,
		parser:  &Parser{},
		valid:   &Validation{},
	}, nil
}

func (f *File) Run() {
	f.ExtractLinks()
	f.ExtractHeaders()
	f.ValidateLinks()
	f.ExtractStats()
}

func (f *File) ExtractLinks() {
	f.Links = f.parser.
		Links(f.Content, f.DirPath).
		RemoveWhiteLinks(f.Config.WhiteListExt, f.Config.WhiteListInt).
		Filter(func(link Link) bool {
			if f.Config.IgnoreInternal && (link.TypeOf == HashInternalLink || link.TypeOf == InternalLink) {
				return false
			}

			if f.Config.IgnoreExternal && link.TypeOf == ExternalLink {
				return false
			}

			return true
		})
}

func (f *File) ExtractHeaders() {
	f.Headers = f.parser.Headers(f.Content)
}

func (f *File) ValidateLinks() {
	f.Links = f.valid.Links(f.Links, f.Headers)
	f.Status = f.Links.CheckStatus()
}

func (f *File) ExtractStats() {
	f.Stats = NewFileStats(f)
}

func (f *File) WriteStats() {
	writeStats(f)
}

func (f *File) Summary() {
	summaryOfFile(f)
}
