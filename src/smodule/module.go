// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package smodule

import (
	"fmt"
)

var attrs = struct {
	use     string
	useFmt  string
	itemFmt string
	depFmt  string
}{
	"use",
	"use %s\n",
	"%s require (\n",
	"%s %s\n",
}

func (m *module) Lang() string {
	return m.lang
}

func (m *module) Items() Items {
	return m.items
}

func (m *module) Main() (Item, error) {
	main := m.items["main"]
	if main == nil {
		return nil, fmt.Errorf("the main item is not found")
	} else {
		return main, nil
	}
}

func (m *module) AddItem(item string) error {
	if _, found := m.items[item]; found {
		return fmt.Errorf("\"%s\" item already exists", item)
	}
	m.items[item] = Item{}
	return nil
}

func (m *module) AddDependency(item, dependency, resolver string, update bool) error {
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

func (m *module) DeleteItem(item string) error {
	delete(m.items, item)
	return nil
}

func (m *module) DeleteDependency(item, dependency string) error {
	if curr, found := m.items[item]; found {
		delete(curr, dependency)
	}
	return nil
}

func (m *module) Language() string {
	return fmt.Sprintf(attrs.useFmt, m.lang)
}

func (m *module) Dependency(item, dep string) string {
	if deps := m.items[item]; deps != nil {
		if res, found := deps[dep]; found {
			return fmt.Sprintf(attrs.depFmt, dep, res)
		}
	}
	return ""
}
