// Copyright 2023 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (v *CmdPrinter) init() {
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Printer != nil {
			apps, err := v.Printer.Apps()
			if err != nil {
				return err
			}
			for _, app := range apps {
				fmt.Println(app)
			}
		}
		return nil
	}
}
