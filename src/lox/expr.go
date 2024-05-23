package lox

type Visitor interface {
	VisitBinary(binary *Binary) any
	VisitGrouping(grouping *Grouping) any
	VisitLiteral(literal *Literal) any
	VisitUnary(unary *Unary) any
}

type Expr interface {
	accept(visitor Visitor) any
}

type Binary struct {
	Left Expr
	Operator *Token
	Right Expr
}

func NewBinary(left Expr, operator *Token, right Expr) *Binary {
	return &Binary{
		Left: left,
		Operator: operator,
		Right: right,
	}
}

func (binary *Binary) accept(visitor Visitor) any {
	return visitor.VisitBinary(binary)
}

type Grouping struct {
	Expr Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		Expr: expr,
	}
}

func (grouping *Grouping) accept(visitor Visitor) any {
	return visitor.VisitGrouping(grouping)
}

type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

func (literal *Literal) accept(visitor Visitor) any {
	return visitor.VisitLiteral(literal)
}

type Unary struct {
	Operator *Token
	Right Expr
}

func NewUnary(operator *Token, right Expr) *Unary {
	return &Unary{
		Operator: operator,
		Right: right,
	}
}

func (unary *Unary) accept(visitor Visitor) any {
	return visitor.VisitUnary(unary)
}

