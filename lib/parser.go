package gaialisp

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(source string) *Parser {
	parser := &Parser{}
	parser.lexer = NewLexer(source)
	parser.lexer.NextToken()

	return parser

}

func (self *Parser) Parse() *Node {
	if self.lexer.tokenType == TTLPAR {
		return self.parseSExpr()

	} else {
		return self.parseFactor()
	}
}

func (self *Parser) parseSExpr() *Node {
	if self.lexer.tokenType != TTLPAR {
		panic("parse error s-expr require (")
	} else {
		node := &Node{nodeType: NTSEXPR, subs: make([]*Node, 0)}
		self.lexer.NextToken()
		// parse first element
		for self.lexer.tokenType != TTRPAR && self.lexer.tokenType != TTEND {
			subNode := self.Parse()
			node.subs = append(node.subs, subNode)
		}

		if self.lexer.tokenType != TTRPAR {
			panic("parse error s-expr require )")
		} else {
			self.lexer.NextToken()
		}
		return node
	}
}

func (self *Parser) parseFactor() *Node {
	switch self.lexer.tokenType {
	case TTID:
		node := &Node{nodeType: NTID, sval: self.lexer.token}
		self.lexer.NextToken()
		return node
	case TTNUM:
		var num float64 = 0
		var err error
		num, err = strconv.ParseFloat(self.lexer.token, 64)
		if err != nil {
			panic(err.Error())
		}
		node := &Node{nodeType: NTNUM, ival: num}
		self.lexer.NextToken()
		return node
	case TTLITERAL:
		node := &Node{nodeType: NTLITERAL, sval: self.lexer.token}
		self.lexer.NextToken()
		return node
	default:
		panic(fmt.Sprintf("unknown tokenType: %d", self.lexer.tokenType))
	}
}

func (self *Parser) Test() {
	self.Parse()
}
