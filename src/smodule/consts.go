package smodule

const (
	moduleExt    string = ".sb"
	MainItemName string = "main"
	// notifications
	ModuleIsCreatedF string = "%s file has been created\n"
	// errors
	ItemExistsF             string = "the %s item already exists in %s module"
	ItemIsMissingF          string = "the %s item does not exist"
	ModuleFilesMissingF     string = "no sb files in %s"
	ModuleLanguageMismatchF string = "the %s language of %s module is mismatch the %s selected language"
)
