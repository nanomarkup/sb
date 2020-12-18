package services

type Generator interface {
	Init(items map[string]map[string]string)
	Clean(сonfiguration string) error
	Generate(сonfiguration string) error
}
