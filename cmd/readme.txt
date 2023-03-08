package cmd // import "github.com/sapplications/sb/cmd"

Package cmd represents Command Line Interface.

CONSTANTS

const (
	// application
	AppsItemName      string = "apps"
	DefaultModuleName string = "apps"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	AppNameMissing          string = "application name is required"
	SubcmdMissing           string = "subcommand is required"
	ItemMissing             string = "item name is required"
	ModOrDepMissing         string = "module name or dependency name is missing"
	ResolverMissing         string = "resolver is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist"
	UnknownSubcmdF          string = "unknown \"%s\" subcommand"
)

FUNCTIONS

func CmdCode(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdCreate(c *sb.SmartCreator) func(cmd *cobra.Command, args []string) error

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
    CmdModIniter command creates a apps.sb module and initialize it with the
    apps item. If the apps item is exist then do nothing.

type CmdPrinter struct {
	Printer
	cobra.Command
}
    CmdPrinter command displays all available applications.

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
	Item(string, [][]string) string
	String(module ModReader) string
}
    ModFormatter describes methods for formatting module attributes and returns
    it as a string.

type ModManager interface {
	AddItem(module, item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	ReadAll(kind string) (ModReader, error)
}
    ModManager describes methods for managing application items and
    dependencies.

type ModReader interface {
	Items() map[string][][]string
	Dependency(string, string) string
}
    ModReader describes methods for getting module attributes.

type Printer interface {
	Apps() ([]string, error)
}
    Printer describes methods for displaying all available applications.

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
	Generator    CmdGenerator
	Builder      CmdBuilder
	Cleaner      CmdCleaner
	Printer      CmdPrinter
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
    Starter command is the apps command contains all commands to display.

