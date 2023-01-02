// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

// Package app implements a Smart Builder application.
// It is the next generation of building applications using independent bussiness components.
package app

import "github.com/hashicorp/go-plugin"

// SmartBuilder manages modules and builds the application.
type SmartBuilder struct {
	Builder         interface{}
	ModManager      ModManager
	PluginHandshake plugin.HandshakeConfig
	Logger          Logger
}

// SmartGenerator generates smart builder unit (.sb) using smart application unit.
type SmartGenerator struct{}

// ModManager describes methods for managing a module.
type ModManager interface {
	AddItem(moduleName, itemName string) error
	AddDependency(itemName, dependencyName, resolver string, update bool) error
	DeleteItem(itemName string) error
	DeleteDependency(itemName, dependencyName string) error
	ReadAll(kind string) (ModReader, error)
	SetLogger(logger Logger)
}

// ModReader describes methods for getting module attributes.
type ModReader interface {
	Items() map[string]map[string]string
	Dependency(itemName, dependencyName string) string
}

// ModHelper performs usuful methods.
type ModHelper struct {
	Manager ModManager
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

var ModKind = struct {
	SA string
	SB string
	SP string
}{
	"sa", // smart application unit
	"sb", // smart builder unit
	"sp", // smart package unit
}

const (
	// application
	AppName           string = "sb"
	AppVersion        string = "1.0"
	AppVersionString  string = AppName + " version " + AppVersion
	AppsItemName      string = "apps"
	DefaultModuleName string = "apps"
	// errors
	ErrorMessageF         string = "Error: %v\n"
	AppIsMissing          string = "does not found any application in the apps"
	AppIsMissingF         string = "the selected \"%s\" application is not found"
	AppIsMissingInSystemF string = "the system cannot find the \"%s\" application"
	AppIsNotSpecified     string = "the application is not specified"
	ItemIsMissingF        string = "the %s item is not found"
	AttrIsMissingF        string = "the \"%s\" attribute is missing for \"%s\" application"
)
