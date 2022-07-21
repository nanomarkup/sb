package app // import "github.com/sapplications/sbuilder/src/app"


CONSTANTS

const (
	AppName           string = "sb"
	AppVersion        string = "1.0"
	AppVersionString  string = AppName + " version " + AppVersion
	DefaultModuleName string = "main"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	LanguageIsNotSupportedF string = "the current \"%s\" language is not supported\n"
	ApplicationIsMissing    string = "does not found any application in the main"
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

type ModManager interface {
	Init(moduleName, language string) error
	AddItem(moduleName, itemName string) error
	AddDependency(itemName, dependencyName, resolver string, update bool) error
	DeleteItem(itemName string) error
	DeleteDependency(itemName, dependencyName string) error
	ReadAll(language string) (ModReader, error)
	SetLogger(logger Logger)
}

type ModReader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(itemName, dependencyName string) string
	Main() (map[string]string, error)
}

type SmartBuilder struct {
	Lang            func() string
	Builder         interface{}
	ModManager      ModManager
	PluginHandshake plugin.HandshakeConfig
	Logger          Logger
}

func (b *SmartBuilder) AddDependency(item, dependency, resolver string, update bool) error

func (b *SmartBuilder) AddItem(module, item string) error

func (b *SmartBuilder) Build(application string) error

func (b *SmartBuilder) Clean(application string) error

func (b *SmartBuilder) Create(application string) error

func (b *SmartBuilder) DeleteDependency(item, dependency string) error

func (b *SmartBuilder) DeleteItem(item string) error

func (b *SmartBuilder) Generate(application string) error

func (b *SmartBuilder) Init(lang string) error

func (b *SmartBuilder) ReadAll(lang string) (ModReader, error)

func (b *SmartBuilder) Run(application string) error

func (b *SmartBuilder) Version() string

type SmartGenerator struct{}

func (b *SmartGenerator) Generate(application string) error

