package gaialisp

import (
	"fmt"
)

type Closure struct {
	scope map[string]*Node
}

func NewClosure()*Closure {
	closure:=&Closure{scope: make(map[string]*Node)}
	return closure
}


func (self* Closure) DefVar(varName string, value* Node) {
	_ ,exist:=self.scope[varName]
	if exist {
		panic(fmt.Sprintf("variable [%s] already exists", varName))
	}

	self.scope[varName] = value

}

func (self *Closure) GetVar(varName string) *Node {
	val, exist:=self.scope[varName]

	if !exist {
		panic(fmt.Sprintf("variable [%s] not defined", varName))
	}
	// fixme if scope[varName] is nil?
	return self.scope[varName]
}



