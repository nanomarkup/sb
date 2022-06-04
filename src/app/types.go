package app

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
