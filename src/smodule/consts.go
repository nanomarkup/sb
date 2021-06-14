package smodule

const (
	moduleExt    string = ".sb"
	MainItemName string = "main"
	// error messages
	ItemExistsF             string = "the %s item already exists in %s module\n"
	ItemIsMissingF          string = "the %s item does not exist\n"
	ModuleIsCreatedF        string = "%s file has been created\n"
	ModuleFilesMissingF     string = "no sb files in %s\n"
	ModuleLanguageMismatchF string = "the %s language of %s module is mismatch the %s selected language\n"
)
