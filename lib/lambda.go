package gaialisp

func DefLambda(vm *VM, args []*Node) *Node {
	if len(args) == 2 {
		//first arg is a s-expr with all identifier nodes
		lambdaArgs := args[0]

		if lambdaArgs.nodeType == NTSEXPR {
			for _, lambdaArg := range lambdaArgs.subs {
				if lambdaArg.nodeType == NTID {

				} else {
					panic("lambda args should be identifier")
				}
			}
			return &Node{nodeType: NTLAMBDA, subs: args, upper: vm.currentFrame()}
		} else {
			panic("lambda should be defined with symbol args")
		}

	} else {
		panic("lambda should be defined with args and body")
	}
}

func GetLambdaArgs(lambda *Node) []string {
	//fixme assume lambda is ok
	args := lambda.subs
	lambdaArgs := args[0]

	argNames:=make([]string, 0)

	if lambdaArgs.nodeType == NTSEXPR {
		for _, lambdaArg := range lambdaArgs.subs {
			if lambdaArg.nodeType == NTID {
				argNames = append(argNames, lambdaArg.sval)
			} else {
				panic("lambda args should be identifier")
			}
		}
		return argNames
	} else {
		panic("lambda should be defined with symbol args")
	}
}

func GetLambdaBody(lambda *Node)*Node {
	args:=lambda.subs
	if len(args) == 2 {
		return args[1]
	} else {
		panic("lambda must be defined with args and body")
	}
}