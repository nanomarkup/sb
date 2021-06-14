package smodule

import (
	"github.com/sapplications/sbuilder/src/services/smodule"
)

type modules []Module

type Manager struct {
	Lang func() string
}

func (m *Manager) Init(module, lang string) error {
	return addItem(module, lang, mainItemName)
}

func (m *Manager) AddItem(module, item string) error {
	return addItem(module, m.Lang(), item)
}

func (m *Manager) AddDependency(item, dependency, resolver string, update bool) error {
	// mod, err := m.loadItems()
	// if err != nil {
	// 	return err
	// } else if err = mod.AddDependency(item, dependency, resolver, update); err != nil {
	// 	return err
	// } else {
	// 	return nil
	// 	// return saveModule(module, mod)
	// }
	return nil
}

func (m *Manager) DeleteItem(item string) error {
	mod, err := findItem(m.Lang(), item)
	if err != nil {
		return err
	} else if mod != nil {
		if err = mod.DeleteItem(item); err != nil {
			return err
		} else {
			return saveModule(mod)
		}
	} else {
		return nil
	}
}

func (m *Manager) DeleteDependency(item, dependency string) error {
	// mod, err := m.loadItems()
	// if err != nil {
	// 	return err
	// } else if err = mod.DeleteDependency(item, dependency); err != nil {
	// 	return err
	// } else {
	// 	return saveModule(module, mod)
	// }
	return nil
}

func (m *Manager) ReadAll(lang string) (smodule.Reader, error) {
	mods, err := loadModules(lang)
	if err == nil {
		return loadItems(mods)
	} else {
		return nil, err
	}
}
