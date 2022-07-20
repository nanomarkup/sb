// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/spf13/cobra"
)

func (v *CmdCoder) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Coder == nil {
			return nil
		}
		if len(args) > 0 {
			return v.Coder.Generate(args[0])
		} else {
			return v.Coder.Generate("")
		}
	}
}
