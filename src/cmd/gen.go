// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type Generator struct {
	Gen cmd.Generator
	cobra.Command
}

func (v *Generator) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Gen == nil {
			return
		}
		configuration := ""
		if len(args) > 0 {
			configuration = args[0]
		}
		if err := v.Gen.Generate(configuration); err != nil {
			common.PrintError(err)
		}
	}
}
