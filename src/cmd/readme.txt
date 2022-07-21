package cmd // import "github.com/sapplications/sbuilder/src/cmd"

Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2022 Vitalii Noha vitalii.noga@gmail.com


Package cmd represents Command Line Interface

Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com

CONSTANTS

const (
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	AppNameMissing          string = "application name is required"
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

FUNCTIONS

func Language() string

TYPES

type Builder interface {
	Build(string) error
}

type Cleaner interface {
	Clean(string) error
}

type CmdBuilder struct {
	Builder
	cobra.Command
}

type CmdCleaner struct {
	Cleaner
	cobra.Command
}

type CmdCoder struct {
	Coder
	cobra.Command
}

type CmdCreator struct {
	Creator
	cobra.Command
}

type CmdGenerator struct {
	Generator
	cobra.Command
}

type CmdManager struct {
	ModManager
	ModFormatter
	cobra.Command
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

type CmdReader struct {
	Reader
	cobra.Command
}

type CmdRunner struct {
	Runner
	cobra.Command
}

type Coder interface {
	Generate(string) error
}

type Creator interface {
	Create(string) error
}

type Generator interface {
	Generate(string) error
}

type ModFormatter interface {
	Item(string, map[string]string) string
	String(module ModReader) string
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

type Reader interface {
	Version() string
}

type Runner interface {
	Run(string) error
}

type SmartBuilder struct {
	Starter      Starter
	Reader       CmdReader
	Runner       CmdRunner
	Creator      CmdCreator
	Generator    CmdGenerator
	Coder        CmdCoder
	Builder      CmdBuilder
	Cleaner      CmdCleaner
	ModManager   CmdManager
	ModAdder     CmdModAdder
	ModDeler     CmdModDeler
	ModIniter    CmdModIniter
	SilentErrors bool
}

func (sb *SmartBuilder) Execute() error

type Starter struct {
	cobra.Command
}

