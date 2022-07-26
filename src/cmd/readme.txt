package cmd // import "github.com/sapplications/sbuilder/src/cmd"

Package cmd represents Command Line Interface.

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
    Language returns the current language.


TYPES

type Builder interface {
	Build(string) error
}
    Builder describes methods for building an application using the generated
    items.

type Cleaner interface {
	Clean(string) error
}
    Cleaner describes methods for removing generated/compiled files.

type CmdBuilder struct {
	Builder
	cobra.Command
}
    CmdBuilder command builds an application using the generated items.

type CmdCleaner struct {
	Cleaner
	cobra.Command
}
    CmdCleaner command removes generated/compiled files.

type CmdCoder struct {
	Coder
	cobra.Command
}
    CmdCoder command generates code to build the application.

type CmdCreator struct {
	Creator
	cobra.Command
}
    CmdCreator command creates an application by generating smart application
    unit (.sa file).

type CmdGenerator struct {
	Generator
	cobra.Command
}
    CmdGenerator command generates smart builder unit (.sb) using smart
    application unit.

type CmdManager struct {
	ModManager
	ModFormatter
	cobra.Command
}
    CmdManager command manages application items and dependencies.

type CmdModAdder struct {
	ModManager
	cobra.Command
}
    CmdModAdder command adds item or dependency to the exsiting item.

type CmdModDeler struct {
	ModManager
	cobra.Command
}
    CmdModDeler command deletes item or dependency from the exsiting item.

type CmdModIniter struct {
	ModManager
	cobra.Command
}
    CmdModIniter command creates a main.sb module and initialize it with the
    main item. If the main item is exist then do nothing.

type CmdReader struct {
	Reader
	cobra.Command
}
    CmdReader command displays a version of the application.

type CmdRunner struct {
	Runner
	cobra.Command
}
    CmdRunner command runs the application.

type Coder interface {
	Generate(string) error
}
    Coder describes methods for generating code to build the application.

type Creator interface {
	Create(string) error
}
    Creator describes methods for creating an application by generating smart
    application unit (.sa file).

type Generator interface {
	Generate(string) error
}
    Generator describes methods for generating smart builder unit (.sb) using
    smart application unit.

type ModFormatter interface {
	Item(string, map[string]string) string
	String(module ModReader) string
}
    ModFormatter describes methods for formatting module attributes and returns
    it as a string.

type ModManager interface {
	Init(lang string) error
	AddItem(module, item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	ReadAll(lang string) (ModReader, error)
}
    ModManager describes methods for managing application items and
    dependencies.

type ModReader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(string, string) string
	Main() (map[string]string, error)
}
    ModReader describes methods for getting module attributes.

type Reader interface {
	Version() string
}
    Reader describes methods for displaying a version of the application.

type Runner interface {
	Run(string) error
}
    Runner describes methods for running the application.

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
    SmartBuilder includes all available commands and handles them.

func (sb *SmartBuilder) Execute() error
    Execute initializes commands and displays them in the console.

type Starter struct {
	cobra.Command
}
    Starter command is the main command contains all commands to display.

