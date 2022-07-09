package helper

import (
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
func New(name string) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Output: os.Stdout,
		Level:  hclog.Level(LogLever.Info),
	})
}
