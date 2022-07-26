// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

var starterFlags struct {
	lang string
}

func (r *Starter) init() {
	r.SilenceUsage = true
	r.CompletionOptions.DisableDefaultCmd = true
	r.PersistentFlags().StringVarP(&starterFlags.lang, "lang", "l", "", "select language")
}
