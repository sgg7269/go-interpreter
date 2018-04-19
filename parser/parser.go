package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// TODO: should make a map of this shit
var (
	EOSToken       = ";"
	SeparatorToken = " "
	// TokenIDList    = map[string]string{
	// 	"eof": "META",
	// 	"eos": "META",
	// 	"var": "TYPE",
	// }
	tokenList = map[string]token.Token{
		"ident": token.Token{
			ID:   -3,
			Type: "IDENT",
		},
		"eos": token.Token{
			ID:   -2,
			Type: "META",
			Value: token.Value{
				String: "EOS",
			},
		},
		"eof": token.Token{
			ID:   -1,
			Type: "META",
			Value: token.Value{
				String: "EOF",
			},
		},
		"var": token.Token{
			ID:   1,
			Type: "TYPE",
			Value: token.Value{
				String: "var",
			},
		},
		"=": token.Token{
			ID:   2,
			Type: "EQUALS",
			Value: token.Value{
				String: "=",
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

// PeekNextChar ...
func (p *Program) PeekNextChar() string {
	if p.Index+1 < p.Length {
		return string(p.Value[p.Index+1])
	}
	return ""
}

// GetNextChar returns the next character in the sequence by incrementing the index and returning the "current" char
func (p *Program) GetNextChar() {
	if p.Index < p.Length-1 {
		p.Char.Accumulator += p.Char.CurrentChar
		p.Char.LastChar = p.Char.CurrentChar
		p.Index++
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
		return
	}
}

// // GetCurrentChar returns the character at the current index
// func (p *Program) GetCurrentChar() (string, error) {
// 	if p.Index < p.Length {
// 		return string(p.Value[p.Index]), nil
// 	}

// 	// FIXME: maybe use this for EOF?
// 	return "", errors.New("EOF")
// }

// GetCurrentChar returns the character at the current index
func (p *Program) GetCurrentChar() string {
	if p.Index < p.Length {
		return string(p.Value[p.Index])
	}

	p.EOF = true
	return ""
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
	p.CollectedTokens = append(p.CollectedTokens, t)
	return
}

// GetNextToken ...
// FIXME: we can start out with going until space, but we will need something for strings after that
// TODO: should actually return the token from a map lookup and switch on it to make sure its what were expecting
// FIXME: somewhere in the program struct should be a "expectedNextToken" struct, "lastToken" struct, etc
// TODO: could we use a select here?
func (p *Program) GetNextToken() token.Token {
	var ok bool
	start := p.GetCurrentIndex()

	for {

		peek := p.PeekNextChar()
		if p.PeekNextChar() == ";" {
			p.GetNextChar()
			// Parse the current accumulator
			var t token.Token
			if t, ok = tokenList[p.Char.Accumulator]; ok {
				t.SetLocation(start, p.GetLastIndex())
			} else {
				t = tokenList["ident"]
				t.SetLocation(start, p.GetLastIndex())
				t.Value.String = p.Char.Accumulator
			}
			p.AddToken(t)
			// p.Char.Accumulator = ""
			return t
		}

		p.GetNextChar()
		fmt.Println("peeked", peek)

		switch {
		case p.EOF:
			t := tokenList["eof"]
			t.SetLocation(start, p.GetCurrentIndex())
			p.AddToken(t)
			return t

		case p.GetCurrentChar() == ";":
			// // Parse the current accumulator
			// var t token.Token
			// if t, ok = tokenList[p.Char.Accumulator]; ok {
			// 	t.SetLocation(start, p.GetLastIndex())
			// } else {
			// 	t = tokenList["ident"]
			// 	t.SetLocation(start, p.GetLastIndex())
			// 	t.Value.String = p.Char.Accumulator
			// }
			// p.AddToken(t)

			// Printout an end-of-statement token
			tt := tokenList["eos"]
			tt.SetLocation(p.GetCurrentIndex(), p.GetCurrentIndex())
			p.AddToken(tt)

			p.GetNextChar()
			p.Char.Accumulator = ""

			return tt

		case p.GetCurrentChar() == " ":
			var t token.Token
			if t, ok = tokenList[p.Char.Accumulator]; ok {
				t.SetLocation(start, p.GetLastIndex())
			} else {
				t = tokenList["ident"]
				t.SetLocation(start, p.GetLastIndex())
				t.Value.String = p.Char.Accumulator
				fmt.Println("i am here")
			}
			p.AddToken(t)

			p.GetNextChar()
			p.Char.Accumulator = ""
			return t

		default:
			// fmt.Println(p.Char.Accumulator)
		}
	}
}
