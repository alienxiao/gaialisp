package gaialisp

func DefLambda(vm * VM, args []* Node) *Node {
	if len(args) == 2 {
		//first arg is a s-expr with all identifier nodes
		lambdaArgs:= args[0]
		
		if lambdaArgs.nodeType == NTSEXPR {
			for _, lambdaArg:=range lambdaArgs.subs {
				if lambdaArg.nodeType == NTID {
					
				} else {
					panic("lambda args should be identifier")
				}
			}
			return &Node{nodeType: NTLAMBDA, subs: args}
		} else {
			panic("lambda should be defined with symbol args")
		}

	} else {
		panic("lambda should be defined with args and body")
	}
}
