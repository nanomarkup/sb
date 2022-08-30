// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func (v *CmdModAdder) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		defer handleError()
		if v.ModManager == nil {
			return nil
		} else if len(args) < 1 {
			return errors.New(ItemMissing)
		} else if len(args) < 2 {
			return errors.New(ModOrDepMissing)
		} else if len(args) == 2 {
			return v.ModManager.AddItem(args[1], args[0])
		} else if len(args) > 2 {
			return v.ModManager.AddDependency(args[0], args[1], args[2], false)
		} else {
			return nil
		}
	}
}
