package sbuilder

type Builder interface {
	Init(Items)
	Build(AppName) error
	Clean(AppName) error
}
