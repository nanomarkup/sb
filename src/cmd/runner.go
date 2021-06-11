// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

import (
	"github.com/spf13/cobra"
)

type Runner struct {
	cobra.Command
}

var runnerFlags struct {
	lang string
}

func Language() string {
	return runnerFlags.lang
}

func (r *Runner) init() {
	r.SilenceUsage = true
	r.PersistentFlags().StringVarP(&runnerFlags.lang, "lang", "l", "", "select language")
}
