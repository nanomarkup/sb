package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services"
)

type DepManager struct {
	Lang   func() string
	Module services.Module
}

func (d *DepManager) Init(lang string) error {
	common.Check(d.Module.Load(lang))
	if _, err := d.Module.Main(); err == nil {
		return fmt.Errorf("The main item of %s language already exists", lang)
	} else {
		d.Module.Init(Version, lang)
		if err := d.Module.Save(ModFileName); err != nil {
			return err
		}
		fmt.Printf("%s file has been created", ModFileName)
	}
	return nil
}

func (d *DepManager) AddItem(item string) error {
	common.Check(d.Module.Load(d.Lang()))
	common.Check(d.Module.AddItem(item))
	return d.Module.Save(ModFileName)
}

func (d *DepManager) AddDependency(item, dependency, resolver string, update bool) error {
	common.Check(d.Module.Load(d.Lang()))
	common.Check(d.Module.AddDependency(item, dependency, resolver, update))
	return d.Module.Save(ModFileName)
}

func (d *DepManager) DeleteItem(item string) error {
	common.Check(d.Module.Load(d.Lang()))
	common.Check(d.Module.DeleteItem(item))
	return d.Module.Save(ModFileName)
}

func (d *DepManager) DeleteDependency(item, dependency string) error {
	common.Check(d.Module.Load(d.Lang()))
	common.Check(d.Module.DeleteDependency(item, dependency))
	return d.Module.Save(ModFileName)
}
