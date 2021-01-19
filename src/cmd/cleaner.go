// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type Cleaner struct {
	Clean cmd.Cleaner
	cobra.Command
}

func (v *Cleaner) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Clean == nil {
			return
		}
		application := ""
		if len(args) > 0 {
			application = args[0]
		}
		if err := v.Clean.Clean(application); err != nil {
			common.PrintError(err)
		}
	}
}
