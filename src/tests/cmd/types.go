package cmd

import (
	"github.com/sapplications/sbuilder/src/app"
	"github.com/sapplications/sbuilder/src/golang"
	"github.com/sapplications/sbuilder/src/services/smodule"
	src "github.com/sapplications/sbuilder/src/smodule"
)

func appItemsToMaps(l map[string]map[string]string) map[string]map[string]string {
	r := map[string]map[string]string{}
	for k, v := range l {
		r[k] = v
	}
	return r
}

func mapsToAppItems(l map[string]map[string]string) map[string]map[string]string {
	r := map[string]map[string]string{}
	for k, v := range l {
		r[k] = v
	}
	return r
}

type srcReader struct {
	imp app.Reader
	src.Module
}

type appSmartBuilder struct {
	app.SmartBuilder
}

func (b *appSmartBuilder) ReadAll(lang string) (smodule.Reader, error) {
	// return r.(smodule.Reader), e
	// return appReaderToSmoduleReader(r), e
	r, e := b.SmartBuilder.ReadAll(lang)
	if e == nil {
		return &srcReader{imp: r}, nil
	} else {
		return nil, e
	}
}

type golangBuilder struct {
	golang.Builder
}

func (b *golangBuilder) Init(items map[string]map[string]string) {
	// r := map[string]map[string]string{}
	// for k, v := range items {
	// 	r[k] = v
	// }
	b.Builder.Init(appItemsToMaps(items))
}

type golangGenerator struct {
	golang.Generator
}

func (g *golangGenerator) Init(items map[string]map[string]string) {
	// r := map[string]map[string]string{}
	// for k, v := range items {
	// 	r[k] = v
	// }
	g.Generator.Init(appItemsToMaps(items))
}

type appReader struct {
	imp smodule.Reader
	src.Module
}

func (r *appReader) Lang() string {
	return r.imp.Lang()
}

func (r *appReader) Items() map[string]map[string]string {
	return mapsToAppItems(r.imp.Items())
}

func (r *appReader) Dependency(itemName, dependencyName string) string {
	return r.imp.Dependency(itemName, dependencyName)
}

func (r *appReader) Main() (map[string]string, error) {
	return r.imp.Main()
}

type smoduleManager struct {
	Lang func() string
	src.Manager
}

func (m *smoduleManager) ReadAll(language string) (app.Reader, error) {
	m.Manager.Lang = m.Lang
	r, e := m.Manager.ReadAll(language)
	if e == nil {
		return &appReader{imp: r}, nil
	} else {
		return nil, e
	}
}
