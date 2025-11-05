package evaluator

import (
	"fmt"

	"github.com/Monsler/devlang/lexer"
	"github.com/Monsler/devlang/parser"
)

func applyOperation(opType lexer.TokenType, left, right int) (int, error) {
	switch opType {
	case lexer.PLUS:
		return left + right, nil
	case lexer.MINUS:
		return left - right, nil
	case lexer.DIV:
		return left / right, nil
	case lexer.MUL:
		return left * right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %d", opType)
	}
}

func Evaluate(node parser.Node) (int, error) {
	switch n := node.(type) {
	case *parser.NumberLiteral:
		return n.Value, nil
	case *parser.BinaryExpression:
		leftVal, err := Evaluate(n.Left)
		if err != nil {
			return 0, err
		}

		rightVal, err := Evaluate(n.Right)
		if err != nil {
			return 0, err
		}

		return applyOperation(n.Operator.Type, leftVal, rightVal)
	default:
		return 0, fmt.Errorf("unknown eval type")
	}
}
