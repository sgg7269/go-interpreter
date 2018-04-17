package main

import (
	"fmt"
	"io/ioutil"
	"os"

	parser "github.com/sgg7269/go-interpreter/parser"
)

// TODO: should make a map of this shit
var (
	EOSToken       = ";"
	SeparatorToken = " "
	TokenIDList    = map[string]string{
		"EOS": "META",
		"INT": "TYPE",
	}
	tokenList = map[string]Token{
		"int": Token{
			Type:  TokenIDList["INT"],
			Value: "int",
		},
		// "string": 2,
	}
)

func main() {
	programName := "program.expr"

	input, err := ioutil.ReadFile(programName)
	if err != nil {
		fmt.Println("error reading input program", programName)
		os.Exit(1)
	}

	// // Make sure we have an eos (;)
	// if input[len(input)-1] != ';' {
	// 	fmt.Println("You must have an end of statement (;)")
	// 	return
	// }

	p := parser.Program{
		Value:  string(input),
		Length: len(input),
		Char: Char{
			CurrentChar: string(input[0]),
		},
	}

	fmt.Printf("Program start: %#v\n\n", p)

	p.GetProgram()

	fmt.Printf("Program end: %#v\n\n", p)

	fmt.Println("Tokens:")
	for _, token := range p.Tokens {
		fmt.Printf("Program end: %#v\n", token)
	}
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

// SetEnd ...
func (t *Token) SetEnd(index int) {
	t.Location[1] = index
	return
}

// NewToken ...
func NewToken(index int) Token {
	return Token{Location: [2]int{index, -1}}
}

// AddToken ...
func (p *Program) AddToken(t Token) {
	p.Tokens = append(p.Tokens, t)
	return
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
