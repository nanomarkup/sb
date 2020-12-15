// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import "github.com/spf13/cobra"

type IClean interface {
	Clean(configuration string)
}

type CleanCmd struct {
	Clean IClean
	cobra.Command
}

func (v *CleanCmd) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Clean == nil {
			return
		}
		if len(args) > 0 {
			v.Clean.Clean(args[0])
		} else {
			v.Clean.Clean("")
		}

	}
}
