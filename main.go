package main

import (
	"fmt"

	"github.com/devilandpuppy/mofalisp/mofa"
)

func main() {
	fmt.Println("hello, mofalisp")

	parser:=mofa.NewParser("")

	parser.Parse()
}
