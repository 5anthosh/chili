package lexer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/5anthosh/chili/parser/token"
	"github.com/shopspring/decimal"
)

const nullTerminater = '\000'

//ErrInvalidNumber #
var ErrInvalidNumber = errors.New("Invalid number")

//ErrEOF #
var ErrEOF = errors.New("EOF")

//Lexer struct
type Lexer struct {
	// expression source
	source []byte
	// Length of expression source
	len uint
	// Scanned tokens
	tokens []token.Token
	// Current column
	column  uint
	start   uint
	current uint
}

//New  creates new lexer
func New(reader io.ReadCloser) (*Lexer, error) {
	source, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FromBytes(source), err
}

//FromBytes #
func FromBytes(source []byte) *Lexer {
	lex := new(Lexer)
	lex.source = source
	lex.len = uint(len(source))
	return lex
}

//FromString #
func FromString(source string) *Lexer {
	return FromBytes([]byte(source))
}

//Next token
func (l *Lexer) Next() (*token.Token, error) {
	if l.isEnd() {
		return l.nextToken(token.EOF{}, nil), nil
	}
	l.start = l.current
	t, err := l.scan()
	if err != nil {
		return nil, err
	}
	l.tokens = append(l.tokens, *t)
	return t, err
}

func (l *Lexer) scan() (*token.Token, error) {
	b := l.eat()
	b = l.space(b)
	switch b {
	case token.PlusChar:
		return l.nextToken(token.Plus{}, nil), nil
	case token.MinusChar:
		return l.nextToken(token.Minus{}, nil), nil
	case token.StarChar:
		return l.nextToken(token.Star{}, nil), nil
	case token.CommonSlashChar:
		return l.nextToken(token.CommonSlash{}, nil), nil
	case token.OpenParenChar:
		return l.nextToken(token.OpenParen{}, nil), nil
	case token.CloseParenChar:
		return l.nextToken(token.CloseParen{}, nil), nil
	case nullTerminater:
		return l.nextToken(token.EOF{}, nil), nil
	case token.CapChar:
		return l.nextToken(token.Cap{}, nil), nil
	case token.ModChar:
		return l.nextToken(token.Mod{}, nil), nil
	case token.CommaChar:
		return l.nextToken(token.Comma{}, nil), nil
	case token.QuoteChar:
		return l.stringLiteral(token.QuoteChar)
	case token.DoubleQuoteChar:
		return l.stringLiteral(token.DoubleQuoteChar)
	case token.QuestionChar:
		return l.nextToken(token.Question{}, nil), nil
	case token.ColonChar:
		return l.nextToken(token.Colon{}, nil), nil
	case token.PipeChar:
		if l.peek(0) == token.PipeChar {
			l.eat()
			return l.nextToken(token.Or{}, nil), nil
		}
		return nil, fmt.Errorf("Expecting '|' after |")
	case token.AndChar:
		if l.peek(0) == token.AndChar {
			l.eat()
			return l.nextToken(token.And{}, nil), nil
		}
		return nil, fmt.Errorf("Expecting '&' after &")
	case token.EqualChar:
		if l.peek(0) == token.EqualChar {
			l.eat()
			return l.nextToken(token.Equal{}, nil), nil
		}
		return nil, fmt.Errorf("Expecting '=' after =")
	case token.PunctuationChar:
		if l.peek(0) == token.EqualChar {
			l.eat()
			return l.nextToken(token.NotEqual{}, nil), nil
		}
		return l.nextToken(token.Not{}, nil), nil
	case token.GreaterChar:
		if l.peek(0) == token.EqualChar {
			l.eat()
			return l.nextToken(token.GreaterEqual{}, nil), nil
		}
		return l.nextToken(token.Greater{}, nil), nil
	case token.LesserChar:
		if l.peek(0) == token.EqualChar {
			l.eat()
			return l.nextToken(token.LesserEqual{}, nil), nil
		}
		return l.nextToken(token.LesserEqual{}, nil), nil
	}
	if isDigit(b) {
		return l.number()
	}

	if isChar(b) {
		return l.variable()
	}
	//ErrUnexpectedToken #
	var errUnexpectedToken = fmt.Errorf("Unexpected character %c", b)
	return nil, errUnexpectedToken
}

func (l *Lexer) number() (*token.Token, error) {
	l.digits()
	if l.match('.') {
		if !isDigit(l.peek(0)) {
			return nil, ErrInvalidNumber
		}
		l.eat()
		l.digits()
	}

	if l.match('e') || l.match('E') {
		if !l.match('-') {
			l.match('+')
		}
		if !isDigit(l.peek(0)) {
			return nil, ErrInvalidNumber
		}
		l.eat()
		l.digits()
	}
	value, err := decimal.NewFromString(string(l.source[l.start:l.current]))
	return l.nextToken(token.Number{}, value), err
}

func (l *Lexer) variable() (*token.Token, error) {
	l.characters()
	value := l.source[int(l.start):int(l.current)]
	if string(value) == "true" {
		return l.nextToken(token.Boolean{}, true), nil
	}
	if string(value) == "false" {
		return l.nextToken(token.Boolean{}, false), nil
	}
	return l.nextToken(token.Variable{}, nil), nil
}

func (l *Lexer) eat() byte {
	l.current++
	l.column++
	return l.source[l.current-1]
}

func (l *Lexer) match(expected byte) bool {
	if l.isEnd() {
		return false
	}
	if l.source[l.current] != expected {
		return false
	}
	l.current++
	l.column++
	return true
}

func (l *Lexer) space(b byte) byte {
	for isEmptySpace(b) {
		if !l.isEnd() {
			l.start = l.current
			b = l.eat()
		} else {
			return nullTerminater
		}
	}
	return b
}

func (l Lexer) peek(b uint) byte {
	if l.current+b >= l.len {
		return nullTerminater
	}
	return l.source[l.current+b]
}

func (l *Lexer) isEnd() bool {
	return l.current >= l.len
}

func isEmptySpace(b byte) bool {
	return b == ' ' || b == '\r' || b == '\t' || b == '\n'
}
func (l *Lexer) digits() {
	for !l.isEnd() && isDigit(l.peek(0)) {
		l.eat()
	}
}
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isChar(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b == '_'
}

func (l *Lexer) characters() {
	for !l.isEnd() && (isChar(l.peek(0)) || isDigit(l.peek(0)) || l.peek(0) == '.') {
		l.eat()
	}
}

func (l *Lexer) nextToken(tokenType token.Type, literal interface{}) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: literal,
		Lexeme:  string(l.source[l.start:l.current]),
		Column:  l.column,
	}
}

func (l *Lexer) stringLiteral(s byte) (*token.Token, error) {
	for {
		if l.isEnd() {
			return nil, fmt.Errorf("Expecting %c but found EOF", s)
		}
		if l.peek(0) == s {
			l.eat()
			break
		}
		l.eat()
	}
	return l.nextToken(token.LiteralString{}, string(l.source[l.start+1:l.current-1])), nil
}
