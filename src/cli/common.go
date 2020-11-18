// Package cli represents common command line methods
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cli

import (
	"io"
	"os"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Recover() {
	if r := recover(); r != nil {
		PrintError(r)
	}
}

func IsDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, nil
}
