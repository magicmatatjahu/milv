package pkg

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
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

	client := &http.Client{}
	url, err := url.ParseRequestURI(link.AbsPath)
	if err != nil {
		return link, err
	}

	resp, err := client.Get(url.String())
	if err != nil {
		return link, err
	}

	if match, _ := regexp.MatchString(`^2*`, strconv.Itoa(resp.StatusCode)); match {
		link.Result.Status = true
	} else {
		link.Result.Status = false
		link.Result.Message = resp.Status
	}
	resp.Body.Close()
	return link, nil
}

func (*Validation) internalLink(link Link) (Link, error) {
	if link.TypeOf != InternalLink {
		return link, nil
	}

	if err := fileExists(link.AbsPath); err == nil {
		link.Result.Status = true
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
		link.Result.Status = false
		link.Result.Message = "The specified header doesn't exist"
	}
	return link, nil
}
