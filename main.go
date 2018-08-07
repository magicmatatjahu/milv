package main

import (
	"fmt"
	"os"

	"github.com/magicmatatjahu/milv/cli"
	milv "github.com/magicmatatjahu/milv/pkg"
)

func main() {
	cliCommands := cli.ParseCommands()
	milv.SetBasePath(cliCommands.BasePath, false)

	config, err := milv.NewConfig(cliCommands)
	if err != nil {
		panic(err)
	}
	files, _ := milv.NewFiles(cliCommands.Files, config)
	files.Run(cliCommands.Verbose)

	if files.Summary() {
		os.Exit(1)
	}

	fmt.Println("NO ISSUES :-)")
}
