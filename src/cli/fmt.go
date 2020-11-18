// Package cli represents common command line methods
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cli

import "fmt"

func PrintError(e interface{}) {
	fmt.Printf(EMessage, e)
}
