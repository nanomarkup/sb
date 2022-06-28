package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func (sb *SmartBuilder) Execute() error {
	sb.Starter.init()
	sb.Reader.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Generator.init()
	sb.ModManager.init()
	sb.ModAdder.init()
	sb.ModDeler.init()
	sb.ModIniter.init()
	sb.Starter.AddCommand(&sb.ModManager.Command)
	sb.Starter.AddCommand(&sb.Generator.Command)
	sb.Starter.AddCommand(&sb.Builder.Command)
	sb.Starter.AddCommand(&sb.Cleaner.Command)
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
