// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/sapplications/sbuilder/src/common"
	"github.com/spf13/cobra"
)

type IClean interface {
	Clean(configuration string) error
}

type CleanCmd struct {
	Clean IClean
	cobra.Command
}

func (v *CleanCmd) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Clean == nil {
			return
		}
		configuration := ""
		if len(args) > 0 {
			configuration = args[0]
		}
		if err := v.Clean.Clean(configuration); err != nil {
			common.PrintError(err)
		}
	}
}
