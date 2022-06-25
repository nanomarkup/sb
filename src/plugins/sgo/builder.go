package sgo

func (b *builder) Build(app string, sources *map[string]map[string]string) error {
	b.builder.Init(*sources)
	if err := b.builder.Build(app); err != nil {
		return err
	}
	return nil
}

func (b *builder) Clean(app string, sources *map[string]map[string]string) error {
	// remove the built files
	b.builder.Init(*sources)
	if err := b.builder.Clean(app); err != nil {
		return err
	}
	// remove the generated files
	b.generator.Init(*sources)
	if err := b.generator.Clean(app); err != nil {
		return err
	}
	return nil
}

func (b *builder) Generate(app string, sources *map[string]map[string]string) error {
	b.generator.Init(*sources)
	if err := b.generator.Generate(app); err != nil {
		return err
	}
	return nil
}
