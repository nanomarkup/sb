package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/golang"
	"github.com/sapplications/sbuilder/src/smod"
)

type SmartBuilder struct {
}

func (sb *SmartBuilder) Build(configuration string) {
	defer cli.Recover()
	// check configuration
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	if err := CheckConfiguration(configuration, &c); err != nil {
		cli.PrintError(err)
		return
	}
	// process configuration
	switch c.Lang {
	case Langs.Go:
		var builder = golang.Builder{
			ModFileName,
			configuration,
		}
		if err := builder.Build(&c); err != nil {
			cli.PrintError(err)
		}
	default:
		cli.PrintError(fmt.Errorf("\"%s\" language is not supported"))
	}
}

func (sb *SmartBuilder) Clean(configuration string) {
	defer cli.Recover()
	// check configuration
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	if err := CheckConfiguration(configuration, &c); err != nil {
		cli.PrintError(err)
		return
	}
	// process configuration
	switch c.Lang {
	case Langs.Go:
		// remove the generated files
		var gen = golang.Generator{
			ModFileName,
			configuration,
		}
		if err := gen.Clean(); err != nil {
			cli.PrintError(err)
		}
		// remove the built files
		var builder = golang.Builder{
			ModFileName,
			configuration,
		}
		if err := builder.Clean(); err != nil {
			cli.PrintError(err)
		}
	default:
		cli.PrintError(fmt.Errorf("\"%s\" language is not supported"))
	}
}

func (sb *SmartBuilder) PrintVersion() {
	fmt.Println(AppVersion)
}
