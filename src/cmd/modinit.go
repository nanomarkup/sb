// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func (v *CmdModIniter) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.ModManager == nil {
			return nil
		} else if len(args) < 1 {
			return errors.New(LanguageMissing)
		} else {
			defer handleError()
			return v.ModManager.Init(args[0])
		}
	}
}
