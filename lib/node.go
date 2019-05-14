package gaialisp

const (
	NTSEXPR = iota
	NTID
	NTNUM
	NTLITERAL
	NTLAMBDA
)

type Node struct {
	nodeType int
	sval     string
	ival     float64
	subs     []*Node
}
