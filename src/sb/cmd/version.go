// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/sb/app"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Smart Builder version",
	Long:  `Version prints the current Smart Builder version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.AppVersion)
	},
}
