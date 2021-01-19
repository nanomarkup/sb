package services

type Generator interface {
	Init(items map[string]map[string]string)
	Clean(application string) error
	Generate(application string) error
}
