package main

import (
	"fmt"

	"github.com/devilandpuppy/mofalisp/mofa"
)

func main() {
	fmt.Println("mofalisp v0.0.1 by devilandpuppy")
	fmt.Println()
	
	code:=mofa.LoadFile("./examples/base.lisp")

	parser:=mofa.NewParser(code)

	node:=parser.Parse()

	mofa.Eval(node)
}
