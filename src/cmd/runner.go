// Package cmd represents Command Line Interface
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cmd

var runnerFlags struct {
	lang string
}

func (r *Runner) init() {
	r.SilenceUsage = true
	r.PersistentFlags().StringVarP(&runnerFlags.lang, "lang", "l", "", "select language")
}
