// Copyright © 2018 Salesforce
// +build linux darwin freebsd

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/organization/repo-name/cmd"
)

var (
	// CompileDate tracks when the binary was compiled.
	CompileDate = "No date provided."

	// GitCommit tracks the SHA of the built binary.
	GitCommit = "No revision provided."

	// Version is the version of the built binary.
	Version = "No version provided."

	// GoVersion details the version of Go this was compiled with.
	GoVersion = runtime.Version()
)

func main() {
	// -v or --version
	ShowVersionInfo()

	// Let's start it up.
	cmd.RootCmd.Execute()
}

// ShowVersionInfo shows information about the compiled binary.
func ShowVersionInfo() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			fmt.Printf("%-8s : %s\n%-8s : %s\n%-8s : %s\n%-8s : %s\n", "Version", Version, "Revision", GitCommit, "Date", CompileDate, "Go", GoVersion)
			os.Exit(0)
		}
	}
}
