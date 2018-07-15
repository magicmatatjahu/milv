package milv

import (
	"regexp"
	"bufio"
	"strings"
	"io/ioutil"
	"net/http"
	"bytes"
	"path"
)

type Parser struct {}

const (
	httpsPattern = "^https://*|^http://*"
	hashPattern = `#(.*)`
)

func (p *Parser) GetLinks(str string) ([]string, error) {
	if match, _ := regexp.MatchString(httpsPattern, str); match {
		return p.GetLinksFromMarkdownUrl(str)
	} else {
		return p.GetLinksFromMarkdownFile(str)
	}
}

func (p *Parser) GetAndExtractLinks(dirPath string, filePath string) ([]Link, error) {
	var links []string
	var err error

	if match, _ := regexp.MatchString(httpsPattern, filePath); match {
		links, err = p.GetLinksFromMarkdownUrl(filePath)
	} else {
		links, err = p.GetLinksFromMarkdownFile(filePath)
	}

	if err != nil {
		return []Link{}, err
	} else {
		return p.ExtractLinks(dirPath, links), nil
	}
}

func (p *Parser) GetLinksFromMarkdownFile(fileName string) ([]string, error) {
	markdown, err := readMarkdownFile(fileName)
	if err != nil {
		return []string{}, err
	}
	//markdown = removeCodeBlocks(markdown)
	return getAllLinks(markdown), nil
}

func (p *Parser) GetLinksFromMarkdownUrl(url string) ([]string, error) {
	markdown, err := downloadMarkdownUrl(url)
	if err != nil {
		return []string{}, err
	}
	//markdown = removeCodeBlocks(markdown)
	return getAllLinks(markdown), nil
}

func (p *Parser) GetAllHeaders(fileName string) ([]string, error) {
	markdown, err := readMarkdownFile(fileName)
	if err != nil {
		return []string{}, err
	}
	return getAllHeaders(markdown), nil
}

func (p *Parser) ExtractLinks(dirPath string, links []string) []Link {
	var extractedLinks []Link
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
			extractedLinks = append(extractedLinks, Link{
				RelPath: link,
				AbsPath: path.Join(dirPath, link),
				TypeOf: InternalLink,
			})
		}
	}
	return extractedLinks
}

func readMarkdownFile(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func downloadMarkdownUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String(), nil
}

func getAllLinks(markdown string) []string {
	var links []string
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if matches != nil {
			links = append(links, strings.Split(matches[0][2], " ")[0])
		}
	}
	return links
}

func getAllHeaders(markdown string) []string {
	var headers []string
	re := regexp.MustCompile(`^#{1,6}? (.*)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if matches != nil {
			headers = append(headers, matches[0][1])
		}
	}
	return headers
}

func removeCodeBlocks(markdown string) string {
	//TODO improve that shit
	re := regexp.MustCompile(`\r?\n`)
	markdown = re.ReplaceAllString(markdown, " ")

	re = regexp.MustCompile(`(\x60{3}?.+?\x60{3}?)`)
	return re.ReplaceAllString(markdown, "")
}