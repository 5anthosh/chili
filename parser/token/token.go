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
	CapChar         = '^'
	ModChar         = '%'
	CommaChar       = ','
	QuoteChar       = '\''
	DoubleQuoteChar = '"'
	EqualChar       = '='
	PunctuationChar = '!'
	PipeChar        = '|'
	AndChar         = '&'
	QuestionChar    = '?'
	ColonChar       = ':'
	GreaterChar     = '>'
	LesserChar      = '<'
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
	CapType
	ModType
	CommaType
	StringType
	EqualType
	NotEqualType
	GreaterType
	GreaterEqualType
	LesserType
	LesserEqualType
	OrType
	AndType
	NotType
	BooleanType
	QuestionType
	ColonType
	EOFType
)

//Type of the token
type Type interface {
	String() string
	Type() uint
}

//Plus symbol "+"
type Plus struct{}

func (Plus) String() string {
	return "Plus"
}

//Type of symbol
func (Plus) Type() uint {
	return PlusType
}

//Star symbol "*"e
type Star struct{}

func (Star) String() string {
	return "Star"
}

//Type of symbol
func (Star) Type() uint {
	return StarType
}

//Minus symbol "-"
type Minus struct{}

func (Minus) String() string {
	return "Minus"
}

//Type of symbol
func (Minus) Type() uint {
	return MinusType
}

//CommonSlash symbol "/"
type CommonSlash struct{}

func (CommonSlash) String() string {
	return "Common Slash"
}

//Type of symbol
func (CommonSlash) Type() uint {
	return CommonSlashType
}

//Number symbol
type Number struct{}

func (Number) String() string {
	return "Number"
}

//Type of symbol
func (Number) Type() uint {
	return NumberType
}

//OpenParen symbol "("
type OpenParen struct{}

func (OpenParen) String() string {
	return "Open Bracket"
}

//Type of symbol
func (OpenParen) Type() uint {
	return OpenParenType
}

//CloseParen symbol ")"
type CloseParen struct{}

func (CloseParen) String() string {
	return "Close Bracket"
}

//Type of symbol
func (CloseParen) Type() uint {
	return CloseParenType
}

//Variable symbol
type Variable struct{}

func (Variable) String() string {
	return "Variable"
}

//Type of symbol
func (Variable) Type() uint {
	return VariableType
}

//Cap symbol
type Cap struct{}

func (Cap) String() string {
	return "Cap"
}

//Type of symbol
func (Cap) Type() uint {
	return CapType
}

//Mod symbol
type Mod struct{}

func (Mod) String() string {
	return "Mod"
}

//Type of symbol
func (Mod) Type() uint {
	return ModType
}

//Comma symbol
type Comma struct{}

func (Comma) String() string {
	return "Comma"
}

//Type of symbol
func (Comma) Type() uint {
	return CommaType
}

//LiteralString value
type LiteralString struct{}

func (LiteralString) String() string {
	return "String"
}

//Type of symbol
func (LiteralString) Type() uint {
	return StringType
}

//Equal == symbol
type Equal struct{}

func (Equal) String() string {
	return "Equal"
}

//Type of token
func (Equal) Type() uint {
	return EqualType
}

//NotEqual != symbol
type NotEqual struct{}

func (NotEqual) String() string {
	return "Not Equal"
}

//Type of token
func (NotEqual) Type() uint {
	return NotEqualType
}

//And && symbol
type And struct{}

func (And) String() string {
	return "And"
}

//Type of token
func (And) Type() uint {
	return AndType
}

//Or || symbol
type Or struct{}

func (Or) String() string {
	return "Or"
}

//Type of token
func (Or) Type() uint {
	return OrType
}

//Boolean || symbol
type Boolean struct{}

func (Boolean) String() string {
	return "Boolean"
}

//Type of token
func (Boolean) Type() uint {
	return BooleanType
}

//Question ? symbol
type Question struct{}

func (Question) String() string {
	return "Question"
}

//Type of token
func (Question) Type() uint {
	return QuestionType
}

//Colon : symbol
type Colon struct{}

func (Colon) String() string {
	return "Colon"
}

//Type of token
func (Colon) Type() uint {
	return ColonType
}

//Not ! symbol
type Not struct{}

func (Not) String() string {
	return "Not"
}

//Type of token
func (Not) Type() uint {
	return NotType
}

//Greater ! symbol
type Greater struct{}

func (Greater) String() string {
	return "Greater"
}

//Type of token
func (Greater) Type() uint {
	return GreaterType
}

//GreaterEqual ! symbol
type GreaterEqual struct{}

func (GreaterEqual) String() string {
	return "GreaterEqual"
}

//Type of token
func (GreaterEqual) Type() uint {
	return GreaterEqualType
}

//Lesser ! symbol
type Lesser struct{}

func (Lesser) String() string {
	return "Lesser"
}

//Type of token
func (Lesser) Type() uint {
	return LesserType
}

//LesserEqual ! symbol
type LesserEqual struct{}

func (LesserEqual) String() string {
	return "LesserEqual"
}

//Type of token
func (LesserEqual) Type() uint {
	return LesserEqualType
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
