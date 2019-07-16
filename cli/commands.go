package cli

import (
	"regexp"

	"time"
	"github.com/docopt/docopt-go"
	"strings"
	"strconv"
)

const usage = `A bot that parses, checks and validates internal & external URLs in markdown files

Usage:
	milv [-c <config-file>] [-d <document-root>] [-e <external-path>] [-w <white-list>] [-b <black-list>] [-t <timeout>] [-r <repeats>] [--ignore-external] [--ignore-internal] [-v] [<filenames>...]

Options:
	-c, --config-file <config-file>  			Path for config file (eg.: ./foo/bar/milv.yaml)
	-d, --document-root <document-root>  		Set document root directory for absolute paths (helpfully for run in docker)
	-e, --external-path <external-path>  		External path for internal links (helpfully for some cases).
	-w, --white-list <white-list>   			Array of regex of links to exclude from checking (separated by a comma without spaces).
	-b, --black-list <black-list>  				Array of regex of files to exclude from checking (separated by a comma without spaces).
	-t, --timeout <timeout>	 					Set timeout for HTTP requests in seconds. Disabled by default.
	-r, --repeats <repeats>  					Repeats of HTTP requests for failed checks (uint8 type).
	--ignore-external  							Ignore checking external links.
	--ignore-internal  							Ignore checking internal links.
	-v, --verbose  								Be verbose.
`

type Commands struct {
	ConfigFile      string
	DocumentRoot	string
	ExternalPath	string
	WhiteList		[]*regexp.Regexp
	BlackList		[]*regexp.Regexp
	Timeout        	time.Duration
	IgnoreExternal 	bool
	IgnoreInternal 	bool
	Repeats 		uint8
	Verbose        	bool
	FileNames		[]string
}

func parseCommands(argv []string) (Commands, error) {
	commands, err := docopt.ParseArgs(usage, argv, "0.1.0")
	if err != nil {
		return Commands{}, err
	}

	if commands["--config-file"] == nil {
		commands["--config-file"] = ""
	}

	if commands["--document-root"] == nil {
		commands["--document-root"] = ""
	}

	if commands["--external-path"] == nil {
		commands["--external-path"] = ""
	}

	whiteList := []*regexp.Regexp(nil)
	if commands["--white-list"] != nil {
		regexs := strings.Split(commands["--white-list"].(string), ",")
		for _, r := range regexs {
			compiledR, err := regexp.Compile(r)
			if err != nil {
				return Commands{}, err
			}
			whiteList = append(whiteList, compiledR)
		}
	}

	blackList := []*regexp.Regexp(nil)
	if commands["--black-list"] != nil {
		regexs := strings.Split(commands["--black-list"].(string), ",")
		for _, r := range regexs {
			compiledR, err := regexp.Compile(r)
			if err != nil {
				return Commands{}, err
			}
			blackList = append(blackList, compiledR)
		}
	}

	timeout := 0.0
	if commands["--timeout"] != nil {
		temp := commands["--timeout"].(string)
		timeout, err = strconv.ParseFloat(temp, 64)

		if err != nil {
			return Commands{}, err
		}
	}

	if commands["--ignore-external"] == nil {
		commands["--ignore-external"] = false
	}

	if commands["--ignore-internal"] == nil {
		commands["--ignore-internal"] = false
	}

	var repeats uint8 = 0
	if commands["--repeats"] != nil {
		r, err := strconv.ParseUint(commands["--repeats"].(string), 10, 8)
		if err != nil {
			return Commands{}, err
		}
		repeats = uint8(r)
	}

	if commands["--verbose"] == nil {
		commands["--verbose"] = false
	}

	return Commands{
		ConfigFile: commands["--config-file"].(string),
		DocumentRoot: commands["--document-root"].(string),
		ExternalPath: commands["--external-path"].(string),
		WhiteList: whiteList,
		BlackList: blackList,
		Timeout: time.Duration(timeout) * time.Second,
		IgnoreExternal: commands["--ignore-external"].(bool),
		IgnoreInternal: commands["--ignore-internal"].(bool),
		Repeats: repeats,
		Verbose: commands["--verbose"].(bool),
		FileNames: commands["<filenames>"].([]string),
	}, nil
}
