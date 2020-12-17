package main

import (
	"github.com/sapplications/sbuilder/src/sb/cmd"
	"github.com/sapplications/sbuilder/src/smod"
	"github.com/sapplications/sbuilder/src/golang"
	"github.com/sapplications/sbuilder/src/sb/app"
)

func Execute() {
	app := UseCmdSmartBuilderConsole()
	app.Execute()
}

func UseAppDepManagerRef() *app.DepManager {
	v := &app.DepManager{}
	v.Module = UseSmodConfigFileRef()
	return v
}

func UseCmdGenCmd() cmd.GenCmd {
	v := cmd.GenCmd{}
	v.Short = "Generate configuration"
	v.Long = "Generates all items for the selected configuration.\n\n'gen' generates all items for the current configuration (update).\n'gen [configuration]' generates all items for a custom configuration."
	v.Gen = UseAppSmartBuilderRef()
	v.Use = "gen"
	return v
}

func UseSmodConfigFileRef() *smod.ConfigFile {
	return &smod.ConfigFile{}
}

func UseGolangBuilderRef() *golang.Builder {
	return &golang.Builder{}
}

func UseCmdCleanCmd() cmd.CleanCmd {
	v := cmd.CleanCmd{}
	v.Short = "Remove generated files"
	v.Long = "Clean removes the generated and built files.\n\n'clean' removes files for the current configuration.\n'clean [configuration]'	removes files for a custom configuration."
	v.Clean = UseAppSmartBuilderRef()
	v.Use = "clean"
	return v
}

func UseCmdVersionCmd() cmd.VersionCmd {
	v := cmd.VersionCmd{}
	v.Long = "Version prints the current Smart Builder version."
	v.Version = UseAppSmartBuilderRef()
	v.Use = "version"
	v.Short = "Print Smart Builder version"
	return v
}

func UseCmdSmartBuilderConsole() cmd.SmartBuilderConsole {
	v := cmd.SmartBuilderConsole{}
	v.Dep = UseCmdDepCmd()
	v.Gen = UseCmdGenCmd()
	v.Build = UseCmdBuildCmd()
	v.Clean = UseCmdCleanCmd()
	v.Version = UseCmdVersionCmd()
	return v
}

func UseCmdBuildCmd() cmd.BuildCmd {
	v := cmd.BuildCmd{}
	v.Use = "build"
	v.Short = "Build application"
	v.Long = "Builds an application using the generated items.\n\n'build' builds an application for the current configuration (rebuild).\n'build [configuration]' builds an application for a custom configuration."
	v.Build = UseAppSmartBuilderRef()
	return v
}

func UseAppSmartBuilderRef() *app.SmartBuilder {
	v := &app.SmartBuilder{}
	v.Module = UseSmodConfigFileRef()
	v.GoBuilder = UseGolangBuilderRef()
	return v
}

func UseCmdDepCmd() cmd.DepCmd {
	v := cmd.DepCmd{}
	v.Use = "dep"
	v.Short = "Manage dependencies"
	v.Long = "Manages application dependencies and configurations for generating items to build application.\n\n'dep init [language]' generates a smart module in the current directory, in effect creating a new application rooted at the current directory.\n'dep add --name [item]' adds a new item.\n'dep add --name [item] --dep [dependency] --resolver [resolver]' adds a new dependency item to the existing item.\n'dep del --name [item]' deletes the item with all dependencies.\n'dep del --name [item] --dep [dependency]' deletes item's dependency.\n'dep edit --name [item] --dep [dependency] --resolver [resolver]' adds/updates dependency to/in the existing item.\n'dep list --name [item]' prints item's dependencies.\n'dep list --name [item] --dep [dependency]' prints resolver of dependency item.\n'dep list --version' prints module version.\n'dep list --lang' prints module language.\n'dep list --all' prints module file."
	v.Manager = UseAppDepManagerRef()
	return v
}

