package cmd

type Builder interface {
	Build(configuration string) error
}
