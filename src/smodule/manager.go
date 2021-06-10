package smodule

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/smodule"
)

type Manager struct {
	Lang func() string
}

func (m *Manager) Init(lang string) error {
	mod, err := loadAll(lang)
	if err == nil {
		if _, err := mod.Main(); err == nil {
			return fmt.Errorf("the main item of %s language already exists", lang)
		} else {
			mod.AddItem("main")
			common.Check(saveModule(ModuleFileName, mod))
			fmt.Print("the main item has been added")
		}
	} else {
		wd, _ := os.Getwd()
		if err.Error() == fmt.Sprintf(ModuleFilesMissingF, wd) {
			mod := &Module{lang, map[string]map[string]string{}}
			mod.AddItem("main")
			common.Check(saveModule(ModuleFileName, mod))
			fmt.Printf(ModuleIsCreatedF, ModuleFileName)
		} else {
			return err
		}
	}
	return nil
}

func (m *Manager) AddItem(module, item string) error {
	mod, err := loadAll(m.Lang())
	common.Check(err)
	common.Check(mod.AddItem(item))
	return saveModule(module, mod)
}

func (m *Manager) AddDependency(module, item, dependency, resolver string, update bool) error {
	mod, err := loadAll(m.Lang())
	common.Check(err)
	common.Check(mod.AddDependency(item, dependency, resolver, update))
	return saveModule(module, mod)
}

func (m *Manager) DeleteItem(module, item string) error {
	mod, err := loadAll(m.Lang())
	common.Check(err)
	common.Check(mod.DeleteItem(item))
	return saveModule(module, mod)
}

func (m *Manager) DeleteDependency(module, item, dependency string) error {
	mod, err := loadAll(m.Lang())
	common.Check(err)
	common.Check(mod.DeleteDependency(item, dependency))
	return saveModule(module, mod)
}

func (m *Manager) ReadAll(lang string) (smodule.Reader, error) {
	return readAll(lang)
}
