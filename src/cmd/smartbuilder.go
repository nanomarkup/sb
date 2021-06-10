package cmd

import (
	"github.com/spf13/cobra"
)

type SmartBuilder struct {
	Runner     Runner
	Reader     Reader
	Builder    Builder
	Cleaner    Cleaner
	Generator  Generator
	DepManager DepManager
}

func (sb *SmartBuilder) Execute() error {
	sb.Runner.init()
	sb.Reader.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Generator.init()
	sb.DepManager.init()
	sb.Runner.AddCommand(&sb.DepManager.Command)
	sb.Runner.AddCommand(&sb.Generator.Command)
	sb.Runner.AddCommand(&sb.Builder.Command)
	sb.Runner.AddCommand(&sb.Cleaner.Command)
	sb.Runner.AddCommand(&sb.Reader.Command)
	sb.Runner.SilenceUsage = true
	return sb.Runner.Execute()
}

func init() {
	cobra.EnableCommandSorting = false
}
