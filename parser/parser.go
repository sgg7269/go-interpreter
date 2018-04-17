package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// TODO: should make a map of this shit
var (
	EOSToken       = ";"
	SeparatorToken = " "
	TokenIDList    = map[string]string{
		"EOS": "META",
		"VAR": "TYPE",
	}
	tokenList = map[string]token.Token{
		"var": token.Token{
			Type: TokenIDList["VAR"],
			Value: token.Value{
				True:   "<var_id>",
				String: "var",
			},
		},
		// "string": 2,
	}
)

// Char holds information about the current char that is being parsed
type Char struct {
	Index       int
	Accumulator string
	CurrentChar string
	LastChar    string
	EOS         bool
}

// func LoadNextStatement() {}

// GetNextChar returns the next character in the sequence by incrementing the index and returning the "current" char
func (p *Program) GetNextChar() {
	if p.Index < p.Length-1 {
		p.Index++
		p.Char.Accumulator += p.Char.CurrentChar
		p.Char.LastChar = p.Char.CurrentChar
		p.Char.CurrentChar = p.GetCurrentChar()

		// FIXME: idk
		if p.Char.CurrentChar == EOSToken {
			p.Char.EOS = true
		}
	} else if p.Index == p.Length-1 {
		p.Char.Accumulator += p.Char.CurrentChar
		p.Char.LastChar = p.Char.CurrentChar
		p.Char.CurrentChar = p.GetCurrentChar()

		// TODO: fix this hack later
		p.Index++
	} else {
		fmt.Println("vert da ferk 6")
		return
	}
}

// GetCurrentChar returns the character at the current index
func (p *Program) GetCurrentChar() string {
	return string(p.Value[p.Index])
}

// GetCurrentIndex ...
func (p *Program) GetCurrentIndex() int {
	return p.Index
}

// GetLastIndex ...
func (p *Program) GetLastIndex() int {
	// TODO: fix this hack later
	if p.Index == p.Length {
		return p.Length - 1
	}
	return p.Index - 1
}

// AddToken ...
func (p *Program) AddToken(t token.Token) {
	p.Tokens = append(p.Tokens, t)
	return
}
