package plugins

type Builder interface {
	Build(app string, sources *map[string]map[string]string) error
	Clean(app string, sources *map[string]map[string]string) error
	Generate(app string, sources *map[string]map[string]string) error
}

type BuilderPlugin struct {
	Impl Builder
}
