package main

import tk "github.com/nsf/gothic"
import "fmt"

func main() {
	ir := tk.NewInterpreter(`
namespace eval go {}
ttk::entry .e -textvariable go::etext
trace add variable go::etext write go::onchange
pack .e -fill x -expand true
	`)

	ir.RegisterCommand("go::onchange", func() {
		fmt.Println(ir.EvalAsString("set go::etext"))
	})
	<-ir.Done
}
