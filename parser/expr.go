package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/sgg7269/go-interpreter/token"
)

var secOpList = map[string]token.Token{
	"=": token.Token{
		ID:   20,
		Type: "EQUALS",
		Value: token.Value{
			String: "=",
		},
	},
	"++": token.Token{
		ID:   21,
		Type: "INCRE",
		Value: token.Value{
			String: "++",
		},
	},
}

// GetExpr ...
func (p *Program) GetExpr() {
	// p.GetAssignment()
	p.GetTerm()
	// TODO: This needs to check EOS
	eos, err := p.GetSeparator()
	if err != nil {
		fmt.Println("OMFG", err)
		os.Exit(6)
	}
	if eos {
		fmt.Println("got to eos")
		return
	}

	op, err := p.GetSecOp()
	if err != nil {
		fmt.Println("OMFG OP", err)
		os.Exit(6)
	}
	if op == "=" {
		fmt.Println("got = op")

		eos, err = p.GetSeparator()
		if err != nil {
			fmt.Println("OMFG", err)
			os.Exit(6)
		} else if eos {
			fmt.Println("Expected expression after operand")
			return
		}

		p.GetConstant()
		return
	}

	eos, err = p.GetSeparator()
	if err != nil {
		fmt.Println("OMFG", err)
		os.Exit(6)
	} else if eos {
		fmt.Println("Expected expression after operand")
		return
	}
	// TODO: This needs to check what it got for an error
	p.GetExpr()
}

// // GetAssignment ...
// func (p *Program) GetAssignment() {
// 	p.GetTerm()
// 	// TODO: This needs to check EOS
// 	p.GetSeparator()
// 	p.GetSecOp()
// 	p.GetSeparator()
// 	// TODO: This needs to check what it got for an error
// 	p.GetVar()
// }

// GetSecOp ...
func (p *Program) GetSecOp() (string, error) {
	t := token.NewToken(p.GetCurrentIndex())

	for {
		if !p.EOF && !p.Char.EOS {
			if p.Char.CurrentChar != " " {
				// t := token.NewToken(p.GetCurrentIndex())
				p.GetNextChar()
			} else {
				if secOpToken, ok := secOpList[p.Char.Accumulator]; ok {
					secOpToken.Location[0] = t.Location[0]
					secOpToken.SetEnd(p.GetLastIndex())
					p.AddToken(secOpToken)
					p.Char.Accumulator = ""
					return secOpToken.Value.String, nil
				}
				fmt.Println("wtf")
				return "", errors.New("wtf")
			}
		}
	}
}
