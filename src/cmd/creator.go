// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (v *CmdCreator) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Creator == nil {
			return nil
		}
		if len(args) > 0 {
			return v.Creator.Create(args[0])
		} else {
			return fmt.Errorf(AppNameMissing)
		}
	}
}
