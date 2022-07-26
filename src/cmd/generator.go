// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"
)

func (v *CmdGenerator) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Generator == nil {
			return nil
		}
		if len(args) > 0 {
			return v.Generator.Generate(args[0])
		} else {
			return v.Generator.Generate("")
		}
	}
}
