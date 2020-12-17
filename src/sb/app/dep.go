package app

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/services"
)

type DepManager struct {
	Module services.IModule
}

func (d *DepManager) Init(lang string) {
	if _, err := os.Stat(ModFileName); err == nil {
		cli.PrintError(fmt.Sprintf("%s already exists", ModFileName))
	} else if !os.IsNotExist(err) {
		cli.PrintError(err)
	} else {
		d.Module.Init(Version, lang)
		cli.Check(d.Module.SaveToFile(ModFileName))
		fmt.Printf("%s file has been created", ModFileName)
	}
}

func (d *DepManager) AddItem(item string) error {
	cli.Check(d.Module.LoadFromFile(ModFileName))
	cli.Check(d.Module.AddItem(item))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) AddDependency(item, dependency, resolver string, update bool) error {
	cli.Check(d.Module.LoadFromFile(ModFileName))
	cli.Check(d.Module.AddDependency(item, dependency, resolver, update))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteItem(item string) error {
	cli.Check(d.Module.LoadFromFile(ModFileName))
	cli.Check(d.Module.DeleteItem(item))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteDependency(item, dependency string) error {
	cli.Check(d.Module.LoadFromFile(ModFileName))
	cli.Check(d.Module.DeleteDependency(item, dependency))
	return d.Module.SaveToFile(ModFileName)
}
