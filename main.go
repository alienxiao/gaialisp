package main

import (
	"fmt"
	gaialisp "nicolas/gaialisp/lib"
	"os"
)

func main() {
	fmt.Println("gaialisp v0.0.1")
	fmt.Println("developed by Nicolas Siu")
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		code := gaialisp.ReadFile(fileName)
		parser := gaialisp.NewParser(code)
		//parser.Test()
		root := parser.Parse()

		vm := gaialisp.NewVM(root)

		vm.Eval()
	}
}
