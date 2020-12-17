// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"
	"strings"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/sb/app"
	"github.com/sapplications/sbuilder/src/smod"
	"github.com/spf13/cobra"
)

type IManager interface {
	Init(lang string) error
	AddItem(item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
}

type DepCmd struct {
	Manager IManager
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

var depCmdFlags struct {
	item     *string
	dep      *string
	resolver *string
	version  *bool
	lang     *bool
	all      *bool
}

func (v *DepCmd) init() {
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
		var itemStr = strings.Trim(*depCmdFlags.item, "\t \n")
		var depStr = strings.Trim(*depCmdFlags.dep, "\t \n")
		var resolverStr = strings.Trim(*depCmdFlags.resolver, "\t \n")
		// handle subcommands
		switch subCmd {
		case subCmds.init:
			if len(args) < 2 {
				common.PrintError("Language parameter is required")
				return
			}
			v.Manager.Init(args[1])
		case subCmds.add:
			if itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr != "" && resolverStr == "" {
				common.PrintError("\"--resolver\" parameter is required")
				return
			}
			if depStr == "" {
				common.Check(v.Manager.AddItem(itemStr))
			} else {
				common.Check(v.Manager.AddDependency(itemStr, depStr, resolverStr, false))
			}
		case subCmds.del:
			if itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				common.Check(v.Manager.DeleteItem(itemStr))
			} else {
				common.Check(v.Manager.DeleteDependency(itemStr, depStr))
			}
		case subCmds.edit:
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
			common.Check(v.Manager.AddDependency(itemStr, depStr, resolverStr, true))
		case subCmds.list:
			if depStr != "" && itemStr == "" {
				common.PrintError("\"--name\" parameter is required")
				return
			}
			var c smod.ConfigFile
			common.Check(c.LoadFromFile(app.ModFileName))
			if *depCmdFlags.all {
				fmt.Println(c.String())
			} else {
				if *depCmdFlags.version {
					fmt.Printf(c.Version())
				}
				if *depCmdFlags.lang {
					fmt.Printf(c.Language())
				}
				if itemStr != "" {
					var item = c.Items()[itemStr]
					if item == nil {
						common.PrintError(fmt.Sprintf("\"%s\" item does not exist", itemStr))
					} else {
						if depStr == "" {
							fmt.Printf(c.Item(itemStr))
						} else {
							if _, found := item[depStr]; found {
								fmt.Printf(c.Dependency(itemStr, depStr))
							} else {
								common.PrintError(fmt.Sprintf("\"%s\" dependency item does not exist", depStr))
							}
						}
					}
				}
			}
		default:
			common.PrintError(fmt.Sprintf("Unknown \"%s\" subcommand", args[0]))
			return
		}
	}
	depCmdFlags.item = v.Command.Flags().StringP("name", "n", "", "item name")
	depCmdFlags.dep = v.Command.Flags().StringP("dep", "d", "", "dependency name")
	depCmdFlags.resolver = v.Command.Flags().StringP("resolver", "r", "", "resolver")
	depCmdFlags.version = v.Command.Flags().BoolP("version", "v", false, "print version")
	depCmdFlags.lang = v.Command.Flags().BoolP("lang", "l", false, "print language")
	depCmdFlags.all = v.Command.Flags().BoolP("all", "a", false, "print module")
}
