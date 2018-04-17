package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

// Program is the Overarching program type that holds information about the current program parse
type (
	Program struct {
		Value     string
		Length    int
		Index     int
		EOF       bool
		Char      Char
		Statement Statement
		Tokens    []Token
	}

	// Char holds information about the current char that is being parsed
	Char struct {
		Index       int
		Accumulator string
		CurrentChar string
		LastChar    string
		EOS         bool
	}

	// Statement holds information about the current statement that is being parsed
	Statement struct {
		Index  int
		Value  string
		Length int
	}

	// Token is a syntactical structure used to hold representative metadata
	Token struct {
		Location [2]int
		Type     string
		Value    string
	}
)

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

	p := Program{
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

// GetSeparator ...
// TODO: fix this
func (p *Program) GetSeparator() {
	for {
		if !p.EOF && !p.Char.EOS {
			if p.Char.CurrentChar == SeparatorToken {
				t := NewToken(p.GetCurrentIndex())
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value = p.Char.Accumulator
				t.Type = "SEP"
				// p.AddToken(t)
				p.Char.Accumulator = ""
				return
			}
		} else {
			return
		}
	}
}

// GetProgram ...
func (p *Program) GetProgram() { p.GetStatement() }

// GetStatement ...
func (p *Program) GetStatement() {
	p.GetExpr()
	p.GetEOS()
}

// GetExpr ...
func (p *Program) GetExpr() { p.GetEqualsExpr() }

// GetEqualsExpr ...
func (p *Program) GetEqualsExpr() {
	p.GetVar()
	p.GetSeparator()
	p.GetEquals()
	p.GetSeparator()
	p.GetTerm()
}

// GetVar ...
func (p *Program) GetVar() {
	p.GetType()
	p.GetSeparator()
	p.GetIdent()
}

// GetEquals ...
func (p *Program) GetEquals() {
	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			t := NewToken(p.GetCurrentIndex())

			if p.GetCurrentChar() == "=" {
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value = p.Char.Accumulator
				t.Type = "EQUALS"
				p.AddToken(t)
				p.Char.Accumulator = ""
				return
			}

			fmt.Println("vert da ferk 3")
			return

		}
	}
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

// GetType gets the next type from the sequence
func (p *Program) GetType() {
	t := NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			p.GetNextChar()
			if token, ok := tokenList[p.Char.Accumulator]; ok {
				t.SetEnd(p.GetLastIndex())
				t.Value = token.Value
				t.Type = token.Type
				p.AddToken(t)
				return
			}
		} else {
			fmt.Println("vert a ferk")
			return
		}
	}
}

// GetIdent ...
func (p *Program) GetIdent() {
	t := NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			p.GetNextChar()
		} else {
			t.SetEnd(p.GetLastIndex())
			t.Value = p.Char.Accumulator
			t.Type = "IDENT"
			p.AddToken(t)
			p.Char.Accumulator = ""
			return
		}
	}
}

// GetTerm ...
func (p *Program) GetTerm() { p.GetConstant() }

// GetConstant ...
func (p *Program) GetConstant() {
	t := NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			p.GetNextChar()
		} else {
			_, err := strconv.Atoi(p.Char.Accumulator)
			if err != nil {
				fmt.Println("vert da ferk 4", err)
				return
			}

			t.SetEnd(p.GetLastIndex())
			t.Value = p.Char.Accumulator
			t.Type = "RAW_INT"
			p.AddToken(t)
			p.Char.Accumulator = ""
			return
		}
	}
}

// GetEOS ...
func (p *Program) GetEOS() {
	for {
		if !p.EOF {
			if p.Char.EOS {
				t := NewToken(p.GetCurrentIndex())
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value = p.Char.Accumulator
				t.Type = "EOS"
				p.AddToken(t)
				p.Char.Accumulator = ""
				return
			}
		} else {
			fmt.Println("vert da ferk 5")
			return
		}
	}
}
