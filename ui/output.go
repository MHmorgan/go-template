package ui

import (
	"fmt"
	"github.com/mbndr/figlet4go"
)

var figlet *figlet4go.AsciiRender

func init() {
	figlet = figlet4go.NewAsciiRender()
}

func Print(obj any, opts ...OutputOption) {
	fmt.Println(obj)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Figlet(s string, opts ...OutputOption) {
	str, err := figlet.Render(s)
	if err != nil {
		fmt.Println(s)
	} else {
		fmt.Println(str)
	}
}

func Error(format string, a ...any) {
	if len(a) > 0 {
		format = fmt.Sprintf(format, a...)
	}
	fmt.Println("Error: ", format)
}

type outputOptions struct{}

type OutputOption func(*outputOptions)
