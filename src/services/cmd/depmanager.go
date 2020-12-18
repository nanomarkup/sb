package cmd

type DepManager interface {
	Init(lang string) error
	AddItem(item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
}
