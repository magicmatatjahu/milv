package pkg

type LinkType string

const (
	ExternalLink     LinkType = "ExternalLink"
	InternalLink     LinkType = "InternalLink"
	HashInternalLink LinkType = "HashInternalLink"
)

type Link struct {
	RelPath string `yaml:"path"`
	AbsPath string
	Config  *LinkConfig `yaml:"config"`
	TypeOf  LinkType
	Result  LinkResult
}

type LinkResult struct {
	Status  bool
	Message string
}
