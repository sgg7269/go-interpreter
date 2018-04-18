package parser

import (
	"fmt"

	"github.com/sgg7269/go-interpreter/token"
)

// Program is the Overarching program type that holds information about the current program parse
type Program struct {
	Value           string
	Length          int
	Index           int
	EOF             bool
	Char            Char
	Statement       Statement
	TokenPipeline   map[string]token.Token
	CollectedTokens []token.Token
}

// // NewToken ...
// func (p *Program) NewToken() Token {
// 	return Token{Location: [2]int{p.Index, -1}}
// }

// GetProgram ...
func (p *Program) GetProgram() {
	// p.GetStatement()
	for {
		t := p.GetNextToken()

		switch t.ID {
		case -1:
			return

		case 1:
			t := p.GetNextToken()

			if t.ID != -3 {
				fmt.Println("wtf1")
			}

		case 2:
			// t := p.GetNextToken()

			// if t.ID != -3 {
			// 	fmt.Println("wtf2")
			// }
			p.GetConstant()
			p.GetEOS()
		}

		// fmt.Println(t)
	}
}

// // GetProgram ...
// func (p *Program) GetProgram() {
// 	fmt.Println(p.Value + "\n")

// 	// Split by semicolon
// 	// Split by encompassers
// 	// Split by space

// 	semicolonSplit := strings.Split(p.Value, ";")
// 	fmt.Printf("%#v\n\n", semicolonSplit)

// 	for _, split := range semicolonSplit {
// 		spaceSplit := strings.Split(split, " ")
// 		fmt.Printf("%#v\n\n", spaceSplit)
// 	}
// }
