package cmd

import (
	"github.com/sapplications/sb/app"
	"github.com/sapplications/sb/cmd"
	"github.com/sapplications/smod/lod"
)

var modType = struct {
	sa string
	sb string
	sp string
}{
	"sa",
	"sb",
	"sp",
}

type appSmartBuilder struct {
	app.SmartBuilder
}

func (b *appSmartBuilder) ReadAll(kind string) (cmd.ModReader, error) {
	return b.SmartBuilder.ReadAll(kind)
}

type smoduleManager struct {
	lod.Manager
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

func (m *smoduleManager) ReadAll(kind string) (app.ModReader, error) {
	return m.Manager.ReadAll(kind)
}

func (m *smoduleManager) SetLogger(logger app.Logger) {
}
