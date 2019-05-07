package gaialisp

import (
	"fmt"
)

type Node struct {
	nodeType int
	sval     string
	ival     float64
}

type Parser struct {
	lexer *Lexer
}

func NewParser(source string) *Parser {
	parser := &Parser{}
	parser.lexer = NewLexer(source)
	parser.lexer.NextToken()

	return parser

}

func (self *Parser) Parse() {
	if self.lexer.tokenType == TTLPAR {
		self.parseSExpr()

	} else {
		self.parseFactor()
	}
}

func (self *Parser) parseSExpr() {
	if self.lexer.tokenType != TTLPAR {
		panic("parse error s-expr require (")
	} else {
		self.lexer.NextToken()
		// parse first element
		for self.lexer.tokenType != TTRPAR && self.lexer.tokenType != TTEND {
			self.Parse()
		}

		if self.lexer.tokenType != TTRPAR {
			panic("parse error s-expr require )")
		} else {
			self.lexer.NextToken()
		}
	}
}

func (self *Parser) parseFactor() {
	switch self.lexer.tokenType {
	case TTID:
		self.lexer.NextToken()
	case TTNUM:
		self.lexer.NextToken()
	case TTLITERAL:
		self.lexer.NextToken()
	default:
		panic(fmt.Sprintf("unknown tokenType: %d", self.lexer.tokenType))
	}
}

func (self * Parser) Test () {
  self.Parse()
}
