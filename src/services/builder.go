package services

type Builder interface {
	Init(items map[string]map[string]string)
	Build(application string) error
	Clean(application string) error
}
