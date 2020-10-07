package parser

import (
	"errors"
	"fmt"

	"github.com/5anthosh/eval/parser/ast/expr"
	"github.com/5anthosh/eval/parser/lexer"
	"github.com/5anthosh/eval/parser/token"
)

//Parser struct
type Parser struct {
	source string
	lex    *lexer.Lexer
	n      uint
	tokens []*token.Token
}

//New Parser
func New(source string) *Parser {
	newParser := new(Parser)
	newParser.source = source
	newParser.lex = lexer.FromString(source)
	newParser.tokens = make([]*token.Token, 0)
	return newParser
}

//Parse the expression and returns AST
func (p *Parser) Parse() (expr.Expr, error) {
	return p.expression()
}

func (p *Parser) expression() (expr.Expr, error) {
	return p.addition()
}

func (p *Parser) addition() (expr.Expr, error) {
	expression, err := p.multiply()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.Plus, token.Minus})
		if err != nil {
			return nil, err
		}
		if ok {
			operator := p.previous()
			right, err := p.multiply()
			if err != nil {
				return nil, err
			}
			expression = &expr.Binary{Left: expression, Right: right, Operator: operator}
			continue
		}
		break
	}
	return expression, nil
}

func (p *Parser) multiply() (expr.Expr, error) {
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.Star, token.CommonSlash})
		if err != nil {
			return nil, err
		}
		if ok {
			operator := p.previous()
			right, err := p.unary()
			if err != nil {
				return nil, err
			}
			expression = &expr.Binary{Left: expression, Right: right, Operator: operator}
			continue
		}
		break
	}
	return expression, nil
}

func (p *Parser) unary() (expr.Expr, error) {
	ok, err := p.match([]uint{token.Plus, token.Minus})
	if err != nil {
		return nil, err
	}
	if ok {
		t := p.previous()
		unaryExpr, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &expr.Unary{Operator: t, Right: unaryExpr}, nil
	}
	return p.term()
}

func (p *Parser) term() (expr.Expr, error) {
	ok, err := p.match([]uint{token.Number})
	if err != nil {
		return nil, err
	}
	if ok {
		numberExpression := p.previous()
		return &expr.Literal{Value: numberExpression.Literal}, nil
	}
	ok, err = p.match([]uint{token.OpenBracket})
	if err != nil {
		return nil, err
	}
	if ok {
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		peekValue := "EOF"
		peekToken, err := p.peek()
		if err != nil {
			return nil, err
		}
		if peekToken.Type != token.EOF {
			peekValue = peekToken.Lexeme
		}
		err = p.consume(token.CloseBracket, fmt.Sprintf("Expect ')' after expression but found %s", peekValue))
		if err != nil {
			return nil, err
		}
		return &expr.Group{Expression: expression}, nil
	}
	t, err := p.peek()
	if err != nil {
		return nil, err
	}
	peekValue := "EOF"
	if t.Type != token.EOF {
		peekValue = t.Lexeme
	}
	return nil, fmt.Errorf("Expect Expression but found %s", peekValue)
}

func (p *Parser) getToken() (*token.Token, error) {
	t, err := p.lex.Next()
	if err == nil {
		if t.Type != token.EOF {
			p.tokens = append(p.tokens, t)
		}
		return t, nil
	}
	return nil, err
}

func (p *Parser) match(tokenTypes []uint) (bool, error) {
	for _, tokenType := range tokenTypes {
		ok, err := p.check(tokenType)
		if err != nil {
			return false, err
		}
		if ok {
			p.increment()
			return true, nil
		}
	}
	return false, nil
}

func (p *Parser) consume(tokenType uint, message string) error {
	ok, err := p.check(tokenType)
	if err != nil {
		return err
	}
	if ok {
		p.increment()
		return nil
	}
	return errors.New(message)
}

func (p *Parser) check(tokenType uint) (bool, error) {
	ok, err := p.isAtEnd()
	if err != nil {
		return false, err
	}
	if ok {
		return false, nil
	}
	t, err := p.peek()
	if err != nil {
		return false, err
	}
	return t.Type == tokenType, nil
}

func (p *Parser) peek() (*token.Token, error) {
	return p.nextToken()
}
func (p *Parser) isAtEnd() (bool, error) {
	t, err := p.nextToken()
	if err != nil {
		return false, err
	}
	return t.Type == token.EOF, nil
}

func (p *Parser) previous() *token.Token {
	return p.tokens[p.n-1]
}

func (p *Parser) nextToken() (*token.Token, error) {
	if p.n < uint(len(p.tokens)) {
		return p.tokens[p.n], nil
	}
	return p.getToken()
}

func (p *Parser) increment() {
	p.n = p.n + 1
}
