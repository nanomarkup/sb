package cmd

type SmartBuilderConsole struct {
	Dep     DepCmd
	Gen     GenCmd
	Build   BuildCmd
	Clean   CleanCmd
	Version VersionCmd
}

func (sb *SmartBuilderConsole) Execute() {
	sb.Dep.init()
	sb.Gen.init()
	sb.Build.init()
	sb.Clean.init()
	sb.Version.init()
	rootCmd.AddCommand(&sb.Dep.Command)
	rootCmd.AddCommand(&sb.Gen.Command)
	rootCmd.AddCommand(&sb.Build.Command)
	rootCmd.AddCommand(&sb.Clean.Command)
	rootCmd.AddCommand(&sb.Version.Command)
	Execute()
}
