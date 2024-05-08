package lox

type Expr struct {
	Left *Expr
	Operator *Token
	Right *Expr
}

func NewExpr(left *Expr, operator *Token, right *Expr) *Expr {
	return &Expr{
		Left: left,
		Operator: operator,
		Right: right,
	}
}

type Binary struct {
	Left *Expr
	Operator *Token
	Right *Expr
}

func NewBinary(left *Expr, operator *Token, right *Expr) *Binary {
	return &Binary{
		Left: left,
		Operator: operator,
		Right: right,
	}
}

type Grouping struct {
	Expr *Expr
}

func NewGrouping(expr *Expr) *Grouping {
	return &Grouping{
		Expr: expr,
	}
}

type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

type Unary struct {
	Operator *Token
	Right *Expr
}

func NewUnary(operator *Token, right *Expr) *Unary {
	return &Unary{
		Operator: operator,
		Right: right,
	}
}

