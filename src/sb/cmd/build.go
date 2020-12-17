// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import "github.com/spf13/cobra"

type IBuild interface {
	Build(configuration string) error
}

type BuildCmd struct {
	Build IBuild
	cobra.Command
}

func (v *BuildCmd) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Build == nil {
			return
		}
		if len(args) > 0 {
			v.Build.Build(args[0])
		} else {
			v.Build.Build("")
		}

	}
}
