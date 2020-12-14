package cmd

type SmartBuilder struct {
	Version VersionCmd
}

func (sb *SmartBuilder) Execute() {
	sb.Version.init()
	rootCmd.AddCommand(&sb.Version.Command)
	Execute()
}
