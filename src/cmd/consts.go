package cmd

const (
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	SubcmdMissing           string = "subcommand is required"
	ItemMissing             string = "item name is required"
	ModOrDepMissing         string = "module name or dependency name is missing"
	LanguageMissing         string = "language parameter is required"
	ResolverMissing         string = "resolver is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist\n"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist\n"
	UnknownSubcmdF          string = "unknown \"%s\" subcommand\n"
)
