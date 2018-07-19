package pkg

import (
	"regexp"
	"bufio"
	"strings"
	"path"
	"fmt"
)

type Parser struct {}

type match func([][]string) string

const (
	linkPattern = `\[([^\]]*)\]\(([^)]*)\)`
	headerPattern = `^#{1,6}? (.*)`
	httpsPattern = `^https?://`
	hashPattern = `#(.*)`
)

func (p *Parser) Links(content, dirPath string) Links {
	return p.extractLinks(p.parse(content, linkPattern, p.getLink), dirPath)
}

func (p *Parser) Headers(markdown string) []string {
	return p.parse(markdown, headerPattern, p.getHeader)
}

func (*Parser) parse(markdown, pattern string, match match) []string {
	var result []string
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(strings.NewReader(markdown))

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if matches != nil {
			result = append(result, match(matches))
		}
	}
	return result
}

func (*Parser) getLink(matches [][]string) string {
	return strings.Split(matches[0][2], " ")[0]
}

func (*Parser) getHeader(matches [][]string) string {
	return matches[0][1]
}

func (*Parser) extractLinks(links []string, dirPath string) Links {
	var extractedLinks Links
	for _, link := range links {
		if match, _ := regexp.MatchString(httpsPattern, link); match {
			extractedLinks = append(extractedLinks, Link{
				RelPath: "",
				AbsPath: link,
				TypeOf: ExternalLink,
			})
		} else if match, _ := regexp.MatchString(hashPattern, link); match {
			extractedLinks = append(extractedLinks, Link{
				RelPath: link,
				AbsPath: "",
				TypeOf: HashInternalLink,
			})
		} else {
			var absPath string
			if strings.HasPrefix(link, "/") {
				absPath = fmt.Sprintf("%s%s", _BASE_PATH, link)
			} else {
				absPath = path.Join(dirPath, link)
			}
			extractedLinks = append(extractedLinks, Link{
				RelPath: link,
				AbsPath: absPath,
				TypeOf: InternalLink,
			})
		}
	}
	return extractedLinks
}