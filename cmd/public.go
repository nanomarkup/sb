// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

// Package cmd represents Command Line Interface.
package cmd

import "github.com/spf13/cobra"

// SmartBuilder includes all available commands and handles them.
type SmartBuilder struct {
	Starter      Starter
	Reader       CmdReader
	Runner       CmdRunner
	Creator      CmdCreator
	Generator    CmdGenerator
	Coder        CmdCoder
	Builder      CmdBuilder
	Cleaner      CmdCleaner
	Printer      CmdPrinter
	ModManager   CmdManager
	ModAdder     CmdModAdder
	ModDeler     CmdModDeler
	ModIniter    CmdModIniter
	SilentErrors bool
}

// Starter command is the apps command contains all commands to display.
type Starter struct {
	cobra.Command
}

// CmdReader command displays a version of the application.
type CmdReader struct {
	Reader
	cobra.Command
}

// Reader describes methods for displaying a version of the application.
type Reader interface {
	Version() string
}

// CmdCreator command creates an application by generating smart application unit (.sa file).
type CmdCreator struct {
	Creator
	cobra.Command
}

// Creator describes methods for creating an application by generating smart application unit (.sa file).
type Creator interface {
	Create(string) error
}

// CmdGenerator command generates smart builder unit (.sb) using smart application unit.
type CmdGenerator struct {
	Generator
	cobra.Command
}

// Generator describes methods for generating smart builder unit (.sb) using smart application unit.
type Generator interface {
	Generate(string) error
}

// CmdCoder command generates code to build the application.
type CmdCoder struct {
	Coder
	cobra.Command
}

// Coder describes methods for generating code to build the application.
type Coder interface {
	Generate(string) error
}

// CmdBuilder command builds an application using the generated items.
type CmdBuilder struct {
	Builder
	cobra.Command
}

// Builder describes methods for building an application using the generated items.
type Builder interface {
	Build(string) error
}

// CmdCleaner command removes generated/compiled files.
type CmdCleaner struct {
	Cleaner
	cobra.Command
}

// Cleaner describes methods for removing generated/compiled files.
type Cleaner interface {
	Clean(string) error
}

// CmdRunner command runs the application.
type CmdRunner struct {
	Runner
	cobra.Command
}

// Runner describes methods for running the application.
type Runner interface {
	Run(string) error
}

// Printer describes methods for displaying all available applications.
type Printer interface {
	Apps() ([]string, error)
}

// CmdPrinter command displays all available applications.
type CmdPrinter struct {
	Printer
	cobra.Command
}

// CmdManager command manages application items and dependencies.
type CmdManager struct {
	ModManager
	ModFormatter
	cobra.Command
}

// ModManager describes methods for managing application items and dependencies.
type ModManager interface {
	AddItem(module, item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	ReadAll(kind string) (ModReader, error)
}

// ModReader describes methods for getting module attributes.
type ModReader interface {
	Items() map[string]map[string]string
	Dependency(string, string) string
}

// ModFormatter describes methods for formatting module attributes and returns it as a string.
type ModFormatter interface {
	Item(string, map[string]string) string
	String(module ModReader) string
}

// CmdModAdder command adds item or dependency to the exsiting item.
type CmdModAdder struct {
	ModManager
	cobra.Command
}

// CmdModDeler command deletes item or dependency from the exsiting item.
type CmdModDeler struct {
	ModManager
	cobra.Command
}

// CmdModIniter command creates a apps.sb module and initialize it with the apps item.
// If the apps item is exist then do nothing.
type CmdModIniter struct {
	ModManager
	cobra.Command
}

const (
	// application
	AppsItemName      string = "apps"
	DefaultModuleName string = "apps"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	AppNameMissing          string = "application name is required"
	SubcmdMissing           string = "subcommand is required"
	ItemMissing             string = "item name is required"
	ModOrDepMissing         string = "module name or dependency name is missing"
	ResolverMissing         string = "resolver is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist"
	UnknownSubcmdF          string = "unknown \"%s\" subcommand"
)
