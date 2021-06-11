// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"errors"

	"github.com/sapplications/sbuilder/src/common"
	src "github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type ModInit struct {
	Manager src.Manager
	cobra.Command
}

func (v *ModInit) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Manager == nil {
			return nil
		} else if len(args) < 1 {
			return errors.New(LanguageMissing)
		} else {
			defer common.Recover()
			return v.Manager.Init(args[0])
		}
	}
}
