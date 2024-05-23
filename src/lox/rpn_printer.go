package lox

import "fmt"

type RpnPrinter struct {
}

func NewRpnPrinter() *RpnPrinter {
	return &RpnPrinter{}
}

func (r *RpnPrinter) Print(expr Expr) string {
	str := fmt.Sprintf("%v", expr.accept(r))
	fmt.Println(str)
	return str
}

func (p *RpnPrinter) VisitBinary(binary *Binary) any {
	return p.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (p *RpnPrinter) VisitGrouping(grouping *Grouping) any {
	return p.parenthesize("grouping", grouping.Expr)
}

func (p *RpnPrinter) VisitLiteral(literal *Literal) any {
	if literal.Value == nil {
		return "nil"
	}
	return literal.Value
}

func (p *RpnPrinter) VisitUnary(unary *Unary) any {
	return p.parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (r *RpnPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder string
	for _, expr := range exprs {
		str := expr.accept(r).(string)
		builder += str + " "
	}
	builder += name
	return builder
}
