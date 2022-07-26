// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package smodule

import (
	"fmt"
)

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
	m.Logger.Trace(fmt.Sprintf("loading modules using \"%s\" language", lang))
	mods, err := loadModules(lang)
	if err == nil {
		if mods == nil {
			return &module{}, fmt.Errorf("cannot load modules using \"%s\" language", lang)
		}
		m.Logger.Trace("reading items")
		return loadItems(mods)
	} else {
		return &module{}, err
	}
}

func (m *Manager) SetLogger(logger Logger) {
	m.Logger = logger
}
