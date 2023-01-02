// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Execute initializes commands and displays them in the console.
func (sb *SmartBuilder) Execute() error {
	sb.Starter.init()
	sb.Reader.init()
	sb.Runner.init()
	sb.Creator.init()
	sb.Generator.init()
	sb.Coder.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Printer.init()
	sb.ModManager.init()
	sb.ModAdder.init()
	sb.ModDeler.init()
	sb.ModIniter.init()
	sb.Starter.AddCommand(&sb.Creator.Command)
	sb.Starter.AddCommand(&sb.Generator.Command)
	sb.Starter.AddCommand(&sb.Coder.Command)
	sb.Starter.AddCommand(&sb.Builder.Command)
	sb.Starter.AddCommand(&sb.Cleaner.Command)
	sb.Starter.AddCommand(&sb.Runner.Command)
	sb.Starter.AddCommand(&sb.ModManager.Command)
	sb.Starter.AddCommand(&sb.Printer.Command)
	sb.Starter.AddCommand(&sb.Reader.Command)
	sb.ModManager.AddCommand(&sb.ModIniter.Command)
	sb.ModManager.AddCommand(&sb.ModAdder.Command)
	sb.ModManager.AddCommand(&sb.ModDeler.Command)
	err := sb.Starter.Execute()
	if (err != nil) && !sb.SilentErrors {
		os.Exit(1)
	}
	return err
}

func init() {
	cobra.EnableCommandSorting = false
}
