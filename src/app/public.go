package app

type SmartBuilder struct {
	Lang        func() string
	ModManager  Manager
	GoBuilder   Builder
	GoGenerator Generator
}

type Builder interface {
	Init(items map[string]map[string]string)
	Build(appName string) error
	Clean(appName string) error
}

type Generator interface {
	Init(items map[string]map[string]string)
	Clean(appName string) error
	Generate(appName string) error
}

type Manager interface {
	Init(moduleName, language string) error
	AddItem(moduleName, itemName string) error
	AddDependency(itemName, dependencyName, resolver string, update bool) error
	DeleteItem(itemName string) error
	DeleteDependency(itemName, dependencyName string) error
	ReadAll(language string) (Reader, error)
}

type Reader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(itemName, dependencyName string) string
	Main() (map[string]string, error)
}

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
