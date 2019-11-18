package main

import (
	"fmt"

	"github.com/devilandpuppy/mofalisp/mofa"
)

func main() {
	fmt.Println("hello, mofalisp")

	code:=mofa.LoadFile("./examples/base.lisp")

	parser:=mofa.NewParser(code)

	parser.Parse()
}
