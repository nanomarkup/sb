package services

type IBuilder interface {
	Init(items map[string]map[string]string)
	Build(configuration string) error
	Clean(configuration string) error
}
