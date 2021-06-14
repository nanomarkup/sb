// Package app represents application items
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package app

const (
	AppName           string = "sb"
	AppVersion        string = "1.0"
	AppVersionString  string = AppName + " version " + AppVersion
	DefaultModuleName string = "main"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	LanguageIsNotSupportedF string = "the current \"%s\" language is not supported\n"
	ApplicationIsMissing    string = "does not found any application in the main"
)

var Langs = struct {
	Go string
}{
	"go",
}

var suppLangs = map[string]bool{
	Langs.Go: true,
}
