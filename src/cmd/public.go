package cmd

import "github.com/spf13/cobra"

const (
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	SubcmdMissing           string = "subcommand is required"
	ItemMissing             string = "item name is required"
	ModOrDepMissing         string = "module name or dependency name is missing"
	LanguageMissing         string = "language parameter is required"
	ResolverMissing         string = "resolver is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist\n"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist\n"
	UnknownSubcmdF          string = "unknown \"%s\" subcommand\n"
)

type SmartBuilder struct {
	Starter      Starter
	Reader       CmdReader
	Runner       CmdRunner
	Builder      CmdBuilder
	Cleaner      CmdCleaner
	Generator    CmdGenerator
	ModManager   CmdManager
	ModAdder     CmdModAdder
	ModDeler     CmdModDeler
	ModIniter    CmdModIniter
	SilentErrors bool
}

type Starter struct {
	cobra.Command
}

type CmdReader struct {
	Reader
	cobra.Command
}

type Reader interface {
	Version() string
}

type CmdBuilder struct {
	Builder
	cobra.Command
}

type Builder interface {
	Build(string) error
}

type CmdCleaner struct {
	Cleaner
	cobra.Command
}

type Cleaner interface {
	Clean(string) error
}

type CmdRunner struct {
	Runner
	cobra.Command
}

type Runner interface {
	Run(string) error
}

type CmdGenerator struct {
	Generator
	cobra.Command
}

type Generator interface {
	Generate(string) error
}

type CmdManager struct {
	ModManager
	ModFormatter
	cobra.Command
}

type ModManager interface {
	Init(lang string) error
	AddItem(module, item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	ReadAll(lang string) (ModReader, error)
}

type ModReader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(string, string) string
	Main() (map[string]string, error)
}

type ModFormatter interface {
	Item(string, map[string]string) string
	String(module ModReader) string
}

type CmdModAdder struct {
	ModManager
	cobra.Command
}

type CmdModDeler struct {
	ModManager
	cobra.Command
}

type CmdModIniter struct {
	ModManager
	cobra.Command
}

func Language() string {
	return starterFlags.lang
}
