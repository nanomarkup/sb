// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/spf13/cobra"
)

type Builder struct {
	Builder cmd.Builder
	cobra.Command
}

func (v *Builder) init() {
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Builder == nil {
			return nil
		}
		if len(args) > 0 {
			return v.Builder.Build(args[0])
		} else {
			return v.Builder.Build("")
		}
	}
}
