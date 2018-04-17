package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// GetEOS ...
func (p *Program) GetEOS() {
	for {
		if !p.EOF {
			if p.Char.EOS {
				t := token.NewToken(p.GetCurrentIndex())
				p.GetNextChar()
				t.SetEnd(p.GetLastIndex())
				t.Value.String = p.Char.Accumulator
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
