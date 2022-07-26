package cmd

import (
	"github.com/sapplications/sbuilder/src/app"
	"github.com/sapplications/sbuilder/src/cmd"
	src "github.com/sapplications/sbuilder/src/smodule"
)

type appSmartBuilder struct {
	app.SmartBuilder
}

func (b *appSmartBuilder) ReadAll(lang string) (cmd.ModReader, error) {
	return b.SmartBuilder.ReadAll(lang)
}

type smoduleManager struct {
	Lang func() string
	src.Manager
}

func (m *smoduleManager) AddItem(module, item string) error {
	m.Manager.Lang = m.Lang
	return m.Manager.AddItem(module, item)
}

func (m *smoduleManager) AddDependency(item, dependency, resolver string, update bool) error {
	m.Manager.Lang = m.Lang
	return m.Manager.AddDependency(item, dependency, resolver, update)
}

func (m *smoduleManager) DeleteItem(item string) error {
	m.Manager.Lang = m.Lang
	return m.Manager.DeleteItem(item)
}

func (m *smoduleManager) DeleteDependency(item, dependency string) error {
	m.Manager.Lang = m.Lang
	return m.Manager.DeleteDependency(item, dependency)
}

func (m *smoduleManager) ReadAll(language string) (app.ModReader, error) {
	m.Manager.Lang = m.Lang
	return m.Manager.ReadAll(language)
}

func (m *smoduleManager) SetLogger(logger app.Logger) {
}
