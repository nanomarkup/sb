package cmd

import "fmt"

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
