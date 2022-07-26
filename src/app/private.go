// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package app

import (
	"fmt"
)

type builder interface {
	Build(app string, sources *map[string]map[string]string) error
	Clean(app string, sources *map[string]map[string]string) error
	Generate(app string, sources *map[string]map[string]string) error
}

var langs = struct {
	Go string
}{
	"go",
}

var suppLangs = map[string]bool{
	langs.Go: true,
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
