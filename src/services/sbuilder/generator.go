package sbuilder

type Generator interface {
	Init(Items)
	Clean(AppName) error
	Generate(AppName) error
}
