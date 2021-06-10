package cmd

const (
	SubcmdMissing           string = "Subcommand is required"
	NameMissing             string = "\"--name\" parameter is required"
	ModuleMissing           string = "\"--mod\" parameter is required"
	LanguageMissing         string = "Language parameter is required"
	ResolverMissing         string = "\"--resolver\" parameter is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist"
	UnknownSubcmdF          string = "Unknown \"%s\" subcommand"
)
