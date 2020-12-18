package cmd

type SmartBuilder struct {
	Reader     Reader
	Builder    Builder
	Cleaner    Cleaner
	Generator  Generator
	DepManager DepManager
}

func (sb *SmartBuilder) Execute() {
	sb.Reader.init()
	sb.Builder.init()
	sb.Cleaner.init()
	sb.Generator.init()
	sb.DepManager.init()
	rootCmd.AddCommand(&sb.DepManager.Command)
	rootCmd.AddCommand(&sb.Generator.Command)
	rootCmd.AddCommand(&sb.Builder.Command)
	rootCmd.AddCommand(&sb.Cleaner.Command)
	rootCmd.AddCommand(&sb.Reader.Command)
	Execute()
}
