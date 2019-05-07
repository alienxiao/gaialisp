package gaialisp

const (
	NTSEXPR = iota
	NTID
	NTNUM
	NTLITERAL
)

type Node struct {
	nodeType int
	sval     string
	ival     float64
	subs     []*Node
}
