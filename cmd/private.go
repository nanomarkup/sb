// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import "fmt"

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
