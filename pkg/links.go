package pkg

import "strings"

type Links []Link

func (l Links) RemoveWhiteLinks(externals, internals []string) Links {
	links := l[:0]
	exist := false

	for _, link := range l {
		exist = false
		if link.TypeOf == ExternalLink {
			for _, w_link := range externals {
				if strings.Contains(link.AbsPath, w_link) {
					exist = true
					break
				}
			}
		} else {
			for _, w_link := range internals {
				if link.RelPath == w_link {
					exist = true
					break
				}
			}
		}
		if !exist {
			links = append(links, link)
		}
	}
	return links
}

func (l Links) CheckStatus() bool {
	for _, link := range l {
		if !link.Result.Status {
			return false
		}
	}
	return true
}