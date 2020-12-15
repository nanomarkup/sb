package cmd

type SmartBuilderConsole struct {
	Build   BuildCmd
	Clean   CleanCmd
	Version VersionCmd
}

func (sb *SmartBuilderConsole) Execute() {
	sb.Build.init()
	sb.Clean.init()
	sb.Version.init()
	rootCmd.AddCommand(&sb.Build.Command)
	rootCmd.AddCommand(&sb.Clean.Command)
	rootCmd.AddCommand(&sb.Version.Command)
	Execute()
}
