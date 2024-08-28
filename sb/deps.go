package main

import (
	p3 "github.com/nanomarkup/sb"
	p4 "github.com/nanomarkup/dl"
	p5 "github.com/hashicorp/go-plugin"
	p6 "github.com/spf13/cobra"
	p7 "github.com/nanomarkup/sb/plugins"
	p1 "github.com/hashicorp/go-hclog"
	p2 "github.com/nanomarkup/sb/cmd"
)

func Execute() {
	app := UseCmdSmartBuilder()
	app.Execute()
}

func UseCreatorLoggerGroupGo_HclogLoggerOptionsRef() *p1.LoggerOptions {
	v := &p1.LoggerOptions{}
	v.Name = "sa"
	v.Level = 1
	v.Output = p2.OSStdout()
	return v
}

func UseSbSmartCreatorRef() *p3.SmartCreator {
	v := &p3.SmartCreator{}
	v.ModManager = UseGen1GroupDlManagerSbModManagerAdapterRef()
	v.Logger = p1.New(UseCreatorLoggerGroupGo_HclogLoggerOptionsRef())
	return v
}

func UseGen1GroupDlManagerRef() *p4.Manager {
	v := &p4.Manager{}
	v.Kind = "sa"
	return v
}

func UseGen2GroupDlManagerRef() *p4.Manager {
	v := &p4.Manager{}
	v.Kind = "sb"
	return v
}

func UseGo_PluginHandshakeConfig() p5.HandshakeConfig {
	v := p5.HandshakeConfig{}
	v.ProtocolVersion = 1
	v.MagicCookieKey = "SMART_PLUGIN"
	v.MagicCookieValue = "sbuilder"
	return v
}

func UseModIniterGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	return v
}

func UseDlFormatterRef() *p4.Formatter {
	v := &p4.Formatter{}
	return v
}

func UseAppsPrinterGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "list"
	v.Short = "Print all applications"
	v.Long = "Prints all available applications."
	v.RunE = p2.CmdList(UseSbModHelperRef())
	v.SilenceUsage = true
	return v
}

func UsePluginsBuilderPluginRef() *p7.BuilderPlugin {
	v := &p7.BuilderPlugin{}
	return v
}

func UseGeneratorGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "gen [application]"
	v.Example = "  generate hello - generate 'hello' application"
	v.Short = "Generate smart units"
	v.Long = "Generates smart builder unit (.sb) using smart application unit."
	v.RunE = p2.CmdGen(UseSbSmartGeneratorRef())
	v.SilenceUsage = true
	return v
}

func UseSbSmartGeneratorRef() *p3.SmartGenerator {
	v := &p3.SmartGenerator{}
	return v
}

func UseGen3GroupDlManagerRef() *p4.Manager {
	v := &p4.Manager{}
	v.Kind = "sb"
	return v
}

func UseVersionPrinterGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "version"
	v.Short = "Print Smart Builder version"
	v.Long = "Prints the current Smart Builder version."
	v.Run = p2.CmdVersion(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseRunnerGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "run [application]"
	v.Example = "  run - run the current application\n  run hello - run 'hello' application"
	v.Short = "Run application"
	v.Long = "Runs the application."
	v.RunE = p2.CmdRun(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseModDelerGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	return v
}

func UseBuilderGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "build [application]"
	v.Example = "  build - build the current application\n  build hello - build 'hello' application"
	v.Short = "Build application"
	v.Long = "Builds an application using the generated items."
	v.RunE = p2.CmdBuild(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseModAdderGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	return v
}

func UseSbModHelperRef() *p3.ModHelper {
	v := &p3.ModHelper{}
	v.Manager = UseGen3GroupDlManagerSbModManagerAdapterRef()
	return v
}

func UseCoderGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "code [application]"
	v.Example = "  code - generate sources to build the application\n  code hello - generate sources for 'hello' application"
	v.Short = "Generate code"
	v.Long = "Generates code to build the application."
	v.RunE = p2.CmdCode(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseSbSmartBuilderRef() *p3.SmartBuilder {
	v := &p3.SmartBuilder{}
	v.Builder = UsePluginsBuilderPluginRef()
	v.ModManager = UseGen2GroupDlManagerSbModManagerAdapterRef()
	v.PluginHandshake = UseGo_PluginHandshakeConfig()
	v.Logger = p1.New(UseBuilderLoggerGroupGo_HclogLoggerOptionsRef())
	return v
}

func UseCleanerGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "clean [application]"
	v.Example = "  clean - remove files for the current application\n  clean hello - remove files for 'hello' application"
	v.Short = "Remove generated files"
	v.Long = "Removes generated/compiled files."
	v.RunE = p2.CmdClean(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseCmdSmartBuilder() p2.SmartBuilder {
	v := p2.SmartBuilder{}
	v.Use = "sb"
	v.Short = "Smart Builder (c)"
	v.Long = "Smart Builder is the next generation of building applications using independent bussiness components."
	v.SilenceUsage = true
	v.CompletionOptions.DisableDefaultCmd = true
	v.AddCommand(UseCreatorGroupCobraCommandRef())
	v.AddCommand(UseGeneratorGroupCobraCommandRef())
	v.AddCommand(UseCoderGroupCobraCommandRef())
	v.AddCommand(UseBuilderGroupCobraCommandRef())
	v.AddCommand(UseCleanerGroupCobraCommandRef())
	v.AddCommand(UseRunnerGroupCobraCommandRef())
	v.AddCommand(UseModManagerGroupCobraCommandRef())
	v.AddCommand(UseAppsPrinterGroupCobraCommandRef())
	v.AddCommand(UseVersionPrinterGroupCobraCommandRef())
	return v
}

func UseModManagerGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "mod"
	v.Example = "  mod edit --name hello --dep Writer --resolver FileWriter - add/update 'Writer' dependency item to/in 'hello' item\n  mod list --name hello - print 'hello' item data\n  mod list --name hello --dep Writer - print 'Writer' dependency item data\n  mod list --all - print all data"
	v.Short = "Manage modules"
	v.Long = "Manages application items and dependencies."
	v.RunE = p2.CmdManageMod(UseSbSmartBuilderRef(), UseDlFormatterRef())
	v.SilenceUsage = true
	v.AddCommand(UseModIniterGroupCobraCommandRef())
	v.AddCommand(UseModAdderGroupCobraCommandRef())
	v.AddCommand(UseModDelerGroupCobraCommandRef())
	return v
}

func UseCreatorGroupCobraCommandRef() *p6.Command {
	v := &p6.Command{}
	v.Use = "new [application]"
	v.Example = "  create hello - create 'hello' application"
	v.Short = "Create new application"
	v.Long = "Creates an application by generating smart application unit (.sa file)."
	v.RunE = p2.CmdCreate(UseSbSmartCreatorRef())
	v.SilenceUsage = true
	return v
}

func UseBuilderLoggerGroupGo_HclogLoggerOptionsRef() *p1.LoggerOptions {
	v := &p1.LoggerOptions{}
	v.Name = "sb"
	v.Level = 6
	v.Output = p2.OSStdout()
	return v
}

type Gen1GroupDlManagerSbModManagerAdapter struct {
	p4.Manager
}

func (o *Gen1GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p3.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p3.ModReader)
	return
}

func (o *Gen1GroupDlManagerSbModManagerAdapter) SetLogger(a1 p3.Logger) {
	b1 := a1.(p4.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen1GroupDlManagerSbModManagerAdapterRef() *Gen1GroupDlManagerSbModManagerAdapter {
	v := &Gen1GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen1GroupDlManagerRef()
	return v
}

type Gen3GroupDlManagerSbModManagerAdapter struct {
	p4.Manager
}

func (o *Gen3GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p3.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p3.ModReader)
	return
}

func (o *Gen3GroupDlManagerSbModManagerAdapter) SetLogger(a1 p3.Logger) {
	b1 := a1.(p4.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen3GroupDlManagerSbModManagerAdapterRef() *Gen3GroupDlManagerSbModManagerAdapter {
	v := &Gen3GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen3GroupDlManagerRef()
	return v
}

type Gen2GroupDlManagerSbModManagerAdapter struct {
	p4.Manager
}

func (o *Gen2GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p3.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p3.ModReader)
	return
}

func (o *Gen2GroupDlManagerSbModManagerAdapter) SetLogger(a1 p3.Logger) {
	b1 := a1.(p4.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen2GroupDlManagerSbModManagerAdapterRef() *Gen2GroupDlManagerSbModManagerAdapter {
	v := &Gen2GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen2GroupDlManagerRef()
	return v
}

