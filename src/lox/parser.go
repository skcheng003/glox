package lox

import (
	"errors"
)

var ErrParseToken = errors.New("parser error")

type Parser struct {
	tokens  []*Token
	current int
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) Parse() Expr {
	return p.expression()
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(PLUS, MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.unary()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(TRUE) {
		return NewLiteral(true)
	}
	if p.match(FALSE) {
		return NewLiteral(false)
	}
	if p.match(NIL) {
		return NewLiteral(nil)
	}
	if p.match(NUMBER, STRING) {
		return NewLiteral(p.previous().literal)
	}
	if p.match(LEFT_PAREN) {
		expr := p.expression()
		_, err := p.consume(RIGHT_PAREN)
		if err != nil {
			ErrorHandler(err, p.peek(), "Expect ')' after expression.")
		}
		return NewGrouping(expr)
	}
	ErrorHandler(ErrParseToken, p.peek(), "Expect expression")
	return nil
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, elem := range tokenTypes {
		if p.check(elem) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tokenType TokenType) (*Token, error) {
	if !p.check(tokenType) {
		return nil, ErrParseToken
	}
	return p.advance(), nil
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().tokenType == tokenType
}

func (p *Parser) advance() *Token {
	if p.isAtEnd() != true {
		p.current += 1
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().tokenType == EOF
}

// peek current token
func (p *Parser) peek() *Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().tokenType == SEMICOLON {
			return
		}
		switch p.peek().tokenType {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}
	}
	p.advance()
}
