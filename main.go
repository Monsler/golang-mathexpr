package main

import (
	"fmt"

	"github.com/Monsler/devlang/evaluator"
	"github.com/Monsler/devlang/lexer"
	"github.com/Monsler/devlang/parser"
)

func main() {
	input := "x + (x / x)"

	lexer := lexer.NewLexer(input)

	tokens := lexer.Tokenize()

	for i, v := range tokens {
		fmt.Printf("%d: (<%s>, %d) == %s\n", i+1, v.Value, v.Position, v.Type.String())
	}

	vars := make(map[string]int)
	vars["x"] = 10

	parser := parser.NewParser(tokens, vars)

	tree := parser.Parse()

	result, err := evaluator.Evaluate(tree)

	fmt.Println("AST: ", tree.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Result: ", result)
}
