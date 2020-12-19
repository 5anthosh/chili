package parser

import (
	"errors"
	"fmt"

	"github.com/5anthosh/chili/parser/ast/expr"
	"github.com/5anthosh/chili/parser/lexer"
	"github.com/5anthosh/chili/parser/token"
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
	return p.ternary()
}

func (p *Parser) ternary() (expr.Expr, error) {
	expression, err := p.logical()
	if err != nil {
		return nil, err
	}
	ok, err := p.match([]uint{token.QuestionType})
	if err != nil {
		return nil, err
	}
	if !ok {
		return expression, nil
	}
	trueExpr, err := p.expression()
	if err != nil {
		return nil, err
	}
	ok, err = p.match([]uint{token.ColonType})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("Expecting : in ternary operation")
	}
	falseExpr, err := p.expression()
	if err != nil {
		return nil, err
	}
	return &expr.Ternary{
		Condition: expression,
		True:      trueExpr,
		False:     falseExpr,
	}, nil
}

func (p *Parser) logical() (expr.Expr, error) {
	expression, err := p.equality()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.AndType, token.OrType})
		if err != nil {
			return nil, err
		}
		if ok {
			operator := p.previous()
			right, err := p.equality()
			if err != nil {
				return nil, err
			}
			expression = &expr.Logical{Left: expression, Right: right, Operator: operator}
			continue
		}
		break
	}
	return expression, nil
}

func (p *Parser) equality() (expr.Expr, error) {
	expression, err := p.addition()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{
			token.EqualType,
			token.NotEqualType,
			token.GreaterType,
			token.GreaterEqualType,
			token.LesserType,
			token.LesserEqualType,
		})
		if err != nil {
			return nil, err
		}
		if ok {
			operator := p.previous()
			right, err := p.addition()
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

func (p *Parser) addition() (expr.Expr, error) {
	expression, err := p.multiply()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.PlusType, token.MinusType})
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
	expression, err := p.exponent()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.StarType, token.CommonSlashType, token.CapType, token.ModType})
		if err != nil {
			return nil, err
		}
		if ok {
			operator := p.previous()
			right, err := p.exponent()
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

func (p *Parser) exponent() (expr.Expr, error) {
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}
	for {
		ok, err := p.match([]uint{token.CapType})
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
	ok, err := p.match([]uint{token.PlusType, token.MinusType, token.NotType})
	if err != nil {
		return nil, err
	}
	if ok {
		t := p.previous()
		unaryExpr, err := p.functionCall()
		if err != nil {
			return nil, err
		}
		return &expr.Unary{Operator: t, Right: unaryExpr}, nil
	}
	return p.functionCall()
}

func (p *Parser) functionCall() (expr.Expr, error) {
	expression, err := p.term()
	if err != nil {
		return nil, err
	}
	ok, err := p.match([]uint{token.OpenParenType})
	if err != nil {
		return nil, err
	}
	if ok {
		switch expression.(type) {
		case *expr.Variable:
			name := expression.(*expr.Variable).Name
			var args []expr.Expr
			ok, err := p.match([]uint{token.CloseParenType})
			if err != nil {
				return nil, err
			}
			if ok {
				return &expr.FunctionCall{Name: name, Args: args}, nil
			}
			for {
				arg, err := p.expression()
				if err != nil {
					return nil, err
				}
				args = append(args, arg)
				ok, err := p.match([]uint{token.CommaType})
				if err != nil {
					return nil, err
				}
				if !ok {
					break
				}
			}
			ok, err = p.match([]uint{token.CloseParenType})
			if err != nil {
				return nil, err
			}
			if ok {
				return &expr.FunctionCall{Name: name, Args: args}, nil
			}
			return nil, errors.New("Expecting ')' after arguments")
		}
	}
	return expression, err
}

func (p *Parser) term() (expr.Expr, error) {
	ok, err := p.match([]uint{token.NumberType, token.StringType, token.BooleanType})
	if err != nil {
		return nil, err
	}
	if ok {
		literal := p.previous()
		return &expr.Literal{Value: literal.Literal}, nil
	}

	ok, err = p.match([]uint{token.VariableType})
	if err != nil {
		return nil, err
	}
	if ok {
		variableExpression := p.previous()
		return &expr.Variable{Name: variableExpression.Lexeme}, nil
	}

	ok, err = p.match([]uint{token.OpenParenType})
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
		if peekToken.Type.Type() != token.EOFType {
			peekValue = peekToken.Lexeme
		}
		err = p.consume(token.CloseParenType, fmt.Sprintf("Expect ')' after expression but found %s", peekValue))
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
	if t.Type.Type() != token.EOFType {
		peekValue = t.Lexeme
	}
	return nil, fmt.Errorf("Expect Expression but found %s", peekValue)
}

func (p *Parser) getToken() (*token.Token, error) {
	t, err := p.lex.Next()
	if err == nil {
		if t.Type.Type() != token.EOFType {
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
	return t.Type.Type() == tokenType, nil
}

func (p *Parser) peek() (*token.Token, error) {
	return p.nextToken()
}
func (p *Parser) isAtEnd() (bool, error) {
	t, err := p.nextToken()
	if err != nil {
		return false, err
	}
	return t.Type.Type() == token.EOFType, nil
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
