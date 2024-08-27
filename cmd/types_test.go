package cmd

import (
	"github.com/nanomarkup/dl"
	"github.com/nanomarkup/sb"
)

type appSmartBuilder struct {
	sb.SmartBuilder
}

func (b *appSmartBuilder) ReadAll(kind string) (ModReader, error) {
	return b.SmartBuilder.ReadAll(kind)
}

type smoduleManager struct {
	dl.Manager
}

func (m *smoduleManager) AddItem(module, item string) error {
	return m.Manager.AddItem(module, item)
}

func (m *smoduleManager) AddDependency(item, dependency, resolver string, update bool) error {
	return m.Manager.AddDependency(item, dependency, resolver, update)
}

func (m *smoduleManager) DeleteItem(item string) error {
	return m.Manager.DeleteItem(item)
}

func (m *smoduleManager) DeleteDependency(item, dependency string) error {
	return m.Manager.DeleteDependency(item, dependency)
}

func (m *smoduleManager) ReadAll() (sb.ModReader, error) {
	return m.Manager.ReadAll()
}

func (m *smoduleManager) SetLogger(logger sb.Logger) {
}
