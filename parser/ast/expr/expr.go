package expr

import "github.com/5anthosh/chili/parser/token"

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
	VisitVariableExpr(variableExpr *Variable) (interface{}, error)
	VisitFunctionCall(functionCallExpr *FunctionCall) (interface{}, error)
	VisitTernary(ternaryExpr *Ternary) (interface{}, error)
	VisitLogicalExpr(logicalExpr *Logical) (interface{}, error)
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

//Variable #
type Variable struct {
	Name string
}

//Accept variable expression
func (v *Variable) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitVariableExpr(v)
}

//FunctionCall #
type FunctionCall struct {
	Name string
	Args []Expr
}

//Accept #
func (f *FunctionCall) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitFunctionCall(f)
}

//Ternary  #
type Ternary struct {
	Condition Expr
	True      Expr
	False     Expr
}

//Accept #
func (t *Ternary) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitTernary(t)
}

//Logical operation
type Logical struct {
	Left     Expr
	Right    Expr
	Operator *token.Token
}

//Accept #
func (t *Logical) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitLogicalExpr(t)
}
