package cmd

type Cleaner interface {
	Clean(configuration string) error
}
