// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type CmdModDeler struct {
	ModManager
	cobra.Command
}

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
