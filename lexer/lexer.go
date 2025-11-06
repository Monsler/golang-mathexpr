package lexer

import (
	"bytes"
	"fmt"
)

type TokenType int

const (
	PLUS TokenType = iota
	MINUS
	DIV
	MUL
	NUMBER
	LPAREN
	RPAREN
	EOF
)

func (t TokenType) String() string {
	switch t {
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case DIV:
		return "DIV"
	case MUL:
		return "MUL"
	case NUMBER:
		return "NUMBER"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case EOF:
		return "EOF"
	default:
		return fmt.Sprintf("UNKNOWN_TOKEN_TYPE(%d)", t)
	}
}

type Token struct {
	Value    string
	Position int
	Type     TokenType
}

type Lexer struct {
	Input       string
	currentChar rune
	currentIdx  int
	tokens      []Token
}

func NewToken(value string, position int, tokenType TokenType) Token {
	return Token{Value: value, Position: position, Type: tokenType}
}

func NewLexer(input string) *Lexer {
	return &Lexer{Input: input, currentChar: rune(input[0]), currentIdx: 0, tokens: []Token{}}
}

func (l *Lexer) GetCurrentChar() rune {
	return l.currentChar
}

func (l *Lexer) step() {
	l.currentIdx++

	if l.currentIdx < len(l.Input) {
		l.currentChar = rune(l.Input[l.currentIdx])
	} else {
		l.currentChar = 0
	}
}

func (l *Lexer) parseNumber() string {
	var buf bytes.Buffer

	for l.currentChar >= '0' && l.currentChar <= '9' {
		buf.WriteRune(l.currentChar)
		l.step()
	}

	return buf.String()
}

func (l *Lexer) skipWhitespaces() {
	for l.currentChar == ' ' {
		l.step()
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.currentIdx < len(l.Input) {
		if l.currentChar >= '0' && l.currentChar <= '9' {
			pos := l.currentIdx
			l.tokens = append(l.tokens, NewToken(l.parseNumber(), pos, NUMBER))
		}

		switch l.currentChar {
		case '(':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, LPAREN))
			l.step()
		case ')':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, RPAREN))
			l.step()
		case '+':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, PLUS))
			l.step()
		case '-':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, MINUS))
			l.step()
		case '/':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, DIV))
			l.step()
		case '*':
			l.tokens = append(l.tokens, NewToken(string(l.currentChar), l.currentIdx, MUL))
			l.step()
		case ' ':
			l.skipWhitespaces()
		default:
			fmt.Printf("WARNING: Unknown token: %c; skipping\n", l.currentChar)
			l.step()
		}
	}

	l.tokens = append(l.tokens, NewToken("EOF", l.currentIdx, EOF))

	return l.tokens
}
