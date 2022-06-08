package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type SmartBuilder struct {
	Runner       Runner
	Reader       CmdReader
	Builder      CmdBuilder
	Cleaner      CmdCleaner
	Generator    CmdGenerator
	ModManager   CmdManager
	ModAdder     CmdModAdder
	ModDeler     CmdModDeler
	ModIniter    CmdModIniter
	SilentErrors bool
}

func (sb *SmartBuilder) Execute() error {
	sb.Runner.init()
	sb.Reader.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Generator.init()
	sb.ModManager.init()
	sb.ModAdder.init()
	sb.ModDeler.init()
	sb.ModIniter.init()
	sb.Runner.AddCommand(&sb.ModManager.Command)
	sb.Runner.AddCommand(&sb.Generator.Command)
	sb.Runner.AddCommand(&sb.Builder.Command)
	sb.Runner.AddCommand(&sb.Cleaner.Command)
	sb.Runner.AddCommand(&sb.Reader.Command)
	sb.ModManager.AddCommand(&sb.ModIniter.Command)
	sb.ModManager.AddCommand(&sb.ModAdder.Command)
	sb.ModManager.AddCommand(&sb.ModDeler.Command)
	err := sb.Runner.Execute()
	if (err != nil) && !sb.SilentErrors {
		os.Exit(1)
	}
	return err
}

func init() {
	cobra.EnableCommandSorting = false
}
