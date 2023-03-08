// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"os"
)

// Execute initializes commands and displays them in the console.
func (sb *SmartBuilder) Execute() error {
	sb.Starter.init()
	err := sb.Starter.Execute()
	if (err != nil) && !sb.SilentErrors {
		os.Exit(1)
	}
	return err
}
