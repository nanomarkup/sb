// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove the generated files",
	Long:  `Clean removes the generated files from source directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")
	},
}
