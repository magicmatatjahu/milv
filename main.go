package main

import (
	"github.com/magicmatatjahu/milv/cli"
	milv "github.com/magicmatatjahu/milv/pkg"
	"os"
	"fmt"
)

func main() {
	cliCommands := cli.ParseCommands()
	milv.SetBasePath(cliCommands.BasePath, false)

	config, err := milv.NewConfig(cliCommands)
	if err != nil {
		panic(err)
	}
	files, _ := milv.NewFiles(cliCommands.Files, config)
	files.Run()

	if files.Summary() {
		os.Exit(1)
	} else {
		fmt.Print("NO ISSUES :-)")
		os.Exit(0)
	}
}