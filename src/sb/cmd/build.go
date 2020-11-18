// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build application",
	Long: `Builds an application using the generated items. 

"build" builds an application for the current configuration (rebuild).
"build [configuration]" builds an application for a custom configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var configuration = ""
		// if len(args) > 0 {
		// 	configuration = args[0]
		// }
		// if err := sb.Build(configuration, cli.ConfigFileName); err != nil {
		// 	cli.PrintError(err)
		// }
	},
}
