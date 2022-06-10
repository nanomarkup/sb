// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (v *CmdReader) init() {
	v.SilenceUsage = true
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Reader != nil {
			fmt.Println(v.Reader.Version())
		}
	}
}
