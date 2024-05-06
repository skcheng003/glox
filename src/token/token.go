package token

import "strconv"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (token *Token) ToString() string {
	// TODO: deal with the any type
	return strconv.Itoa(int(token.tokenType)) + " " + token.lexeme
}
