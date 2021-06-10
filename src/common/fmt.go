// Package common represents common command line methods
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package common

import "fmt"

func PrintError(e interface{}) {
	fmt.Printf(ErrorMessage, e)
}
