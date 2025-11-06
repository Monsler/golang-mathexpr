package parser

import (
	"fmt"
	"strconv"

	"github.com/Monsler/devlang/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	idx    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{Tokens: tokens}
}

func (p *Parser) peek() lexer.Token {
	if p.idx >= len(p.Tokens) {
		return p.Tokens[len(p.Tokens)-1]
	}

	return p.Tokens[p.idx]
}

func (p *Parser) next() (lexer.Token, error) {
	if p.idx >= len(p.Tokens) {
		return lexer.Token{Type: lexer.EOF, Value: "", Position: 0}, nil
	}

	token := p.Tokens[p.idx]
	p.idx++

	return token, nil
}

func (p *Parser) expect(expectedType lexer.TokenType) error {
	currentTk := p.peek()

	if currentTk.Type == expectedType {
		p.next()
		return nil
	}

	err := fmt.Errorf("error: unexpected type; expected %d got %d", expectedType, currentTk.Type)

	return err
}

func (p *Parser) parseTerm() Node {
	left := p.parseFactor()

	for p.peek().Type == lexer.MUL || p.peek().Type == lexer.DIV {
		nextOp, err := p.next()

		if err == nil {
			right := p.parseFactor()

			left = &BinaryExpression{Left: left, Operator: nextOp, Right: right}
		} else {
			panic(err)
		}

	}

	return left
}

func (p *Parser) parseFactor() Node {
	token := p.peek()

	if token.Type == lexer.NUMBER {
		p.next()

		val, _ := strconv.Atoi(token.Value)
		return &NumberLiteral{Token: token, Value: val}
	}

	if token.Type == lexer.LPAREN {

		if err := p.expect(lexer.LPAREN); err != nil {
			panic(err)
		}
		node := p.parseExpression()
		if err := p.expect(lexer.RPAREN); err != nil {
			panic(err)
		}
		return node
	}

	panic("unexpected token: " + token.Value)
}

func (p *Parser) parseExpression() Node {
	left := p.parseTerm()

	for p.peek().Type == lexer.PLUS || p.peek().Type == lexer.MINUS {
		nextOp, err := p.next()

		if err == nil {
			right := p.parseTerm()

			left = &BinaryExpression{Left: left, Operator: nextOp, Right: right}
		} else {
			panic(err)
		}
	}

	return left
}

func (p *Parser) Parse() Node {
	ast := p.parseExpression()

	if p.peek().Type != lexer.EOF {
		unexpected := p.peek()
		panic(fmt.Errorf("unexpected value: \"%s\" (token of type %s) at position %d", unexpected.Value, unexpected.Type.String(), p.idx))
	}

	return ast
}
