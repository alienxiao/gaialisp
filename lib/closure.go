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
	// fmt.Printf(">>>>get var [%s]\n", varName)
	val, exist:=self.scope[varName]

	if !exist {
		// fmt.Printf(">>>>var not exist in current [%s]\n", varName)
		// find in upper closure
		if self.upper != nil {
			// fmt.Printf(">>>>find var in upper [%s]\n", varName)
			val = self.upper.GetVar(varName)
			exist = val != nil
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
