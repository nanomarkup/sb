// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sapplications/sbuilder/src/common"
	src "github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/sapplications/sbuilder/src/services/smodule"
	"github.com/spf13/cobra"
)

type Manager struct {
	Manager   src.Manager
	Formatter smodule.Formatter
	cobra.Command
}

var subCmds = struct {
	add  string
	del  string
	edit string
	list string
}{
	"add",
	"del",
	"edit",
	"list",
}

var depFlags struct {
	mod      *string
	item     *string
	dep      *string
	resolver *string
	all      *bool
}

func (v *Manager) init() {
	depFlags.mod = v.Command.Flags().StringP("mod", "m", "", "module name")
	depFlags.item = v.Command.Flags().StringP("name", "n", "", "item name")
	depFlags.dep = v.Command.Flags().StringP("dep", "d", "", "dependency name")
	depFlags.resolver = v.Command.Flags().StringP("resolver", "r", "", "resolver")
	depFlags.all = v.Command.Flags().BoolP("all", "a", false, "print module")
	v.SilenceUsage = true
	v.Command.RunE = func(cmd *cobra.Command, args []string) error {
		if v.Manager == nil {
			return nil
		}
		if len(args) == 0 {
			return errors.New(SubcmdMissing)
		}
		defer common.Recover()
		var subCmd = args[0]
		var modStr = strings.Trim(*depFlags.mod, "\t \n")
		var itemStr = strings.Trim(*depFlags.item, "\t \n")
		var depStr = strings.Trim(*depFlags.dep, "\t \n")
		var resolverStr = strings.Trim(*depFlags.resolver, "\t \n")
		// handle subcommands
		switch subCmd {
		case subCmds.add:
			if modStr == "" {
				return errors.New(ModuleMissing)
			}
			if itemStr == "" {
				return errors.New(NameMissing)
			}
			if depStr != "" && resolverStr == "" {
				return errors.New(ResolverMissing)
			}
			if depStr == "" {
				return v.Manager.AddItem(modStr, itemStr)
			} else {
				return v.Manager.AddDependency(modStr, itemStr, depStr, resolverStr, false)
			}
		case subCmds.del:
			if modStr == "" {
				return errors.New(ModuleMissing)
			}
			if itemStr == "" {
				return errors.New(NameMissing)
			}
			if depStr == "" {
				return v.Manager.DeleteItem(modStr, itemStr)
			} else {
				return v.Manager.DeleteDependency(modStr, itemStr, depStr)
			}
		case subCmds.edit:
			if modStr == "" {
				return errors.New(ModuleMissing)
			}
			if itemStr == "" {
				return errors.New(NameMissing)
			}
			if depStr == "" {
				return errors.New(DependencyMissing)
			}
			if resolverStr == "" {
				return errors.New(ResolverMissing)
			}
			return v.Manager.AddDependency(modStr, itemStr, depStr, resolverStr, true)
		case subCmds.list:
			if depStr != "" && itemStr == "" {
				return errors.New(NameMissing)
			}
			mod, err := v.Manager.ReadAll(Language())
			common.Check(err)
			if *depFlags.all {
				fmt.Println(v.Formatter.String(mod))
			} else if itemStr != "" {
				var item = mod.Items()[itemStr]
				if item == nil {
					return fmt.Errorf(ItemDoesNotExistF, itemStr)
				} else {
					if depStr == "" {
						fmt.Printf(v.Formatter.Item(itemStr, item))
					} else {
						if _, found := item[depStr]; found {
							fmt.Printf(mod.Dependency(itemStr, depStr))
						} else {
							return fmt.Errorf(DependencyDoesNotExistF, depStr)
						}
					}
				}
			}
		default:
			return fmt.Errorf(UnknownSubcmdF, args[0])
		}
		return nil
	}
}
