package lox

import "fmt"

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) Print(expr Expr) string {
	str := fmt.Sprintf("%v", expr.accept(p))
	fmt.Println(str)
	return str
}

func (p *AstPrinter) VisitBinary(binary *Binary) any {
	return p.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (p *AstPrinter) VisitGrouping(grouping *Grouping) any {
	return p.parenthesize("grouping", grouping.Expr)
}

func (p *AstPrinter) VisitLiteral(literal *Literal) any {
	if literal.Value == nil {
		return "nil"
	}
	return literal.Value
}

func (p *AstPrinter) VisitUnary(unary *Unary) any {
	return p.parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (p *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder string
	builder += "(" + name
	for _, expr := range exprs {
		str := fmt.Sprintf("%v", expr.accept(p))
		builder += " " + str
	}
	builder += ")"
	return builder
}
