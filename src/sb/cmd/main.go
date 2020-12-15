package cmd

type SmartBuilderConsole struct {
	Clean   CleanCmd
	Version VersionCmd
}

func (sb *SmartBuilderConsole) Execute() {
	sb.Clean.init()
	sb.Version.init()
	rootCmd.AddCommand(&sb.Clean.Command)
	rootCmd.AddCommand(&sb.Version.Command)
	Execute()
}
