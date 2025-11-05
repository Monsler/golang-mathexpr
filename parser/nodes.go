package parser

import "github.com/Monsler/devlang/lexer"

type Node interface {
	String() string
}

type NumberLiteral struct {
	Token lexer.Token
	Value int
}

func (nl *NumberLiteral) String() string {
	return nl.Token.Value
}

type BinaryExpression struct {
	Left     Node
	Operator lexer.Token
	Right    Node
}

func (be *BinaryExpression) String() string {
	return "(" + be.Left.String() + " " + be.Operator.Value + " " + be.Right.String() + ")"
}
