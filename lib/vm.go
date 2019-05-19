package gaialisp

import (
	"fmt"
)

type VM struct {
	root   *Node
	frames []*Closure
}

func NewVM(root *Node) *VM {
	vm := &VM{}
	vm.root = root
	vm.frames = make([]*Closure, 0)
	globalScope := NewClosure()
	vm.pushFrame(globalScope)
	return vm

}

func (self *VM) pushFrame(c *Closure) {
	self.frames = append(self.frames, c)
}
func (self *VM) popFrame() {
	if len(self.frames) <= 1 {
		panic("popframe error")
	}
	self.frames = self.frames[0 : len(self.frames)-1]
}
func (self *VM) currentFrame() *Closure {
	if len(self.frames) <= 0 {
		panic("frame is empty")
	}

	return self.frames[len(self.frames)-1]
}

func (self *VM) Eval() {
	self.evalNode(self.root)
}

func (self *VM) evalNode(node *Node) *Node {
	//fmt.Printf("eval node: %d sval %s ival %f\n", node.nodeType, node.sval, node.ival)
	switch node.nodeType {
	case NTSEXPR:
		return self.evalSExpr(node)
	case NTNUM:
		return node
	case NTLITERAL:
		return node
	case NTLAMBDA:
		return node
	case NTID:
		//fixme should get local var
		return self.GetVar(node.sval)
	default:
		panic(fmt.Sprintf("vm error: unknown nodeType %d", node.nodeType))
	}
}

func (self *VM) evalSExpr(node *Node) *Node {
	if node.nodeType != NTSEXPR {
		panic("vm error: eval s-expr is not s-expr")
	} else {
		subs := node.subs
		if len(subs) <= 0 {
			panic("vm error: not support empty s-expr yet")
		}
		firstNode := subs[0]
		// eval first node #begin
		if firstNode.nodeType == NTID {
			// todo support variable
			// current only support internal functions
			return self.callInternalFunction(firstNode.sval, subs[1:])
		} else {
			// todo support normal node eval
			panic("only identifier is support in s-expr's first node currently")
		}

		// eval first node #end
	}
}

func (self *VM) GetVar(varName string) *Node {
	currentFrame := self.currentFrame()
	return currentFrame.GetVar(varName)
}

func (self *VM) DefVar(varName string, value *Node) *Node {
	//fixme return nil node
	ret := &Node{nodeType: NTNUM, ival: 0}

	currentFrame := self.currentFrame()
	currentFrame.DefVar(varName, value)

	return ret
}

func (self *VM) CallLambda(lambda *Node, args []*Node) *Node {
	argNames := GetLambdaArgs(lambda)
	body:=GetLambdaBody(lambda)
	//number of args must be same
	if len(argNames) == len(args) {
		//bind args
		frame := NewClosure()
		frame.upper = lambda.upper
		self.pushFrame(frame)

		for x, argName := range argNames {
			frame.DefVar(argName, self.evalNode(args[x]))
		}

		ret:= self.evalNode(body)
		self.popFrame()
		return ret

	} else {
		panic("number of args must be same")
	}

}

func (self *VM) callInternalFunction(functionName string, args []*Node) *Node {
	if functionName == "print" {
		return Buildins__print(self, args)
	} else if functionName == "+" {
		return Buildins__add(self, args)
	} else if functionName == "-" {
		return Buildins__sub(self, args)
	} else if functionName == "*" {
		return Buildins__mul(self, args)
	} else if functionName == "/" {
		return Buildins__div(self, args)
	} else if functionName == "sqrt" {
		return Buildins__sqrt(self, args)
	} else if functionName == "progn" {
		//sequence execution
		node := &Node{nodeType: NTNUM, ival: 0}

		if len(args) > 0 {
			//fixme default return node should be nil node
			for _, arg := range args {
				node = self.evalNode(arg)
			}
		} else {
			fmt.Println("progn requires at least 1 arg")
		}
		return node

	} else if functionName == "defvar" {
		if len(args) == 2 {
			//first arg is id
			if args[0].nodeType == NTID {

				value := self.evalNode(args[1])
				return self.DefVar(args[0].sval, value)
			} else {
				panic("define first arg should be an identifier")
			}
		} else {
			panic("define syntax error: requires 2 args")
		}
	} else if functionName == "import" {
		if len(args) == 1 && args[0].nodeType == NTLITERAL {
			return Require(self, args[0].sval)
		} else {
			panic("require syntax error: requires 1 string arg")
		}
	} else if functionName == "lambda" {
		return DefLambda(self, args)
	} else {
		// find node by identifier

		firstNode := self.GetVar(functionName)

		// only if firstNode is lambda can be called
		// todo this logic should be outside

		if firstNode.nodeType == NTLAMBDA {
			return self.CallLambda(firstNode, args)
		} else {
			panic(fmt.Sprintf("unknown function name: %s", functionName))
		}
		
	}

}
