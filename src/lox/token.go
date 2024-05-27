package lox

import (
	"strconv"
)

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	BANG       // "!"
	BANG_EQUAL // "!="
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER

	AND
	OR
	CLASS
	IF
	ELSE
	WHILE
	TRUE
	FALSE
	FUN
	FOR
	NIL
	PRINT
	RETURN
	SUPER
	THIS
	VAR

	EOF
)

type Token struct {
	tokenType TokenType
	Lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		tokenType: tokenType,
		Lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (token *Token) ToString() string {
	// TODO: deal with the any type
	return strconv.Itoa(int(token.tokenType)) + " " + token.Lexeme
}
