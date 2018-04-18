package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sgg7269/go-interpreter/parser"
	"github.com/sgg7269/go-interpreter/token"
)

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

	p := parser.Program{
		Value:  string(input),
		Length: len(input),
		Char: parser.Char{
			CurrentChar: string(input[0]),
		},
		TokenPipeline: map[string]token.Token{
			"last":     token.Token{},
			"current":  token.Token{},
			"expected": token.Token{},
		},
	}

	fmt.Printf("Program start: %#v\n\n", p)

	p.GetProgram()

	fmt.Printf("Program end: %#v\n\n", p)

	fmt.Println("Tokens:")
	for _, token := range p.CollectedTokens {
		fmt.Printf("Program end: %#v\n", token)
	}
}
