package lox

type Visitor interface {
	VisitExpr(expr *Expr) any
	VisitBinary(binary *Binary) any
	VisitGrouping(grouping *Grouping) any
	VisitLiteral(literal *Literal) any
	VisitUnary(unary *Unary) any
}

type Expression interface {
	accept(visitor Visitor) any
}

type Expr struct {
	Left     Expression
	Operator *Token
	Right    Expression
}

func NewExpr(left Expression, operator *Token, right Expression) *Expr {
	return &Expr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (expr *Expr) accept(visitor Visitor) any {
	return visitor.VisitExpr(expr)
}

type Binary struct {
	Left     Expression
	Operator *Token
	Right    Expression
}

func NewBinary(left Expression, operator *Token, right Expression) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (binary *Binary) accept(visitor Visitor) any {
	return visitor.VisitBinary(binary)
}

type Grouping struct {
	Expr Expression
}

func NewGrouping(expr Expression) *Grouping {
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
	Right    Expression
}

func NewUnary(operator *Token, right Expression) *Unary {
	return &Unary{
		Operator: operator,
		Right:    right,
	}
}

func (unary *Unary) accept(visitor Visitor) any {
	return visitor.VisitUnary(unary)
}
