// Package smod manages smart module
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package smodule

import (
	"fmt"
)

var attrs = struct {
	sb      string
	sbFmt   string
	lang    string
	langFmt string
	itemFmt string
	depFmt  string
}{
	"sb",
	"sb %s\n",
	"lang",
	"lang %s\n",
	"%s require (\n",
	"%s %s\n",
}

type Module struct {
	sb    string
	lang  string
	items map[string]map[string]string
}

func (m *Module) Sb() string {
	return m.sb
}

func (m *Module) Lang() string {
	return m.lang
}

func (m *Module) Items() map[string]map[string]string {
	return m.items
}

func (m *Module) Main() (map[string]string, error) {
	main := m.items["main"]
	if main == nil {
		return nil, fmt.Errorf("the main item is not found")
	} else {
		return main, nil
	}
}

func (m *Module) AddItem(item string) error {
	if _, found := m.items[item]; found {
		return fmt.Errorf("\"%s\" item already exists", item)
	}
	m.items[item] = map[string]string{}
	return nil
}

func (m *Module) AddDependency(item, dependency, resolver string, update bool) error {
	curr, found := m.items[item]
	if !found {
		return fmt.Errorf("\"%s\" item does not exist", item)
	}
	if _, found := curr[dependency]; found && !update {
		return fmt.Errorf("\"%s\" already exists for \"%s\" item", dependency, item)
	}
	curr[dependency] = resolver
	return nil
}

func (m *Module) DeleteItem(item string) error {
	delete(m.items, item)
	return nil
}

func (m *Module) DeleteDependency(item, dependency string) error {
	if curr, found := m.items[item]; found {
		delete(curr, dependency)
	}
	return nil
}

func (m *Module) Version() string {
	return fmt.Sprintf(attrs.sbFmt, m.sb)
}

func (m *Module) Language() string {
	return fmt.Sprintf(attrs.langFmt, m.lang)
}

func (m *Module) Dependency(item, dep string) string {
	if deps := m.items[item]; deps != nil {
		if res, found := deps[dep]; found {
			return fmt.Sprintf(attrs.depFmt, dep, res)
		}
	}
	return ""
}
