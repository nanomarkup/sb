// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type IVersion interface {
	Version() string
}

type VersionCmd struct {
	Version IVersion
	cobra.Command
}

func (v *VersionCmd) init() {
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Version != nil {
			fmt.Println(v.Version.Version())
		}
	}
}
