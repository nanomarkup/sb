package app

import "fmt"

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
