package main

import (
	p6 "github.com/hashicorp/go-hclog"
	p5 "github.com/hashicorp/go-plugin"
	p3 "github.com/nanomarkup/dl"
	p7 "github.com/nanomarkup/sb"
	p2 "github.com/nanomarkup/sb/cmd"
	p4 "github.com/nanomarkup/sb/plugins"
	p1 "github.com/spf13/cobra"
)

func Execute() {
	app := UseCmdSmartBuilder()
	app.Execute()
}

func UseVersionPrinterGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "version"
	v.Short = "Print Smart Builder version"
	v.Long = "Prints the current Smart Builder version."
	v.Run = p2.CmdVersion(UseSbSmartBuilderRef())
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

func UseCleanerGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "clean [application]"
	v.Example = "  clean - remove files for the current application\n  clean hello - remove files for 'hello' application"
	v.Short = "Remove generated files"
	v.Long = "Removes generated/compiled files."
	v.RunE = p2.CmdClean(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseGen3GroupDlManagerRef() *p3.Manager {
	v := &p3.Manager{}
	v.Kind = "sb"
	return v
}

func UseGeneratorGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "gen [application]"
	v.Example = "  generate hello - generate 'hello' application"
	v.Short = "Generate smart units"
	v.Long = "Generates smart builder unit (.sb) using smart application unit."
	v.RunE = p2.CmdGen(UseSbSmartGeneratorRef())
	v.SilenceUsage = true
	return v
}

func UseDlFormatterRef() *p3.Formatter {
	v := &p3.Formatter{}
	return v
}

func UseModIniterGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	return v
}

func UsePluginsBuilderPluginRef() *p4.BuilderPlugin {
	v := &p4.BuilderPlugin{}
	return v
}

func UseGo_PluginHandshakeConfig() p5.HandshakeConfig {
	v := p5.HandshakeConfig{}
	v.ProtocolVersion = 1
	v.MagicCookieKey = "SMART_PLUGIN"
	v.MagicCookieValue = "sbuilder"
	return v
}

func UseBuilderGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "build [application]"
	v.Example = "  build - build the current application\n  build hello - build 'hello' application"
	v.Short = "Build application"
	v.Long = "Builds an application using the generated items."
	v.RunE = p2.CmdBuild(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseModManagerGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
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

func UseCreatorLoggerGroupGo_HclogLoggerOptionsRef() *p6.LoggerOptions {
	v := &p6.LoggerOptions{}
	v.Name = "sa"
	v.Level = 1
	v.Output = p2.OSStdout()
	return v
}

func UseSbSmartGeneratorRef() *p7.SmartGenerator {
	v := &p7.SmartGenerator{}
	return v
}

func UseModDelerGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	return v
}

func UseSbSmartCreatorRef() *p7.SmartCreator {
	v := &p7.SmartCreator{}
	v.ModManager = UseGen1GroupDlManagerSbModManagerAdapterRef()
	v.Logger = p6.New(UseCreatorLoggerGroupGo_HclogLoggerOptionsRef())
	return v
}

func UseRunnerGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "run [application]"
	v.Example = "  run - run the current application\n  run hello - run 'hello' application"
	v.Short = "Run application"
	v.Long = "Runs the application."
	v.RunE = p2.CmdRun(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseModAdderGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	return v
}

func UseCreatorGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "new [application]"
	v.Example = "  create hello - create 'hello' application"
	v.Short = "Create new application"
	v.Long = "Creates an application by generating smart application unit (.sa file)."
	v.RunE = p2.CmdCreate(UseSbSmartCreatorRef())
	v.SilenceUsage = true
	return v
}

func UseAppsPrinterGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "list"
	v.Short = "Print all applications"
	v.Long = "Prints all available applications."
	v.RunE = p2.CmdList(UseSbModHelperRef())
	v.SilenceUsage = true
	return v
}

func UseSbSmartBuilderRef() *p7.SmartBuilder {
	v := &p7.SmartBuilder{}
	v.Builder = UsePluginsBuilderPluginRef()
	v.ModManager = UseGen2GroupDlManagerSbModManagerAdapterRef()
	v.PluginHandshake = UseGo_PluginHandshakeConfig()
	v.Logger = p6.New(UseBuilderLoggerGroupGo_HclogLoggerOptionsRef())
	return v
}

func UseGen2GroupDlManagerRef() *p3.Manager {
	v := &p3.Manager{}
	v.Kind = "sb"
	return v
}

func UseBuilderLoggerGroupGo_HclogLoggerOptionsRef() *p6.LoggerOptions {
	v := &p6.LoggerOptions{}
	v.Name = "sb"
	v.Level = 6
	v.Output = p2.OSStdout()
	return v
}

func UseCoderGroupCobraCommandRef() *p1.Command {
	v := &p1.Command{}
	v.Use = "code [application]"
	v.Example = "  code - generate sources to build the application\n  code hello - generate sources for 'hello' application"
	v.Short = "Generate code"
	v.Long = "Generates code to build the application."
	v.RunE = p2.CmdCode(UseSbSmartBuilderRef())
	v.SilenceUsage = true
	return v
}

func UseGen1GroupDlManagerRef() *p3.Manager {
	v := &p3.Manager{}
	v.Kind = "sa"
	return v
}

func UseSbModHelperRef() *p7.ModHelper {
	v := &p7.ModHelper{}
	v.Manager = UseGen3GroupDlManagerSbModManagerAdapterRef()
	return v
}

type Gen1GroupDlManagerSbModManagerAdapter struct {
	p3.Manager
}

func (o *Gen1GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p7.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p7.ModReader)
	return
}

func (o *Gen1GroupDlManagerSbModManagerAdapter) SetLogger(a1 p7.Logger) {
	b1 := a1.(p3.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen1GroupDlManagerSbModManagerAdapterRef() *Gen1GroupDlManagerSbModManagerAdapter {
	v := &Gen1GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen1GroupDlManagerRef()
	return v
}

type Gen2GroupDlManagerSbModManagerAdapter struct {
	p3.Manager
}

func (o *Gen2GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p7.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p7.ModReader)
	return
}

func (o *Gen2GroupDlManagerSbModManagerAdapter) SetLogger(a1 p7.Logger) {
	b1 := a1.(p3.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen2GroupDlManagerSbModManagerAdapterRef() *Gen2GroupDlManagerSbModManagerAdapter {
	v := &Gen2GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen2GroupDlManagerRef()
	return v
}

type Gen3GroupDlManagerSbModManagerAdapter struct {
	p3.Manager
}

func (o *Gen3GroupDlManagerSbModManagerAdapter) ReadAll() (r1 p7.ModReader, r2 error) {
	v1, r2 := o.Manager.ReadAll()
	r1 = v1.(p7.ModReader)
	return
}

func (o *Gen3GroupDlManagerSbModManagerAdapter) SetLogger(a1 p7.Logger) {
	b1 := a1.(p3.Logger)
	o.Manager.SetLogger(b1)
}

func UseGen3GroupDlManagerSbModManagerAdapterRef() *Gen3GroupDlManagerSbModManagerAdapter {
	v := &Gen3GroupDlManagerSbModManagerAdapter{}
	v.Manager = *UseGen3GroupDlManagerRef()
	return v
}
