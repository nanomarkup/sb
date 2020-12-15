// Package app represents application items
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/smod"
)

func CheckConfiguration(configuration string, config *smod.ConfigFile) error {
	// check version
	if _, found := versions[config.Sb]; !found {
		return fmt.Errorf("The current \"%s\" version is not supported", config.Sb)
	}
	// check language
	if _, found := suppLangs[config.Lang]; !found {
		return fmt.Errorf("The current \"%s\" language is not supported", config.Lang)
	}
	// read the main item
	var main = config.Items["main"]
	if main == nil {
		return fmt.Errorf("The main item is not found")
	}
	// check the number of existing configurations
	if len(main) == 0 {
		return fmt.Errorf("Does not found any configuration in the main")
	}
	// check the configuration is exist
	if _, found := main[configuration]; !found && configuration != "" {
		return fmt.Errorf("The selected \"%s\" configuration is not found", configuration)
	}
	return nil
}
