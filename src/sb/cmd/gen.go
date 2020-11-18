// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate configuration",
	Long: `Generates all items for the selected configuration. 

"gen" generates all items for the current configuration (update).
"gen [configuration]" generates all items for a custom configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var configuration = ""
		// if len(args) > 0 {
		// 	configuration = args[0]
		// }
		// if err := sb.Generate(configuration, cli.ConfigFileName); err != nil {
		// 	cli.PrintError(err)
		// }
	},
}
