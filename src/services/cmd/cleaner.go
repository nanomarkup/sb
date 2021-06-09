package cmd

type Cleaner interface {
	Clean(AppName) error
}
