// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"errors"

	src "github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type ModAdder struct {
	Manager src.Manager
	cobra.Command
}

func (v *ModAdder) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		defer handleError()
		if v.Manager == nil {
			return nil
		} else if len(args) < 1 {
			return errors.New(ItemMissing)
		} else if len(args) < 2 {
			return errors.New(ModOrDepMissing)
		} else if len(args) == 2 {
			return v.Manager.AddItem(args[1], args[0])
		} else if len(args) > 2 {
			return v.Manager.AddDependency(args[0], args[1], args[2], false)
		}
		return nil
	}
}
