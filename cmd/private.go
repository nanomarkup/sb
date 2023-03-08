// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

func init() {
	cobra.EnableCommandSorting = false
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
