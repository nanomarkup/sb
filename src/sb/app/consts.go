// Package app represents application items
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package app

const (
	// Version constant returns the current version of application
	Version = "1.0"
	// AppName constant returns application name
	AppName = "sb"
	// AppVersion constant returns application version
	AppVersion = AppName + " version " + Version
	// ModFileName constant returns smart module file name
	ModFileName = "smod.sm"
)

var Langs = struct {
	Go string
}{
	"go",
}

var suppLangs = map[string]bool{
	Langs.Go: true,
}

var versions = map[string]bool{
	"1.0": true,
}
