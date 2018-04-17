package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// GetType gets the next type from the sequence
func (p *Program) GetType() {
	t := token.NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS && p.Char.CurrentChar != " " {
			p.GetNextChar()
			if tok, ok := tokenList[p.Char.Accumulator]; ok {
				t.SetEnd(p.GetLastIndex())
				t.Value.String = tok.Value.String
				t.Type = tok.Type
				p.AddToken(t)
				return
			}
		} else {
			fmt.Println("vert a ferk")
			return
		}
	}
}
