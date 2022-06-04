package smodule

type Reader interface {
	Lang() Language
	Items() map[string]map[string]string
	Dependency(ItemName, DependencyName) ResolverName
	Main() (map[string]string, error)
}

type Writer interface {
	AddItem(ItemName) error
	AddDependency(ItemName, DependencyName, ResolverName, DoUpdate) error
	DeleteItem(ItemName) error
	DeleteDependency(ItemName, DependencyName) error
}

type ReadWriter interface {
	Reader
	Writer
}

type Formatter interface {
	Item(ItemName, map[string]string) string
	String(module Reader) string
}
