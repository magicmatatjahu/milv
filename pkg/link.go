package pkg

type LinkType string

const (
	ExternalLink     LinkType = "ExternalLink"
	InternalLink     LinkType = "InternalLink"
	HashInternalLink LinkType = "HashInternalLink"
)

type Link struct {
	RelPath string
	AbsPath string
	Timeout int8
	ReguestTimes int8
	TypeOf  LinkType
	Result  LinkResult
}

type LinkResult struct {
	Status  bool
	Message string
}
