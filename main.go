package main

import (
	"fmt"
	gaialisp "nicolas/gaialisp/lib"
)

func main() {
	fmt.Println("gaialisp v0.0.1")
	fmt.Println("developed by Nicolas Siu")
	code := gaialisp.ReadFile("./test.lisp")
	parser := gaialisp.NewParser(code)
	//parser.Test()
	root := parser.Parse()

	vm := gaialisp.NewVM(root)

	vm.Eval()
}
