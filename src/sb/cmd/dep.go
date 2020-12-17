// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"
	"strings"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/sb/app"
	"github.com/sapplications/sbuilder/src/smod"
	"github.com/spf13/cobra"
)

type IManager interface {
	Init(lang string)
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
			cli.PrintError("Subcommand is required")
			return
		}
		defer cli.Recover()
		var subCmd = args[0]
		var itemStr = strings.Trim(*depCmdFlags.item, "\t \n")
		var depStr = strings.Trim(*depCmdFlags.dep, "\t \n")
		var resolverStr = strings.Trim(*depCmdFlags.resolver, "\t \n")
		// handle subcommands
		switch subCmd {
		case subCmds.init:
			if len(args) < 2 {
				cli.PrintError("Language parameter is required")
				return
			}
			v.Manager.Init(args[1])
		case subCmds.add:
			if itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr != "" && resolverStr == "" {
				cli.PrintError("\"--resolver\" parameter is required")
				return
			}
			if depStr == "" {
				cli.Check(v.Manager.AddItem(itemStr))
			} else {
				cli.Check(v.Manager.AddDependency(itemStr, depStr, resolverStr, false))
			}
		case subCmds.del:
			if itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				cli.Check(v.Manager.DeleteItem(itemStr))
			} else {
				cli.Check(v.Manager.DeleteDependency(itemStr, depStr))
			}
		case subCmds.edit:
			if itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				cli.PrintError("\"--dep\" parameter is required")
				return
			}
			if resolverStr == "" {
				cli.PrintError("\"--resolver\" parameter is required")
				return
			}
			cli.Check(v.Manager.AddDependency(itemStr, depStr, resolverStr, true))
		case subCmds.list:
			if depStr != "" && itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
			var c smod.ConfigFile
			cli.Check(c.LoadFromFile(app.ModFileName))
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
						cli.PrintError(fmt.Sprintf("\"%s\" item does not exist", itemStr))
					} else {
						if depStr == "" {
							fmt.Printf(c.Item(itemStr))
						} else {
							if _, found := item[depStr]; found {
								fmt.Printf(c.Dependency(itemStr, depStr))
							} else {
								cli.PrintError(fmt.Sprintf("\"%s\" dependency item does not exist", depStr))
							}
						}
					}
				}
			}
		default:
			cli.PrintError(fmt.Sprintf("Unknown \"%s\" subcommand", args[0]))
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
