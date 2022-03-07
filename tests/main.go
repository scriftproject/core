package main

import (
	"github.com/scriftproject/core"
)

func main() {
	src := core.ReadSource("./examples/basic_01.scr")
	src.Lex()
}
