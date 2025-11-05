package main

import (
	"fmt"

	"github.com/Monsler/devlang/evaluator"
	"github.com/Monsler/devlang/lexer"
	"github.com/Monsler/devlang/parser"
)

func main() {
	input := "(2 + 2 * 2) / 3"

	lexer := lexer.NewLexer(input)

	tokens := lexer.Tokenize()

	for i, v := range tokens {
		fmt.Printf("%d: (<%s>, %d)\n", i+1, v.Value, v.Position)
	}

	parser := parser.NewParser(tokens)

	tree := parser.Parse()

	result, err := evaluator.Evaluate(tree)

	fmt.Println("AST: ", tree.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Result: ", result)
}
