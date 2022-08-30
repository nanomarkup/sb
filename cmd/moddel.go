// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func (v *CmdModDeler) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		defer handleError()
		if v.ModManager == nil {
			return nil
		} else if len(args) < 1 {
			return errors.New(ItemMissing)
		} else if len(args) == 1 {
			return v.ModManager.DeleteItem(args[0])
		} else if len(args) == 2 {
			return v.ModManager.DeleteDependency(args[1], args[0])
		} else {
			return nil
		}
	}
}
