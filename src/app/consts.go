// Package app represents application items
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package app

const (
	// AppName constant returns application name
	AppName string = "sb"
	// Version constant returns the current version of application
	AppVersion string = "1.0"
	// AppVersion constant returns application version
	AppVersionString string = AppName + " version " + AppVersion
	// ModFileName constant returns smart module file name
	DefaultModuleFileName string = "main.sb"
	// error messages
	LanguageIsNotSupportedF string = "the current \"%s\" language is not supported"
)

var Langs = struct {
	Go string
}{
	"go",
}

var suppLangs = map[string]bool{
	Langs.Go: true,
}
