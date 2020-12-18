package cmd

import (
	"fmt"

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

func (sb *SmartBuilder) Execute() {
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
	if err := sb.Runner.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	cobra.EnableCommandSorting = false
}
