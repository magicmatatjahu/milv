package milv

type FileStats struct {
	SucessedExternalLinks		SucessedExternalLink
	SucessedInternalLinks		SucessedInternalLink
	FailedExternalLinks			FailedExternalLink
	FailedInternalLinks			FailedInternalLink
}

type SucessedExternalLink struct {
	Count	int
	Links	[]Link
}

type SucessedInternalLink struct {
	Count	int
	Links	[]Link
}

type FailedExternalLink struct {
	Count	int
	Links	[]Link
}

type FailedInternalLink struct {
	Count	int
	Links	[]Link
}