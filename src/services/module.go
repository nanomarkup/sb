package services

type IModule interface {
	Sb() string
	Lang() string
	Items() map[string]map[string]string
	Init(version, language string)
	Main() (map[string]string, error)
	AddItem(item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	LoadFromFile(filePath string) error
	SaveToFile(filePath string) error
}
