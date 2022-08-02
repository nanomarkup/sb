package app // import "github.com/sapplications/sbuilder/src/app"

Package app implements a Smart Builder application. It is the next
generation of building applications using independent bussiness components.

CONSTANTS

const (
	AppName           string = "sb"
	AppVersion        string = "1.0"
	AppVersionString  string = AppName + " version " + AppVersion
	DefaultModuleName string = "apps"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	LanguageIsNotSupportedF string = "the current \"%s\" language is not supported\n"
	ApplicationIsMissing    string = "does not found any application in the apps"
)

TYPES

type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	IsError() bool
}
    Logger describes methods for logging messages.

type ModManager interface {
	Init(moduleName, language string) error
	AddItem(moduleName, itemName string) error
	AddDependency(itemName, dependencyName, resolver string, update bool) error
	DeleteItem(itemName string) error
	DeleteDependency(itemName, dependencyName string) error
	ReadAll(language string) (ModReader, error)
	SetLogger(logger Logger)
}
    ModManager describes methods for managing a module.

type ModReader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(itemName, dependencyName string) string
	Apps() (map[string]string, error)
}
    ModReader describes methods for getting module attributes.

type SmartBuilder struct {
	Lang            func() string
	Builder         interface{}
	ModManager      ModManager
	PluginHandshake plugin.HandshakeConfig
	Logger          Logger
}
    SmartBuilder manages modules and builds the application.

func (b *SmartBuilder) AddDependency(item, dependency, resolver string, update bool) error
    AddDependency adds a dependency to the item.

func (b *SmartBuilder) AddItem(module, item string) error
    AddItem adds an item to the module.

func (b *SmartBuilder) Build(application string) error
    Build builds an application using the generated items.

func (b *SmartBuilder) Clean(application string) error
    Clean removes generated/compiled files.

func (b *SmartBuilder) Create(application string) error
    Create creates an application by generating smart application unit (.sa
    file).

func (b *SmartBuilder) DeleteDependency(item, dependency string) error
    DeleteDependency deletes the dependency from the item.

func (b *SmartBuilder) DeleteItem(item string) error
    DeleteItem deletes the item from the module.

func (b *SmartBuilder) Generate(application string) error
    Generate generates smart builder unit (.sb) using smart application unit.

func (b *SmartBuilder) Init(lang string) error
    Init creates a apps.sb module and initialize it with the apps item. If the
    apps item is exist then do nothing.

func (b *SmartBuilder) ReadAll(lang string) (ModReader, error)
    ReadAll loads modules.

func (b *SmartBuilder) Run(application string) error
    Run runs the application.

func (b *SmartBuilder) Version() string
    Version displays a version of the application.

type SmartGenerator struct{}
    SmartGenerator generates smart builder unit (.sb) using smart application
    unit.

func (b *SmartGenerator) Generate(application string) error
    Generate generates smart builder unit (.sb) using smart application unit.

