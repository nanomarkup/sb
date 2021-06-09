package smodule

type Reader interface {
	Sb() Version
	Lang() Language
	Items() Items
	Dependency(ItemName, DependencyName) ResolverName
	Main() (Item, error)
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
	Item(ItemName, Item) string
	String(module Reader) string
}
