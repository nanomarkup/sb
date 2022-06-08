package smodule

import (
	"fmt"
)

type Reader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(string, string) string
	Main() (map[string]string, error)
}

type modules []Module

type Manager struct {
	Lang func() string
}

func (m *Manager) Init(module, lang string) error {
	return addItem(module, lang, MainItemName)
}

func (m *Manager) AddItem(module, item string) error {
	return addItem(module, m.Lang(), item)
}

func (m *Manager) AddDependency(item, dependency, resolver string, update bool) error {
	mod, err := findItem(m.Lang(), item)
	if err != nil {
		return err
	} else if mod == nil {
		return fmt.Errorf(ItemIsMissingF, item)
	} else {
		if err = mod.AddDependency(item, dependency, resolver, update); err != nil {
			return err
		} else {
			return saveModule(mod)
		}
	}
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
	mod, err := findItem(m.Lang(), item)
	if err != nil {
		return err
	} else if mod != nil {
		if err = mod.DeleteDependency(item, dependency); err != nil {
			return err
		} else {
			return saveModule(mod)
		}
	} else {
		return nil
	}
}

func (m *Manager) ReadAll(lang string) (Reader, error) {
	mods, err := loadModules(lang)
	if err == nil {
		return loadItems(mods)
	} else {
		return nil, err
	}
}
