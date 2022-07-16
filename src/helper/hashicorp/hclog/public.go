package helper

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
)

var LogLever = struct {
	NoLevel uint
	Trace   uint
	Debug   uint
	Info    uint
	Warn    uint
	Error   uint
	Off     uint
}{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

// integer is not supported yet
//func New(name string, level uint) hclog.Logger {
func NewStdOut(name string) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Level:  hclog.Level(LogLever.Info),
		Output: os.Stdout,
	})
}

func NewFileOut(name string) hclog.Logger {
	f, err := os.Create(fmt.Sprintf("%s.log", name))
	if err != nil {
		panic(err)
	}
	return hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Level:  hclog.Level(LogLever.Info),
		Output: f,
	})
}
