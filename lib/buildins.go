package gaialisp

import (
	"fmt"
	"math"
)

func Buildins__print(self *VM, args []*Node) *Node {

	// todo support more args
	if len(args) > 0 {
		argOrg := args[0]
		arg := self.evalNode(argOrg)
		if arg.nodeType == NTNUM {
			fmt.Printf("%f\n", arg.ival)
		} else if arg.nodeType == NTLITERAL {
			fmt.Printf("%s\n", arg.sval)
		} else {
			fmt.Println("print only can print number or literal")
		}
	} else {
		fmt.Println("print requires at least 1 arg")
	}
	//fixme return a nil node instead
	return &Node{nodeType: NTNUM, ival: 0}
}

func Buildins__add(self *VM, args []*Node) *Node {
	var sumNum float64 = 0
	if len(args) > 0 {
		for _, arg := range args {
			node := self.evalNode(arg)
			if node.nodeType == NTNUM {
				sumNum += node.ival
			} else {
				fmt.Println("sum only support number")
				break
			}
		}

	} else {
		fmt.Println("sum requires at least 1 arg")
	}
	return &Node{nodeType: NTNUM, ival: sumNum}

}

func Buildins__sub(self *VM, args []*Node) *Node {
	var sumNum float64 = 0
	first := true
	if len(args) >= 2 {
		for _, arg := range args {
			node := self.evalNode(arg)
			if node.nodeType == NTNUM {
				if first {
					sumNum = node.ival
					first = false
				} else {
					sumNum -= node.ival
				}
			} else {
				fmt.Println("sub only support number")
				break
			}
		}

	} else {
		fmt.Println("sub requires at least 2 args")
	}
	return &Node{nodeType: NTNUM, ival: sumNum}

}

func Buildins__mul(self *VM, args []*Node) *Node {
	var sumNum float64 = 1
	if len(args) > 0 {
		for _, arg := range args {
			node := self.evalNode(arg)
			if node.nodeType == NTNUM {
				sumNum *= node.ival
			} else {
				fmt.Println("mul only support number")
				break
			}
		}

	} else {
		fmt.Println("mul requires at least 1 arg")
	}
	return &Node{nodeType: NTNUM, ival: sumNum}

}

func Buildins__div(self *VM, args []*Node) *Node {
	var sumNum float64 = 1
	first := true
	if len(args) >= 2 {
		for _, arg := range args {
			node := self.evalNode(arg)
			if node.nodeType == NTNUM {
				if first {
					sumNum = node.ival
					first = false
				} else {
					sumNum /= node.ival

				}
			} else {
				fmt.Println("div only support number")
				break
			}
		}

	} else {
		fmt.Println("div requires at least 2 args")
	}
	return &Node{nodeType: NTNUM, ival: sumNum}

}



func Buildins__sqrt(self *VM, args []*Node) *Node {
	var result float64 = 0
	if len(args) == 1 {
			node := self.evalNode(args[0])
			if node.nodeType == NTNUM {
				result = math.Sqrt(node.ival)
			} else {
				fmt.Println("sqrt only support number")
			}

	} else {
		fmt.Println("sqrt requires 1 arg")
	}
	return &Node{nodeType: NTNUM, ival: result}

}
