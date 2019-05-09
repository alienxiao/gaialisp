package gaialisp

func Require(vm *VM, fileName string)*Node {
		code := ReadFile(fileName)
		parser := NewParser(code)
		//parser.Test()
		root := parser.Parse()

		return vm.evalNode(root)

}
