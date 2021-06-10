// Package app represents application items
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package app

const (
	// Version constant returns the current version of application
	Version string = "1.0"
	// AppName constant returns application name
	AppName string = "sb"
	// AppVersion constant returns application version
	AppVersion string = AppName + " version " + Version
	// ModFileName constant returns smart module file name
	ModFileName string = "main.sb"
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
