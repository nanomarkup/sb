package app

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services"
)

type DepManager struct {
	Module services.Module
}

func (d *DepManager) Init(lang string) error {
	if _, err := os.Stat(ModFileName); err == nil {
		return fmt.Errorf("%s already exists", ModFileName)
	} else if !os.IsNotExist(err) {
		return err
	} else {
		d.Module.Init(Version, lang)
		if err := d.Module.SaveToFile(ModFileName); err != nil {
			return err
		}
		fmt.Printf("%s file has been created", ModFileName)
	}
	return nil
}

func (d *DepManager) AddItem(item string) error {
	common.Check(d.Module.LoadFromFile(ModFileName))
	common.Check(d.Module.AddItem(item))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) AddDependency(item, dependency, resolver string, update bool) error {
	common.Check(d.Module.LoadFromFile(ModFileName))
	common.Check(d.Module.AddDependency(item, dependency, resolver, update))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteItem(item string) error {
	common.Check(d.Module.LoadFromFile(ModFileName))
	common.Check(d.Module.DeleteItem(item))
	return d.Module.SaveToFile(ModFileName)
}

func (d *DepManager) DeleteDependency(item, dependency string) error {
	common.Check(d.Module.LoadFromFile(ModFileName))
	common.Check(d.Module.DeleteDependency(item, dependency))
	return d.Module.SaveToFile(ModFileName)
}
