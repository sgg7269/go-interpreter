package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// Program is the Overarching program type that holds information about the current program parse
type Program struct {
	Value     string
	Length    int
	Index     int
	EOF       bool
	Char      Char
	Statement Statement
	Tokens    []token.Token
}

// GetProgram ...
func (p *Program) GetProgram() {
	// p.GetStatement()
	fmt.Println(p.GetNextToken())
}
