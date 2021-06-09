package cmd

import "github.com/sapplications/sbuilder/src/services/smodule"

type Manager interface {
	Init(lang string) error
	AddItem(module, item string) error
	AddDependency(module, item, dependency, resolver string, update bool) error
	DeleteItem(module, item string) error
	DeleteDependency(module, item, dependency string) error
	ReadAll(lang string) (smodule.Reader, error)
}
