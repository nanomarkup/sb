package cmd

type Builder interface {
	Build(application string) error
}
