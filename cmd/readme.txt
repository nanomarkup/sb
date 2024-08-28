package cmd // import "github.com/nanomarkup/sb/cmd"
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
func CmdAddToMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdBuild(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdClean(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdCode(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdCreate(c *sb.SmartCreator) func(cmd *cobra.Command, args []string) error
func CmdDelFromMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdGen(c *sb.SmartGenerator) func(cmd *cobra.Command, args []string) error
func CmdInitMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdList(c *sb.ModHelper) func(cmd *cobra.Command, args []string) error
func CmdManageMod(c *sb.SmartBuilder, f *dl.Formatter) func(cmd *cobra.Command, args []string) error
func CmdRun(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error
func CmdVersion(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string)
func OSStdout() *os.File
TYPES
type AppsPrinter interface {
	Apps() ([]string, error)
}
    AppsPrinter describes methods for displaying all available applications.
type Builder interface {
	Build(string) error
}
    Builder describes methods for building an application using the generated
    items.
type Cleaner interface {
	Clean(string) error
}
    Cleaner describes methods for removing generated/compiled files.
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
type Runner interface {
	Run(string) error
}
    Runner describes methods for running the application.
type SmartBuilder struct {
	cobra.Command
}
    SmartBuilder includes all available commands and handles them.
type VersionPrinter interface {
	Version() string
}
    VersionPrinter describes methods for displaying a version of the
    application.
