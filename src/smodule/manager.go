package smodule

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/smodule"
)

type Manager struct {
	Lang func() string
}

func (m *Manager) Init(version, lang string) error {
	mod, err := readAll(lang)
	common.Check(err)
	if _, err := mod.Main(); err == nil {
		return fmt.Errorf("the main item of %s language already exists", lang)
	} else {
		mod := &Module{version, lang, map[string]map[string]string{}}
		mod.AddItem("main")
		filename := "main.sb"
		common.Check(saveModule(filename, mod))
		fmt.Printf("%s file has been created", filename)
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
