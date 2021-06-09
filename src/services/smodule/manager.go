package smodule

type Manager interface {
	Init(Language) error
	AddItem(ModuleName, ItemName) error
	AddDependency(ModuleName, ItemName, DependencyName, ResolverName, DoUpdate) error
	DeleteItem(ModuleName, ItemName) error
	DeleteDependency(ModuleName, ItemName, DependencyName) error
	ReadAll(Language) (Reader, error)
}
