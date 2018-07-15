package milv_cli

import (
	"flag"
	"strings"
	"os/exec"
	"log"
	"fmt"
)

type Commands struct {
	BasePath	string
	ConfigFile	string
	Files 		[]string
	WhiteList 	[]string
	BlackList	[]string
	AllowRed	bool
	AllowDup	bool
	AllowSSL	bool
	Timeout		int
}

func ParseCommands() Commands {
	basePath := flag.String("basePath", "", "The root source directories used to search for files")
	configFile := flag.String("configFile", "milv.config.yaml", "The root source directories used to search for files")
	whiteList := flag.String("whiteList", "", "The root source directories used to search for files")
	blackList := flag.String("blackList", "", "The root source directories used to search for files")
	allowRed := flag.Bool("allowRed", false, "The root source directories used to search for files")
	allowDup := flag.Bool("allowDup", false, "The root source directories used to search for files")
	allowSSL := flag.Bool("allowSSL", false, "The root source directories used to search for files")
	docker := flag.Bool("docker", false, "The root source directories used to search for files")
	timeout := flag.Int("timeout", 5, "The timeout in seconds used when calling the URL")

	flag.Parse()
	files := flag.Args()

	if *basePath == "" {
		out := runCmd("pwd | tr -d '\n'", true)
		*basePath = string(out)
	}
	if len(files) == 0 {
		out := runCmd("find . -name \"*.md\"", true)
		files = strings.Split(string(out), "\n")
	}

	if *docker {
		*configFile = fmt.Sprintf("%s/%s", "/milv/mds", *configFile)
	}

	return Commands{
		BasePath: *basePath,
		ConfigFile: *configFile,
		Files: files,
		WhiteList: strings.Split(*whiteList, ","),
		BlackList: strings.Split(*blackList, ","),
		AllowRed: *allowRed,
		AllowDup: *allowDup,
		AllowSSL: *allowSSL,
		Timeout: *timeout,
	}
}

func runCmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("/bin/bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}