package formatting

import (
	"io"

	"github.com/fatih/color"
)

type WrapColor struct {
	innerColor *color.Color
}

var (
	Err  *WrapColor
	Warn *WrapColor
	Info *WrapColor
)

func init() {
	Err = NewColor(color.FgRed)
	Warn = NewColor(color.FgYellow)
	Info = NewColor(color.FgGreen)
}

func (wc *WrapColor) Fprintln(w io.Writer, a ...interface{}) {
	_, err := wc.innerColor.Fprintln(w, a)
	if err != nil {
		//skip
	}
}

func (wc *WrapColor) Fprintf(w io.Writer, format string, a ...interface{}) {
	_, err := wc.innerColor.Fprintf(w, format, a)
	if err != nil {
		//skip
	}
}

func (wc *WrapColor) Fprint(w io.Writer, a ...interface{}) {
	_, err := wc.innerColor.Fprint(w, a)
	if err != nil {
		//skip
	}
}

func NewColor(attr color.Attribute) *WrapColor {
	return &WrapColor{
		innerColor: color.New(attr),
	}
}
