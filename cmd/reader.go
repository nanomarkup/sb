// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

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
