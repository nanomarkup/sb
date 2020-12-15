package cmd

type SmartBuilderConsole struct {
	Gen     GenCmd
	Build   BuildCmd
	Clean   CleanCmd
	Version VersionCmd
}

func (sb *SmartBuilderConsole) Execute() {
	sb.Gen.init()
	sb.Build.init()
	sb.Clean.init()
	sb.Version.init()
	rootCmd.AddCommand(&sb.Gen.Command)
	rootCmd.AddCommand(&sb.Build.Command)
	rootCmd.AddCommand(&sb.Clean.Command)
	rootCmd.AddCommand(&sb.Version.Command)
	Execute()
}
