// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

var starterFlags struct {
	lang string
}

func (r *Starter) init() {
	r.SilenceUsage = true
	r.CompletionOptions.DisableDefaultCmd = true
	r.PersistentFlags().StringVarP(&starterFlags.lang, "lang", "l", "", "select language")
}
