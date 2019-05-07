package gaialisp

import (
	"fmt"
)

const (
	TTLPAR = iota
	TTRPAR
	TTID
	TTNUM
	TTLITERAL
	TTSTART
	TTEND
)

func isSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func isAlpha(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '+' || c == '-' || c == '*' || c == '/' || c == '_'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAlphanum(c rune) bool {
	return isAlpha(c) || isNumber(c)
}

type Lexer struct {
	source    []rune
	sourceLen int
	pc        int
	cc        rune
	token     string
	tokenType int
}

func NewLexer(source string) *Lexer {
	lexer := &Lexer{}
	lexer.source = []rune(source)
	lexer.sourceLen = len(lexer.source)
	lexer.pc = 0
	lexer.cc = 0
	lexer.token = ""
	lexer.tokenType = TTSTART
	lexer.getchar()
	return lexer

}

func (self *Lexer) getchar() {
	if self.pc < self.sourceLen {
		self.cc = self.source[self.pc]
		self.pc++
	} else {
		self.cc = 0
	}
}

func (self *Lexer) NextToken() {
	for isSpace(self.cc) {
		self.getchar()
	}

	if self.cc == '(' {
		self.token = "("
		self.tokenType = TTLPAR
		self.getchar()
	} else if self.cc == ')' {
		self.token = ")"
		self.tokenType = TTRPAR
		self.getchar()
	} else if isAlpha(self.cc) {
		self.token = string(self.cc)

		self.getchar()
		for isAlphanum(self.cc) {
			self.token += string(self.cc)
			self.getchar()
		}

		self.tokenType = TTID

	} else if isNumber(self.cc) {
		//todo number with dot not support
		self.token = string(self.cc)
		self.getchar()

		for isNumber(self.cc) {
			self.token += string(self.cc)
			self.getchar()
		}

		self.tokenType = TTNUM

	} else if self.cc == '"' {
		// string literal
		self.token = ""

		self.getchar()

		for self.cc != '"' && self.cc != '\r' && self.cc != '\n' && self.cc != 0 {
			self.token += string(self.cc)
			self.getchar()

		}

		if self.cc != '"' {
			panic("string lack quote")
		} else {
			self.getchar()
		}

	} else if self.cc == 0 {
		self.token = ""
		self.tokenType = TTEND

	} else {
		panic(fmt.Sprintf("tokenize error: unknown character: %c", self.cc))
	}
}

func (self *Lexer) Test() {
	self.NextToken()

	for self.tokenType != TTEND {
		fmt.Println(self.token)
		self.NextToken()
	}
}
