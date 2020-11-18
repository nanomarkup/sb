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

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove generated files",
	Long: `Clean removes the generated and built files.

"clean" removes files for the current configuration.
"clean [configuration]"	removes files for a custom configuration.`,
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
			// remove the generated files
			var gen = golang.Generator{
				configuration,
			}
			if err := gen.Clean(); err != nil {
				cli.PrintError(err)
			}
			// remove the built files
			var builder = golang.Builder{
				configuration,
			}
			if err := builder.Clean(); err != nil {
				cli.PrintError(err)
			}
		default:
			cli.PrintError(fmt.Errorf("\"%s\" language is not supported"))
		}
	},
}
