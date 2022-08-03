// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

func (r *Starter) init() {
	r.SilenceUsage = true
	r.CompletionOptions.DisableDefaultCmd = true
}
