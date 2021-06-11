// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type Cleaner struct {
	Cleaner cmd.Cleaner
	cobra.Command
}

func (v *Cleaner) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Cleaner == nil {
			return nil
		}
		if len(args) > 0 {
			return v.Cleaner.Clean(args[0])
		} else {
			return v.Cleaner.Clean("")
		}
	}
}
