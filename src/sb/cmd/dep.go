// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/sb/app"
	"github.com/sapplications/sbuilder/src/smod"
	"github.com/spf13/cobra"
)

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

// depCmd represents the mod command
var depCmd = &cobra.Command{
	Use:   "dep",
	Short: "Manage dependencies",
	Long: `Manages application dependencies and configurations for generating items to build application.
	
"dep init [language]" generates a smart module in the current directory, in effect creating a new application rooted at the current directory.
"dep add --name [item]" adds a new item.
"dep add --name [item] --dep [dependency] --resolver [resolver]" adds a new dependency item to the existing item.
"dep del --name [item]" deletes the item with all dependencies.
"dep del --name [item] --dep [dependency]" deletes item's dependency.
"dep edit --name [item] --dep [dependency] --resolver [resolver]" adds/updates dependency to/in the existing item.
"dep list --name [item]" prints item's dependencies.
"dep list --name [item] --dep [dependency]" prints resolver of dependency item.
"dep list --version" prints module version.
"dep list --lang" prints module language.
"dep list --all" prints module file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check input arguments
		if len(args) == 0 {
			cli.PrintError("Subcommand is required")
			return
		}
		defer cli.Recover()
		var subCmd = args[0]
		var itemStr = strings.Trim(*depCmdFlags.item, "\t \n")
		var depStr = strings.Trim(*depCmdFlags.dep, "\t \n")
		var resolverStr = strings.Trim(*depCmdFlags.resolver, "\t \n")
		// load module
		var c smod.ConfigFile
		switch subCmd {
		case subCmds.add, subCmds.del, subCmds.edit, subCmds.list:
			cli.Check(c.LoadFromFile(app.ModFileName))
		}
		// handle subcommands
		switch subCmd {
		case subCmds.init:
			if len(args) < 2 {
				cli.PrintError("Language parameter is required")
				return
			}
			// create a module file
			if _, err := os.Stat(app.ModFileName); err == nil {
				cli.PrintError(fmt.Sprintf("%s already exists", app.ModFileName))
			} else if !os.IsNotExist(err) {
				cli.PrintError(err)
			} else {
				c = smod.ConfigFile{
					Sb:   app.Version,
					Lang: args[1],
					Items: map[string]map[string]string{
						"main": map[string]string{},
					},
				}
				cli.Check(c.SaveToFile(app.ModFileName))
				fmt.Printf("%s file has been created", app.ModFileName)
			}
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
				cli.Check(c.AddItem(itemStr))
			} else {
				cli.Check(c.AddDependency(itemStr, depStr, resolverStr, false))
			}
		case subCmds.del:
			if itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
			if depStr == "" {
				cli.Check(c.DeleteItem(itemStr))
			} else {
				cli.Check(c.DeleteDependency(itemStr, depStr))
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
			cli.Check(c.AddDependency(itemStr, depStr, resolverStr, true))
		case subCmds.list:
			if depStr != "" && itemStr == "" {
				cli.PrintError("\"--name\" parameter is required")
				return
			}
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
					var item = c.Items[itemStr]
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
		// save the changes into module file
		switch subCmd {
		case subCmds.add, subCmds.del, subCmds.edit:
			cli.Check(c.SaveToFile(app.ModFileName))
		}
	},
}

var depCmdFlags struct {
	item     *string
	dep      *string
	resolver *string
	version  *bool
	lang     *bool
	all      *bool
}

func init() {
	depCmdFlags.item = depCmd.Flags().StringP("name", "n", "", "item name")
	depCmdFlags.dep = depCmd.Flags().StringP("dep", "d", "", "dependency name")
	depCmdFlags.resolver = depCmd.Flags().StringP("resolver", "r", "", "resolver")
	depCmdFlags.version = depCmd.Flags().BoolP("version", "v", false, "print version")
	depCmdFlags.lang = depCmd.Flags().BoolP("lang", "l", false, "print language")
	depCmdFlags.all = depCmd.Flags().BoolP("all", "a", false, "print module")
}
