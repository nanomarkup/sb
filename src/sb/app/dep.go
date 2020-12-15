package app

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/smod"
)

type DepManager struct {
}

func (d *DepManager) Init(lang string) {
	if _, err := os.Stat(ModFileName); err == nil {
		cli.PrintError(fmt.Sprintf("%s already exists", ModFileName))
	} else if !os.IsNotExist(err) {
		cli.PrintError(err)
	} else {
		c := smod.ConfigFile{
			Sb:   Version,
			Lang: lang,
			Items: map[string]map[string]string{
				"main": map[string]string{},
			},
		}
		cli.Check(c.SaveToFile(ModFileName))
		fmt.Printf("%s file has been created", ModFileName)
	}
}

func (d *DepManager) AddItem(item string) error {
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	cli.Check(c.AddItem(item))
	return c.SaveToFile(ModFileName)
}

func (d *DepManager) AddDependency(item, dependency, resolver string, update bool) error {
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	cli.Check(c.AddDependency(item, dependency, resolver, update))
	return c.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteItem(item string) error {
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	cli.Check(c.DeleteItem(item))
	return c.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteDependency(item, dependency string) error {
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(ModFileName))
	cli.Check(c.DeleteDependency(item, dependency))
	return c.SaveToFile(ModFileName)
}
