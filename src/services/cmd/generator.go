package cmd

type Generator interface {
	Generate(configuration string) error
}
