package token

import "fmt"

// Token character
const (
	PlusChar        = '+'
	MinusChar       = '-'
	StarChar        = '*'
	CommonSlashChar = '/'
	OpenParenChar   = '('
	CloseParenChar  = ')'
)

// Type of tokens
const (
	PlusType = 1 + iota
	MinusType
	StarType
	CommonSlashType
	NumberType
	OpenParenType
	CloseParenType
	VariableType
	EOFType
)

//Type of the token
type Type interface {
	String() string
	Type() uint
}

//Plus symbol "+"
type Plus struct{}

func (p Plus) String() string {
	return "Plus"
}

//Type of symbol
func (p Plus) Type() uint {
	return PlusType
}

//Star symbol "*"e
type Star struct{}

func (s Star) String() string {
	return "Star"
}

//Type of symbol
func (s Star) Type() uint {
	return StarType
}

//Minus symbol "-"
type Minus struct{}

func (m Minus) String() string {
	return "Minus"
}

//Type of symbol
func (m Minus) Type() uint {
	return MinusType
}

//CommonSlash symbol "/"
type CommonSlash struct{}

func (c CommonSlash) String() string {
	return "Common Slash"
}

//Type of symbol
func (c CommonSlash) Type() uint {
	return CommonSlashType
}

//Number symbol
type Number struct{}

func (n Number) String() string {
	return "Number"
}

//Type of symbol
func (n Number) Type() uint {
	return NumberType
}

//OpenParen symbol "("
type OpenParen struct{}

func (o OpenParen) String() string {
	return "Open Bracket"
}

//Type of symbol
func (o OpenParen) Type() uint {
	return OpenParenType
}

//CloseParen symbol ")"
type CloseParen struct{}

func (c CloseParen) String() string {
	return "Close Bracket"
}

//Type of symbol
func (c CloseParen) Type() uint {
	return CloseParenType
}

//Variable symbol
type Variable struct{}

func (v Variable) String() string {
	return "Variable"
}

//Type of symbol
func (v Variable) Type() uint {
	return VariableType
}

//EOF symbol
type EOF struct{}

func (e EOF) String() string {
	return "EOF"
}

//Type of symbol
func (e EOF) Type() uint {
	return EOFType
}

//Token character stream of expression is token into token
type Token struct {
	Type    Type
	Literal interface{}
	Lexeme  string
	Column  uint
}

func (t Token) String() string {
	return fmt.Sprintf("< %s %s %v %d>", t.Type.String(), t.Lexeme, t.Literal, t.Column)
}
