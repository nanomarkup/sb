// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import "github.com/spf13/cobra"

type IGen interface {
	Generate(configuration string)
}

type GenCmd struct {
	Gen IGen
	cobra.Command
}

func (v *GenCmd) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Gen == nil {
			return
		}
		if len(args) > 0 {
			v.Gen.Generate(args[0])
		} else {
			v.Gen.Generate("")
		}

	}
}
