package parser

import (
	"glox/src/internal/scanner"
)

type Expr struct {
	left     *Expr
	operator *scanner.Token
	right    *Expr
}

func NewExpr(left *Expr, operator *scanner.Token, right *Expr) *Expr {
	return &Expr{
		left:     left,
		operator: operator,
		right:    right,
	}
}

type Binary struct {
	Expr *Expr
}

func NewBinary(left, right *Expr, operator *scanner.Token) *Binary {
	return &Binary{
		Expr: NewExpr(left, operator, right),
	}
}
