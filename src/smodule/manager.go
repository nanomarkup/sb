package smodule

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/services/smodule"
)

type Manager struct {
	Lang func() string
}

func (m *Manager) Init(module, lang string) error {
	mod, err := loadAll(lang)
	if err == nil {
		if _, err := mod.Main(); err == nil {
			return fmt.Errorf("the main item of %s language already exists", lang)
		} else {
			mod.AddItem("main")
			if err = saveModule(module, mod); err != nil {
				return err
			}
			fmt.Print("the main item has been added")
		}
	} else {
		wd, _ := os.Getwd()
		if err.Error() == fmt.Sprintf(ModuleFilesMissingF, wd) {
			mod := &Module{lang, map[string]map[string]string{}}
			mod.AddItem("main")
			if err = saveModule(module, mod); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (m *Manager) AddItem(module, item string) error {
	mod, err := loadAll(m.Lang())
	if err != nil {
		wd, _ := os.Getwd()
		if err.Error() == fmt.Sprintf(ModuleFilesMissingF, wd) {
			mod = &Module{m.Lang(), map[string]map[string]string{}}
		} else {
			return err
		}
	}
	if err = mod.AddItem(item); err != nil {
		return err
	} else {
		return saveModule(module, mod)
	}
}

func (m *Manager) AddDependency(module, item, dependency, resolver string, update bool) error {
	mod, err := loadAll(m.Lang())
	if err != nil {
		return err
	} else if err = mod.AddDependency(item, dependency, resolver, update); err != nil {
		return err
	} else {
		return saveModule(module, mod)
	}
}

func (m *Manager) DeleteItem(module, item string) error {
	mod, err := loadAll(m.Lang())
	if err != nil {
		return err
	} else if err = mod.DeleteItem(item); err != nil {
		return err
	} else {
		return saveModule(module, mod)
	}
}

func (m *Manager) DeleteDependency(module, item, dependency string) error {
	mod, err := loadAll(m.Lang())
	if err != nil {
		return err
	} else if err = mod.DeleteDependency(item, dependency); err != nil {
		return err
	} else {
		return saveModule(module, mod)
	}
}

func (m *Manager) ReadAll(lang string) (smodule.Reader, error) {
	return readAll(lang)
}
