// Package cli represents common command line methods
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package cli

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
