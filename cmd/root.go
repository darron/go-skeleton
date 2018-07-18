// Copyright © 2018 Salesforce
// +build linux darwin freebsd

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// JSON shows data in JSON format.
	JSON bool
)

// RootCmd is what runs by default.
var RootCmd = &cobra.Command{
	Use:  "binary-name",
	Long: `Some description goes here.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("`binary-name -h` for help information.")
		fmt.Println("`binary-name -v` for version information.")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&JSON, "json", "j", false, "output in json")
}
