package parser

import (
	"fmt"
	"strconv"

	"github.com/sgg7269/go-interpreter/token"
)

// GetConstant ...
func (p *Program) GetConstant() {
	t := token.NewToken(p.GetCurrentIndex())

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
			t.Value.String = p.Char.Accumulator
			t.Value.True = accumulatorInt
			t.Type = "INT_LIT"
			p.AddToken(t)
			p.Char.Accumulator = ""
			return
		}
	}
}
