// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"
	"strings"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/cmd"
	"github.com/sapplications/sbuilder/src/services/smodule"
	"github.com/spf13/cobra"
)

type DepManager struct {
	Manager   cmd.Manager
	Formatter smodule.Formatter
	cobra.Command
}

var subCmds = struct {
	init string
	add  string
	del  string
	edit string
	list string
}{
	"init",
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

func (v *DepManager) init() {
	depFlags.mod = v.Command.Flags().StringP("mod", "m", "", "module name")
	depFlags.item = v.Command.Flags().StringP("name", "n", "", "item name")
	depFlags.dep = v.Command.Flags().StringP("dep", "d", "", "dependency name")
	depFlags.resolver = v.Command.Flags().StringP("resolver", "r", "", "resolver")
	depFlags.all = v.Command.Flags().BoolP("all", "a", false, "print module")
	v.Command.Run = func(cmd *cobra.Command, args []string) {
		if v.Manager == nil {
			return
		}
		if len(args) == 0 {
			common.PrintError("Subcommand is required")
			return
		}
		defer common.Recover()
		var subCmd = args[0]
		var modStr = strings.Trim(*depFlags.mod, "\t \n")
		var itemStr = strings.Trim(*depFlags.item, "\t \n")
		var depStr = strings.Trim(*depFlags.dep, "\t \n")
		var resolverStr = strings.Trim(*depFlags.resolver, "\t \n")
		// handle subcommands
		switch subCmd {
		case subCmds.init:
			if len(args) < 2 {
				common.PrintError("Language parameter is required")
				return
			}
			common.Check(v.Manager.Init(args[1]))
		case subCmds.add:
			if modStr == "" {
				common.PrintError("\"--mod\" parameter is required")
				return
			}
			if itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr != "" && resolverStr == "" {
				common.PrintError("\"--resolver\" parameter is required")
				return
			}
			if depStr == "" {
				common.Check(v.Manager.AddItem(modStr, itemStr))
			} else {
				common.Check(v.Manager.AddDependency(modStr, itemStr, depStr, resolverStr, false))
			}
		case subCmds.del:
			if modStr == "" {
				common.PrintError("\"--mod\" parameter is required")
				return
			}
			if itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				common.Check(v.Manager.DeleteItem(modStr, itemStr))
			} else {
				common.Check(v.Manager.DeleteDependency(modStr, itemStr, depStr))
			}
		case subCmds.edit:
			if modStr == "" {
				common.PrintError("\"--mod\" parameter is required")
				return
			}
			if itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				common.PrintError("\"--dep\" parameter is required")
				return
			}
			if resolverStr == "" {
				common.PrintError("\"--resolver\" parameter is required")
				return
			}
			common.Check(v.Manager.AddDependency(modStr, itemStr, depStr, resolverStr, true))
		case subCmds.list:
			if depStr != "" && itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			mod, err := v.Manager.ReadAll(Language())
			common.Check(err)
			if *depFlags.all {
				fmt.Println(v.Formatter.String(mod))
			} else if itemStr != "" {
				var item = mod.Items()[itemStr]
				if item == nil {
					common.PrintError(fmt.Sprintf("\"%s\" item does not exist", itemStr))
				} else {
					if depStr == "" {
						fmt.Printf(v.Formatter.Item(itemStr, item))
					} else {
						if _, found := item[depStr]; found {
							fmt.Printf(mod.Dependency(itemStr, depStr))
						} else {
							common.PrintError(fmt.Sprintf("\"%s\" dependency item does not exist", depStr))
						}
					}
				}
			}
		default:
			common.PrintError(fmt.Sprintf("Unknown \"%s\" subcommand", args[0]))
			return
		}
	}
}
