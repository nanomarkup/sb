// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/golang"
	"github.com/sapplications/sbuilder/src/sb/app"
	"github.com/sapplications/sbuilder/src/smod"
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
		var configuration = ""
		if len(args) > 0 {
			configuration = args[0]
		}
		defer cli.Recover()
		// check configuration
		var c smod.ConfigFile
		cli.Check(c.LoadFromFile(app.ModFileName))
		if err := app.CheckConfiguration(configuration, &c); err != nil {
			cli.PrintError(err)
			return
		}
		// process configuration
		switch c.Lang {
		case app.Langs.Go:
			var gen = golang.Generator{
				configuration,
			}
			if err := gen.Generate(&c); err != nil {
				cli.PrintError(err)
			}
		default:
			cli.PrintError(fmt.Errorf("\"%s\" language is not supported"))
		}
	},
}
