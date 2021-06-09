package cmd

type Builder interface {
	Build(AppName) error
}
