package cmd

type Generator interface {
	Generate(AppName) error
}
