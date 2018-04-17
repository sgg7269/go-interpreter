package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// GetExpr ...
func (p *Program) GetExpr() {
	p.GetAssignment()
}

// GetAssignment ...
func (p *Program) GetAssignment() {
	p.GetVar()
	p.GetSeparator()
	p.GetEqualsSign()
	p.GetSeparator()
	p.GetTerm()
	// p.GetSeparator()
	// p.GetEqualsSign()
	// p.GetTerm()
	// p.GetSeparator()
	// p.GetVar()
}

// GetEqualsSign ...
func (p *Program) GetEqualsSign() {
	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			t := token.NewToken(p.GetCurrentIndex())

			if p.GetCurrentChar() == "=" {
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value.String = p.Char.Accumulator
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
