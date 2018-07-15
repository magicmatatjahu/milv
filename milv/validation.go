package milv

import (
	"net/http"
	"regexp"
	"strconv"
	"os"
	"fmt"
	"strings"
)

type Validation struct{}

func (v *Validation) ValidateLinks(links []Link, headers []string) ([]Link, error) {
	var validatedLinks []Link
	for _, link := range links {
		if link.TypeOf == ExternalLink {
			link, _ = validateExternalLink(link)
			validatedLinks = append(validatedLinks, link)
		} else if link.TypeOf == InternalLink {
			link, _ = validateInternalLink(link)
			validatedLinks = append(validatedLinks, link)
		} else {
			link, _ = validateHashInternalLink(link, headers)
			validatedLinks = append(validatedLinks, link)
		}
	}
	return validatedLinks, nil
}

func (v *Validation) ValidateExternalLinks(links []Link) ([]Link, error) {
	for _, link := range links {
		link, _ = validateExternalLink(link)
	}
	return links, nil
}

func (v *Validation) ValidateInternalLinks(links []Link) ([]Link, error) {
	for _, link := range links {
		link, _ = validateInternalLink(link)
	}
	return links, nil
}

func (v *Validation) ValidateHashInternalLinks(links []Link, headers []string) ([]Link, error) {
	for _, link := range links {
		link, _ = validateHashInternalLink(link, headers)
	}
	return links, nil
}

func validateExternalLink(link Link) (Link, error) {
	if link.TypeOf != ExternalLink {
		return link, nil
	}

	resp, err := http.Get(link.AbsPath)
	if err != nil {
		return link, err
	}

	if match, _ := regexp.MatchString("^2*", strconv.Itoa(resp.StatusCode)); match {
		link.Result.Status = true
		link.Result.Message = "Not found issue"
	} else {
		link.Result.Status = false
		link.Result.Message = fmt.Sprintf("Status code: %s", strconv.Itoa(resp.StatusCode))
	}
	return link, nil
}

func validateInternalLink(link Link) (Link, error) {
	if link.TypeOf != InternalLink {
		return link, nil
	}

	//TODO change that
	if strings.HasPrefix(link.RelPath, "/") {
		link.AbsPath = fmt.Sprintf("/milv/mds%s", link.RelPath)
	}

	if err := fileExists(link.AbsPath); err == nil {
		link.Result.Status = true
		link.Result.Message = "Not found issue"
	} else {
		link.Result.Status = false
		link.Result.Message = "The specified file doesn't exist"
	}
	return link, nil
}

func validateHashInternalLink(link Link, headers []string) (Link, error) {
	if link.TypeOf != HashInternalLink {
		return link, nil
	}

	if match := headerExists(link.RelPath, headers); match {
		link.Result.Status = true
		link.Result.Message = "Not found issue"
	} else {
		link.Result.Status = false
		link.Result.Message = "The specified header doesn't exist"
	}
	return link, nil
}

func fileExists(file string) error {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func headerExists(link string, headers []string) bool {
	link = strings.TrimPrefix(link, "#")
	for _, header := range headers {
		if link == strings.ToLower(strings.Replace(header, " ", "-", -1)) {
			return true
		}
	}
	return false
}