// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package app

import (
	"fmt"
)

const (
	coderAttrName string = "coder"
)

type builder interface {
	Build(app string) error
	Clean(app string, sources *map[string]map[string]string) error
	Generate(app string, sources *map[string]map[string]string) error
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}