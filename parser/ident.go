package parser

import "github.com/sgg7269/go-interpreter/token"

// GetIdent ...
func (p *Program) GetIdent() {
	t := token.NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			p.GetNextChar()
		} else {
			t.SetEnd(p.GetLastIndex())
			t.Value.String = p.Char.Accumulator
			t.Type = "IDENT"
			p.AddToken(t)
			p.Char.Accumulator = ""
			return
		}
	}
}
