// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

// Package cmd represents Command Line Interface.
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/nanomarkup/dl"
	"github.com/nanomarkup/sb"
	"github.com/spf13/cobra"
)

// SmartBuilder includes all available commands and handles them.
type SmartBuilder struct {
	cobra.Command
}

// Creator describes methods for creating an application by generating smart application unit (.sa file).
type Creator interface {
	Create(string) error
}

// Generator describes methods for generating smart builder unit (.sb) using smart application unit.
type Generator interface {
	Generate(string) error
}

// Coder describes methods for generating code to build the application.
type Coder interface {
	Generate(string) error
}

// Builder describes methods for building an application using the generated items.
type Builder interface {
	Build(string) error
}

// Cleaner describes methods for removing generated/compiled files.
type Cleaner interface {
	Clean(string) error
}

// Runner describes methods for running the application.
type Runner interface {
	Run(string) error
}

// AppsPrinter describes methods for displaying all available applications.
type AppsPrinter interface {
	Apps() ([]string, error)
}

// VersionPrinter describes methods for displaying a version of the application.
type VersionPrinter interface {
	Version() string
}

// ModManager describes methods for managing application items and dependencies.
type ModManager interface {
	AddItem(module, item string) error
	AddDependency(item, dependency, resolver string, update bool) error
	DeleteItem(item string) error
	DeleteDependency(item, dependency string) error
	ReadAll(kind string) (ModReader, error)
}

// ModReader describes methods for getting module attributes.
type ModReader interface {
	Items() map[string][][]string
	Dependency(string, string) string
}

// ModFormatter describes methods for formatting module attributes and returns it as a string.
type ModFormatter interface {
	Item(string, [][]string) string
	String(module ModReader) string
}

const (
	// application
	AppsItemName      string = "apps"
	DefaultModuleName string = "apps"
	// error messages
	ErrorMessageF           string = "Error: %v\n"
	AppNameMissing          string = "application name is required"
	SubcmdMissing           string = "subcommand is required"
	ItemMissing             string = "item name is required"
	ModOrDepMissing         string = "module name or dependency name is missing"
	ResolverMissing         string = "resolver is required"
	DependencyMissing       string = "\"--dep\" parameter is required"
	ItemDoesNotExistF       string = "\"%s\" item does not exist"
	DependencyDoesNotExistF string = "\"%s\" dependency item does not exist"
	UnknownSubcmdF          string = "unknown \"%s\" subcommand"
)

func CmdCreate(c *sb.SmartCreator) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Create(args[0])
			} else {
				return fmt.Errorf(AppNameMissing)
			}
		}
	}
}

func CmdGen(c *sb.SmartGenerator) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Generate(args[0])
			} else {
				return c.Generate("")
			}
		}
	}
}

func CmdCode(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Generate(args[0])
			} else {
				return c.Generate("")
			}
		}
	}
}

func CmdBuild(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Build(args[0])
			} else {
				return c.Build("")
			}
		}
	}
}

func CmdClean(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Clean(args[0])
			} else {
				return c.Clean("")
			}
		}
	}
}

func CmdRun(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return c.Run(args[0])
			} else {
				return c.Run("")
			}
		}
	}
}

func CmdList(c *sb.ModHelper) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			apps, err := c.Apps()
			if err != nil {
				return err
			}
			for _, app := range apps {
				fmt.Println(app)
			}
			return nil
		}
	}
}

func CmdVersion(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) {
			fmt.Println(c.Version())
		}
	}
}

func CmdManageMod(c *sb.SmartBuilder, f *dl.Formatter) func(cmd *cobra.Command, args []string) error {
	if c == nil || f == nil {
		return nil
	} else {
		// depFlags.mod = v.Command.Flags().StringP("mod", "m", "", "module name")
		// depFlags.item = v.Command.Flags().StringP("name", "n", "", "item name")
		// depFlags.dep = v.Command.Flags().StringP("dep", "d", "", "dependency name")
		// depFlags.resolver = v.Command.Flags().StringP("resolver", "r", "", "resolver")
		// depFlags.all = v.Command.Flags().BoolP("all", "a", false, "print module")
		return func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New(SubcmdMissing)
			}
			defer handleError()
			var subCmd = args[0]
			var modStr = strings.Trim(*depFlags.mod, "\t \n")
			var itemStr = strings.Trim(*depFlags.item, "\t \n")
			var depStr = strings.Trim(*depFlags.dep, "\t \n")
			var resolverStr = strings.Trim(*depFlags.resolver, "\t \n")
			// handle subcommands
			switch subCmd {
			case subCmds.del:
				// if modStr == "" {
				// 	// return errors.New(ModuleMissing)
				// }
				// if itemStr == "" {
				// 	return errors.New(ItemMissing)
				// }
				// if depStr == "" {
				// 	return v.Manager.DeleteItem(modStr, itemStr)
				// } else {
				// 	return v.Manager.DeleteDependency(modStr, itemStr, depStr)
				// }
			case subCmds.edit:
				if modStr == "" {
					// return errors.New(ModuleMissing)
				}
				if itemStr == "" {
					return errors.New(ItemMissing)
				}
				if depStr == "" {
					return errors.New(DependencyMissing)
				}
				if resolverStr == "" {
					return errors.New(ResolverMissing)
				}
				return c.AddDependency(itemStr, depStr, resolverStr, true)
			case subCmds.list:
				if depStr != "" && itemStr == "" {
					return errors.New(ItemMissing)
				}
				mod, err := c.ReadAll("sb")
				if err != nil {
					return err
				} else if *depFlags.all {
					fmt.Println(f.String(mod))
				} else if itemStr != "" {
					var item = mod.Items()[itemStr]
					if item == nil {
						return fmt.Errorf(ItemDoesNotExistF, itemStr)
					} else {
						if depStr == "" {
							fmt.Print(f.Item(itemStr, item))
						} else {
							found := false
							for _, row := range item {
								if row[0] == depStr {
									found = true
									break
								}
							}
							if found {
								fmt.Print(mod.Dependency(itemStr, depStr))
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
}

func CmdInitMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			defer handleError()
			return c.AddItem(DefaultModuleName, AppsItemName)
		}
	}
}

func CmdAddToMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			defer handleError()
			if len(args) < 1 {
				return errors.New(ItemMissing)
			} else if len(args) < 2 {
				return errors.New(ModOrDepMissing)
			} else if len(args) == 2 {
				return c.AddItem(args[1], args[0])
			} else if len(args) > 2 {
				return c.AddDependency(args[0], args[1], args[2], false)
			} else {
				return nil
			}
		}
	}
}

func CmdDelFromMod(c *sb.SmartBuilder) func(cmd *cobra.Command, args []string) error {
	if c == nil {
		return nil
	} else {
		return func(cmd *cobra.Command, args []string) error {
			defer handleError()
			if len(args) < 1 {
				return errors.New(ItemMissing)
			} else if len(args) == 1 {
				return c.DeleteItem(args[0])
			} else if len(args) == 2 {
				return c.DeleteDependency(args[1], args[0])
			} else {
				return nil
			}
		}
	}
}

func OSStdout() *os.File {
	return os.Stdout
}
