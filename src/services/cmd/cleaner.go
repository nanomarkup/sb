package cmd

type Cleaner interface {
	Clean(application string) error
}
