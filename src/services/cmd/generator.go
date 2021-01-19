package cmd

type Generator interface {
	Generate(application string) error
}
