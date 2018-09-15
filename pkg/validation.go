package pkg

import (
	"fmt"
	"github.com/schollz/closestmatch"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Validation struct {
}

func (v *Validation) Links(links []Link, optionalHeaders ...Headers) []Link {
	if len(links) == 0 {
		return []Link{}
	}

	var headers Headers
	var headersExist bool
	if len(optionalHeaders) == 1 {
		headers = optionalHeaders[0]
		headersExist = len(headers) > 0
	}

	var validatedLinks []Link
	for _, link := range links {
		if link.TypeOf == ExternalLink {
			link, _ = v.externalLink(link)
			validatedLinks = append(validatedLinks, link)
		} else if link.TypeOf == InternalLink {
			link, _ = v.internalLink(link)
			validatedLinks = append(validatedLinks, link)
		} else {
			if headersExist {
				link, _ = v.hashInternalLink(link, headers)
				validatedLinks = append(validatedLinks, link)
			}
		}
	}
	return validatedLinks
}

func (v *Validation) ExternalLinks(links Links) (Links, error) {
	for _, link := range links {
		if link.TypeOf == ExternalLink {
			link, _ = v.externalLink(link)
		}
	}
	return links, nil
}

func (v *Validation) InternalLink(links Links) (Links, error) {
	for _, link := range links {
		if link.TypeOf == InternalLink {
			link, _ = v.externalLink(link)
		}
	}
	return links, nil
}

func (v *Validation) HashInternalLinks(links Links, headers Headers) (Links, error) {
	for _, link := range links {
		if link.TypeOf == HashInternalLink {
			link, _ = v.hashInternalLink(link, headers)
		}
	}
	return links, nil
}

func (*Validation) externalLink(link Link) (Link, error) {
	if link.TypeOf != ExternalLink {
		return link, nil
	}

	var status bool
	message := ""

	url, err := url.Parse(link.AbsPath)
	if err != nil {
		link.Result.Status = false
		link.Result.Message = err.Error()
		return link, err
	}
	absPath := fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, url.Path)

	client := &http.Client{}
	if link.Config != nil && link.Config.Timeout != nil && *link.Config.Timeout != 0 {
		client.Timeout = time.Duration(int(time.Second) * (*link.Config.Timeout))
	} else {
		client.Timeout = time.Duration(int(time.Second) * 30)
	}

	requestRepeats := int8(1)
	if link.Config != nil && link.Config.ReguestRepeats != nil && *link.Config.ReguestRepeats > 0 {
		requestRepeats = *link.Config.ReguestRepeats
	}

	for i := int8(0); i < requestRepeats; i++ {
		resp, err := client.Get(absPath)
		if err != nil {
			status = false
			message = err.Error()
			continue
		}

		allowRedirect := false
		if link.Config != nil && link.Config.AllowRedirect != nil {
			allowRedirect = *link.Config.AllowRedirect
		}

		statusCode, regexpPattern := strconv.Itoa(resp.StatusCode), `^2[0-9][0-9]`
		if allowRedirect {
			regexpPattern = `^2[0-9][0-9]|^3[0-9][0-9]`
		}

		if match, _ := regexp.MatchString(regexpPattern, statusCode); match && resp != nil {
			status = true

			if !allowRedirect && url.Fragment != "" {
				match, _ = regexp.MatchString(`[a-zA-Z]`, string(url.Fragment[0]))
				if !match {
					break
				}

				parser := &Parser{}
				anchors := parser.Anchors(resp.Body)

				if contains(anchors, url.Fragment) {
					status = true
				} else {
					cm := closestmatch.New(anchors, []int{4, 3, 5})
					closestAnchor := cm.Closest(url.Fragment)

					status = false
					if closestAnchor != "" {
						message = fmt.Sprintf("The specified anchor doesn't exist in website. Did you mean about #%s?", closestAnchor)
					} else {
						message = "The specified anchor doesn't exist"
					}
				}
			}

			resp.Body.Close()
			break
		} else {
			status = false
			message = resp.Status
			resp.Body.Close()
		}
	}

	link.Result.Status = status
	link.Result.Message = message
	return link, nil
}

func (v *Validation) internalLink(link Link) (Link, error) {
	if link.TypeOf != InternalLink {
		return link, nil
	}

	splitted := strings.Split(link.AbsPath, "#")

	if err := fileExists(splitted[0]); err == nil {
		link.Result.Status = true

		if len(splitted) == 2 {
			if !v.isHashInFile(splitted[0], splitted[1]) {
				link.Result.Status = false
				link.Result.Message = "The specified header doesn't exist in file"
			}
		}
	} else {
		link.Result.Status = false
		link.Result.Message = "The specified file doesn't exist"
	}
	return link, nil
}

func (*Validation) hashInternalLink(link Link, headers Headers) (Link, error) {
	if link.TypeOf != HashInternalLink {
		return link, nil
	}

	if match := headerExists(link.RelPath, headers); match {
		link.Result.Status = true
	} else {
		cm := closestmatch.New(headers, []int{4, 3, 5})
		closestHeader := cm.Closest(link.RelPath)
		closestHeader = strings.Replace(closestHeader, " ", "-", -1)
		closestHeader = strings.ToLower(closestHeader)

		link.Result.Status = false
		link.Result.Message = "The specified header doesn't exist in file"
		//if closestHeader != "" {
		//	link.Result.Message = fmt.Sprintf("The specified header doesn't exist in file. Did you mean about #%s?", closestHeader)
		//} else {
		//	link.Result.Message = "The specified header doesn't exist in file"
		//}
	}
	return link, nil
}

func (*Validation) isHashInFile(file, header string) bool {
	markdown, err := readMarkdown(file)
	if err != nil {
		return false
	}

	parser := Parser{}
	return headerExists(header, parser.Headers(markdown))
}
