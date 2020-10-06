package lexer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/5anthosh/eval/parser/token"
)

const nullTerminater = '\000'

//ErrInvalidNumber #
var ErrInvalidNumber = errors.New("Invalid number")

//ErrEOF #
var ErrEOF = errors.New("EOF")

type Lexer struct {
	source  []byte
	len     uint
	tokens  []token.Token
	column  uint
	start   uint
	current uint
}

//Lexer creates new lexer
func New(reader io.ReadCloser) (*Lexer, error) {
	source, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return LexerFromBytes(source), err
}

//LexerFromBytes #
func LexerFromBytes(source []byte) *Lexer {
	lex := new(Lexer)
	lex.source = source
	lex.len = uint(len(source))
	return lex
}

//LexerFromString #
func LexerFromString(source string) *Lexer {
	return LexerFromBytes([]byte(source))
}

//Next token
func (l *Lexer) Next() (*token.Token, error) {
	if l.isEnd() {
		return nil, ErrEOF
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
	}
	if isDigit(b) {
		return l.number()
	}
	//ErrUnexpectedToken #
	var errUnexpectedToken = errors.New(fmt.Sprintf("Unexpected character %c", b))
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
	lexeme, err := strconv.ParseFloat(string(l.source[l.start:l.current]), 32)
	if err != nil && !strings.Contains(err.Error(), "value out of range") {
		return nil, ErrInvalidNumber
	}

	return l.nextToken(token.Number, lexeme), nil
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

func (l *Lexer) nextToken(tokenType uint, literal interface{}) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: literal,
		Lexeme:  string(l.source[l.start:l.current]),
		Column:  l.column,
	}
}
