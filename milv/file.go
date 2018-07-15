package milv

import (
	"path/filepath"
	"regexp"
	"strings"
)

type File struct {
	RelPath		string			`yaml:"path"`
	AbsPath		string
	DirPath		string
	Links		[]Link
	Headers		[]string
	Status		bool
	Config 		FileConfig		`yaml:"config"`
	Stats		FileStats
	parser 		*Parser
	valid 		*Validation
}

func NewFile(filePath string) (*File, error) {
	if match, _ := regexp.MatchString(".md$", filePath); !match {
		return nil, nil
	}
	absPath, _ := filepath.Abs(filePath)

	return &File{
		RelPath: filePath,
		AbsPath: absPath,
		DirPath: filepath.Dir(filePath),
		Config: FileConfig{},
		parser: &Parser{},
		valid: &Validation{},
	}, nil
}

func (f *File) ExtractLinks() error {
	extractedLinks, err := f.parser.GetAndExtractLinks(f.DirPath, f.AbsPath)
	if err != nil {
		return err
	}

	f.Links = extractedLinks
	f.Links = f.removeWhiteLinks()

	return nil
}

func (f *File) ExtractHeaders() error {
	extractedHeaders, err := f.parser.GetAllHeaders(f.AbsPath)
	if err != nil {
		return err
	}

	f.Headers = extractedHeaders
	return nil
}

func (f *File) ValidateLinks() error {
	validatedLinks, err := f.valid.ValidateLinks(f.Links, f.Headers)
	if err != nil {
		return err
	}

	f.Links = validatedLinks
	f.Status = f.CheckValidatedLinks()
	return nil
}

func (f *File) ExtractStats() {
	var fileStat FileStats
	for _, link := range f.Links {
		if link.Result.Status {
			if link.TypeOf == ExternalLink {
				fileStat.SucessedExternalLinks.Count++
				fileStat.SucessedExternalLinks.Links = append(fileStat.SucessedExternalLinks.Links, link)
			} else if link.TypeOf == InternalLink {
				fileStat.SucessedInternalLinks.Count++
				fileStat.SucessedInternalLinks.Links = append(fileStat.SucessedInternalLinks.Links, link)
			}
		} else {
			if link.TypeOf == ExternalLink {
				fileStat.FailedExternalLinks.Count++
				fileStat.FailedExternalLinks.Links = append(fileStat.FailedExternalLinks.Links, link)
			} else if link.TypeOf == InternalLink {
				fileStat.FailedInternalLinks.Count++
				fileStat.FailedInternalLinks.Links = append(fileStat.FailedInternalLinks.Links, link)
			}
		}
	}
	f.Stats = fileStat
}

func (f *File) CheckValidatedLinks() bool {
	for _, link := range f.Links {
		if !link.Result.Status {
			return false
		}
	}
	return true
}

func (f *File) removeWhiteLinks() []Link {
	//TODO: Improve that shit
	tmp := f.Links[:0]
	exist := false
	for _, link := range f.Links {
		exist = false
		if link.TypeOf == ExternalLink {
			for _, w_link := range f.Config.WhiteListExt {
				if strings.Contains(link.AbsPath, w_link) {
					exist = true
					break
				}
			}
		} else {
			for _, w_link := range f.Config.WhiteListInt {
				if link.RelPath == w_link {
					exist = true
					break
				}
			}
		}
		if !exist {
			tmp = append(tmp, link)
		}
	}
	return tmp
}