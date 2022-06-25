package sgo

type Plugin struct {
	Builder   Builder
	Generator Generator
}

type Builder interface {
	Init(items map[string]map[string]string)
	Build(appName string) error
	Clean(appName string) error
}

type Generator interface {
	Init(items map[string]map[string]string)
	Clean(appName string) error
	Generate(appName string) error
}
