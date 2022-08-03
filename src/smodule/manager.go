// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package smodule

import (
	"fmt"
)

func (m *Manager) Init(module string) error {
	return addItem(module, "sb", AppsItemName)
}

func (m *Manager) AddItem(module, item string) error {
	return addItem(module, "sb", item)
}

func (m *Manager) AddDependency(item, dependency, resolver string, update bool) error {
	mod, err := findItem("sb", item)
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
	mod, err := findItem("sb", item)
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
	mod, err := findItem("sb", item)
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

func (m *Manager) ReadAll(kind string) (Reader, error) {
	m.logTrace(fmt.Sprintf("loading \"%s\" modules", kind))
	mods, err := loadModules(kind)
	if err == nil {
		if mods == nil {
			return &module{}, fmt.Errorf("cannot load \"%s\" modules", kind)
		}
		m.logTrace("reading items")
		return loadItems(mods)
	} else {
		return &module{}, err
	}
}

func (m *Manager) SetLogger(logger Logger) {
	m.Logger = logger
}

func (m *Manager) logTrace(message string) {
	if m.Logger != nil {
		m.Logger.Trace(message)
	}
}

func (m *Manager) logDebug(message string) {
	if m.Logger != nil {
		m.Logger.Debug(message)
	}
}

func (m *Manager) logInfo(message string) {
	if m.Logger != nil {
		m.Logger.Info(message)
	}
}

func (m *Manager) logWarn(message string) {
	if m.Logger != nil {
		m.Logger.Warn(message)
	}
}

func (m *Manager) logError(message string) {
	if m.Logger != nil {
		m.Logger.Error(message)
	}
}
