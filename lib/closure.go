package gaialisp

import (
	"fmt"
)

type Closure struct {
	scope map[string]*Node
	upper *Closure
}

func NewClosure()*Closure {
	closure:=&Closure{scope: make(map[string]*Node), upper: nil}
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
		// find in upper closure
		if self.upper != nil {
			val = self.upper.GetVar(varName)
			if val == nil {
				exist = false
			}
		}
	}
	
	if !exist {
		panic(fmt.Sprintf("variable [%s] not defined", varName))
	
	}

	// fixme if scope[varName] is nil?
	if val == nil {
		panic(fmt.Sprintf("variable [%s] is nil<internal error>", varName))
	}
	return val
}
