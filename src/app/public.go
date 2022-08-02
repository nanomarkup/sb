// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

// Package app implements a Smart Builder application.
// It is the next generation of building applications using independent bussiness components.
package app

import "github.com/hashicorp/go-plugin"

const (
	AppName           string = "sb"
	AppVersion        string = "1.0"
	AppVersionString  string = AppName + " version " + AppVersion
	DefaultModuleName string = "apps"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	LanguageIsNotSupportedF string = "the current \"%s\" language is not supported\n"
	ApplicationIsMissing    string = "does not found any application in the apps"
)

// SmartBuilder manages modules and builds the application.
type SmartBuilder struct {
	Lang            func() string
	Builder         interface{}
	ModManager      ModManager
	PluginHandshake plugin.HandshakeConfig
	Logger          Logger
}

// SmartGenerator generates smart builder unit (.sb) using smart application unit.
type SmartGenerator struct{}

// ModManager describes methods for managing a module.
type ModManager interface {
	Init(moduleName, language string) error
	AddItem(moduleName, itemName string) error
	AddDependency(itemName, dependencyName, resolver string, update bool) error
	DeleteItem(itemName string) error
	DeleteDependency(itemName, dependencyName string) error
	ReadAll(language string) (ModReader, error)
	SetLogger(logger Logger)
}

// ModReader describes methods for getting module attributes.
type ModReader interface {
	Lang() string
	Items() map[string]map[string]string
	Dependency(itemName, dependencyName string) string
	Apps() (map[string]string, error)
}

// Logger describes methods for logging messages.
type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	IsError() bool
}
