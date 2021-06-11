package cmd

import (
	"github.com/spf13/cobra"
)

type SmartBuilder struct {
	Runner    Runner
	Reader    Reader
	Builder   Builder
	Cleaner   Cleaner
	Generator Generator
	Manager   Manager
	ModInit   ModInit
}

func (sb *SmartBuilder) Execute() error {
	sb.Runner.init()
	sb.Reader.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Generator.init()
	sb.Manager.init()
	sb.ModInit.init()
	sb.Runner.AddCommand(&sb.Manager.Command)
	sb.Runner.AddCommand(&sb.Generator.Command)
	sb.Runner.AddCommand(&sb.Builder.Command)
	sb.Runner.AddCommand(&sb.Cleaner.Command)
	sb.Runner.AddCommand(&sb.Reader.Command)
	sb.Manager.AddCommand(&sb.ModInit.Command)
	return sb.Runner.Execute()
}

func init() {
	cobra.EnableCommandSorting = false
}
