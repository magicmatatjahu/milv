package main

import (
	"fmt"
	"github.com/magicmatatjahu/milv/milv"
	"github.com/magicmatatjahu/milv/cli"
	"strings"
	"os"
)

func main() {
	cliCommands := milv_cli.ParseCommands()

	config, err := milv.NewConfig(cliCommands.ConfigFile)
	if err != nil {
		panic(err)
	}

	files, _ := milv.NewFiles(cliCommands.Files, config)

	for _, file := range files {
		file.ExtractLinks()
		file.ExtractHeaders()
		file.ValidateLinks()
		file.ExtractStats()
		writeStats(file)
		fmt.Printf("\n")
	}

	failed := writeSummary(files)
	if failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func writeStats(file *milv.File) {
	fmt.Printf("----- %s - status: %v\n", strings.TrimPrefix(file.AbsPath, "/milv/mds/"), file.Status)
	for _, link := range file.Links {
		if link.TypeOf == milv.ExternalLink {
			fmt.Printf("- %s", strings.TrimPrefix(link.AbsPath, "/milv/mds/"))
		} else {
			fmt.Printf("- %s", link.RelPath)
		}
		fmt.Printf(" - status: %v", link.Result.Status)
		if link.Result.Message != "" {
			fmt.Printf(", message: %s", link.Result.Message)
		}
		fmt.Printf("\n")
	}
}

func writeSummary(files milv.Files) bool {
	fmt.Print("\n############### Summary ###############\n")
	failed := false
	for _, file := range files {
		failedLinks := file.Stats.FailedExternalLinks.Count + file.Stats.FailedInternalLinks.Count
		if failedLinks > 0 {
			fmt.Printf("%s\n", strings.TrimPrefix(file.AbsPath, "/milv/mds/"))
			failed = true
		}
		if file.Stats.FailedExternalLinks.Count > 0 {
			fmt.Printf("* External links:\n")
		}
		for _, failedLink := range file.Stats.FailedExternalLinks.Links {
			fmt.Printf("--- %s\t message: %s\n", strings.TrimPrefix(failedLink.AbsPath, "/milv/mds/"), failedLink.Result.Message)
		}
		if file.Stats.FailedInternalLinks.Count > 0 {
			fmt.Printf("* Internal links:\n")
		}
		for _, failedLink := range file.Stats.FailedInternalLinks.Links {
			fmt.Printf("--- %s\t message: %s\n", strings.TrimPrefix(failedLink.RelPath, "/milv/mds/"), failedLink.Result.Message)
		}
		if failedLinks > 0 {
			fmt.Printf("\n")
		}
	}
	if !failed {
		fmt.Print("No issues :-)\n")
	}
	return failed
}