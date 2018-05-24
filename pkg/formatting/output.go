package formatting

import (
	"github.com/fatih/color"
)

var (
	Err  *color.Color
	Warn *color.Color
	Info *color.Color
)

func init() {
	Err = color.New(color.FgRed)
	Warn = color.New(color.FgYellow)
	Info = color.New(color.FgGreen)
}
