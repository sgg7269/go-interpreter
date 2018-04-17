package program

import (
	"fmt"
	"strconv"
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
		Value    interface{}
	}
)

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
			accumulatorInt, err := strconv.Atoi(p.Char.Accumulator)
			if err != nil {
				fmt.Println("vert da ferk 4", err)
				return
			}

			t.SetEnd(p.GetLastIndex())
			t.Value = accumulatorInt
			t.Type = "INT_LIT"
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
