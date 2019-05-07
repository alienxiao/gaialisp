package main

import (
	"fmt"
	gaialisp "nicolas/gaialisp/lib"
)

func main() {
	fmt.Println("Hello, gaialisp")
	code := gaialisp.ReadFile("./test.lisp")

	fmt.Println(code)
}
