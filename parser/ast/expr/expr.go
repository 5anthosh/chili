package expr

import "github.com/5anthosh/eval/parser/token"

//Expr interface
type Expr interface {
	Accept(Visitor) (interface{}, error)
}

//Visitor interface
type Visitor interface {
	VisitBinaryExpr(binaryExpression *Binary) (interface{}, error)
	VisitGroupExpr(groupExpression *Group) (interface{}, error)
	VisitLiteralExpr(LiteralExpression *Literal) (interface{}, error)
	VisitUnaryExpr(unaryExpr *Unary) (interface{}, error)
}

//Binary #
type Binary struct {
	Left     Expr
	Right    Expr
	Operator *token.Token
}

//Accept binary operation
func (b *Binary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitBinaryExpr(b)
}

//Group #
type Group struct {
	Expression Expr
}

//Accept Group exp
func (g *Group) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitGroupExpr(g)
}

//Literal #
type Literal struct {
	Value interface{}
}

//Accept Literal expression
func (l *Literal) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitLiteralExpr(l)
}

//Unary #
type Unary struct {
	Operator *token.Token
	Right    Expr
}

//Accept Unary expr
func (u *Unary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitUnaryExpr(u)
}
