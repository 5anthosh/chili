package lexer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

const nullTerminater = '\000'

//ErrInvalidNumber #
var ErrInvalidNumber = errors.New("Invalid number")

//ErrEOF #
var ErrEOF = errors.New("EOF")

//Lexer struct
type Lexer struct {
	source  []byte
	len     uint
	tokens  []token.Token
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
		return l.nextToken(token.EOF, nil), nil
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
		return l.nextToken(token.Plus, nil), nil
	case token.MinusChar:
		return l.nextToken(token.Minus, nil), nil
	case token.StarChar:
		return l.nextToken(token.Star, nil), nil
	case token.CommonSlashChar:
		return l.nextToken(token.CommonSlash, nil), nil
	case token.OpenBracketChar:
		return l.nextToken(token.OpenBracket, nil), nil
	case token.CloseBracketChar:
		return l.nextToken(token.CloseBracket, nil), nil
	case nullTerminater:
		return l.nextToken(token.EOF, nil), nil
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
	return l.nextToken(token.Number, value), err
}

func (l *Lexer) variable() (*token.Token, error) {
	l.characters()
	return l.nextToken(token.Variable, nil), nil
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
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func (l *Lexer) characters() {
	for !l.isEnd() && isChar(l.peek(0)) || isDigit(l.peek(0)) {
		l.eat()
	}
}

func (l *Lexer) nextToken(tokenType uint, literal interface{}) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: literal,
		Lexeme:  string(l.source[l.start:l.current]),
		Column:  l.column,
	}
}
