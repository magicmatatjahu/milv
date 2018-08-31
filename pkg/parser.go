package pkg

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"path"
	"regexp"
	"strings"
)

type Parser struct{}

type match func([][]string) string

const (
	linkPattern   = `\[([^\]]*)\]\(([^)]*)\)|\bhttps?://\S*\b`
	headerPattern = `^#{1,6}? (.*)`
	httpsPattern  = `^https?://`
	hashPattern   = `#(.*)`
)

func (p *Parser) Links(markdown, dirPath string) Links {
	return p.extractLinks(p.parse(markdown, linkPattern, p.getLink), dirPath)
}

func (p *Parser) Headers(markdown string) []string {
	return p.parse(markdown, headerPattern, p.getHeader)
}

func (p *Parser) Anchors(body io.ReadCloser) (ids []string) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			id := p.getId(z.Token())
			if id != "" {
				// github always add "user-content-" prefix to anchor in .md files
				id = p.removePrefixFomAnchor(id)
				ids = append(ids, id)
			}
		}
	}
	return
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
	substring := strings.Split(matches[0][2], " ")[0]
	if substring == "" {
		return matches[0][0]
	}
	return substring
}

func (p *Parser) getHeader(matches [][]string) string {
	header := matches[0][1]

	re := regexp.MustCompile(`]\(([^)]*)\)`)
	header = re.ReplaceAllString(header, "")

	reg, err := regexp.Compile("[^a-zA-Z0-9 -]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(header, "")
}

func (p *Parser) extractLinks(links []string, dirPath string) Links {
	var extractedLinks Links
	for _, link := range links {
		if match, _ := regexp.MatchString(httpsPattern, link); match {
			extractedLinks = append(extractedLinks, p.externalLink(link))
		} else if match, _ := regexp.MatchString(hashPattern, link); match {
			extractedLinks = append(extractedLinks, p.hashInternalLink(link))
		} else {
			extractedLinks = append(extractedLinks, p.internalLink(link, dirPath))
		}
	}
	return extractedLinks
}

func (p *Parser) externalLink(link string) Link {
	return Link{
		RelPath: "",
		AbsPath: link,
		TypeOf:  ExternalLink,
	}
}

func (p *Parser) hashInternalLink(link string) Link {
	return Link{
		RelPath: link,
		AbsPath: "",
		TypeOf:  HashInternalLink,
	}
}

func (p *Parser) internalLink(link, dirPath string) Link {
	var absPath string

	if strings.HasPrefix(link, "/") {
		absPath = fmt.Sprintf("%s%s", _BASE_PATH, link)
	} else {
		absPath = path.Join(dirPath, link)
	}

	return Link{
		RelPath: link,
		AbsPath: absPath,
		TypeOf:  InternalLink,
	}
}

func (*Parser) getId(t html.Token) string {
	for _, attr := range t.Attr {
		if attr.Key == "id" || (t.Data == "a" && attr.Key == "name") {
			return attr.Val
		}
	}
	return ""
}

func (*Parser) removePrefixFomAnchor(anchor string) string {
	prefixes := []string{
		//Github .md files
		"user-content-",
		//Stackoverflow
		"comments-",
		"comment-",
		"link-post-",
		"comments-link-",
		"answer-",
	}

	for _, prefix := range prefixes {
		anchor = strings.TrimPrefix(anchor, prefix)
	}
	return anchor
}