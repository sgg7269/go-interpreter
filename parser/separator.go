package parser

import (
	"errors"

	"github.com/sgg7269/go-interpreter/token"
)

// GetSeparator ...
// TODO: fix this
func (p *Program) GetSeparator() (bool, error) {
	for {
		if !p.EOF && !p.Char.EOS {
			if p.Char.CurrentChar == SeparatorToken {
				t := token.NewToken(p.GetCurrentIndex())
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value.String = p.Char.Accumulator
				t.Type = "SEP"
				// p.AddToken(t)
				p.Char.Accumulator = ""
				return false, nil
			}
			return false, errors.New("not a separator")
		}
		return true, nil
	}
}
