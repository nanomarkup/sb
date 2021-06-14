package smodule

type Manager interface {
	Init(ModuleName, Language) error
	AddItem(ModuleName, ItemName) error
	AddDependency(ItemName, DependencyName, ResolverName, DoUpdate) error
	DeleteItem(ItemName) error
	DeleteDependency(ItemName, DependencyName) error
	ReadAll(Language) (Reader, error)
}
