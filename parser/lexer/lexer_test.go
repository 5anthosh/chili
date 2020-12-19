package lexer

import (
	"testing"

	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

func TestLexerNormalExpression(t *testing.T) {
	expression := "34534 + 345.34 - 222 / 43435 * 745.234 () () '3453453' \"apple is good$\" sum(8E-3,097e+34)"
	lex := FromString(expression)
	number1, _ := decimal.NewFromString("34534")
	number2, _ := decimal.NewFromString("345.34")
	number3, _ := decimal.NewFromString("222")
	number4, _ := decimal.NewFromString("43435")
	number5, _ := decimal.NewFromString("745.234")
	number6, _ := decimal.NewFromString("8E-3")
	number7, _ := decimal.NewFromString("097e+34")
	tokens := []token.Token{
		{Type: token.Number{}, Literal: number1, Lexeme: "34534", Column: 5},
		{Type: token.Plus{}, Literal: nil, Lexeme: "+", Column: 7},
		{Type: token.Number{}, Literal: number2, Lexeme: "345.34", Column: 14},
		{Type: token.Minus{}, Literal: nil, Lexeme: "-", Column: 16},
		{Type: token.Number{}, Literal: number3, Lexeme: "222", Column: 20},
		{Type: token.CommonSlash{}, Literal: nil, Lexeme: "/", Column: 22},
		{Type: token.Number{}, Literal: number4, Lexeme: "43435", Column: 28},
		{Type: token.Star{}, Literal: nil, Lexeme: "*", Column: 30},
		{Type: token.Number{}, Literal: number5, Lexeme: "745.234", Column: 38},
		{Type: token.OpenParen{}, Literal: nil, Lexeme: "(", Column: 40},
		{Type: token.CloseParen{}, Literal: nil, Lexeme: ")", Column: 41},
		{Type: token.OpenParen{}, Literal: nil, Lexeme: "(", Column: 43},
		{Type: token.CloseParen{}, Literal: nil, Lexeme: ")", Column: 44},
		{Type: token.LiteralString{}, Literal: "3453453", Lexeme: "'3453453'", Column: 54},
		{Type: token.LiteralString{}, Literal: "apple is good$", Lexeme: "\"apple is good$\"", Column: 71},
		{Type: token.Variable{}, Literal: nil, Lexeme: "sum", Column: 75},
		{Type: token.OpenParen{}, Literal: nil, Lexeme: "(", Column: 76},
		{Type: token.Number{}, Literal: number6, Lexeme: "8E-3", Column: 80},
		{Type: token.Comma{}, Literal: nil, Lexeme: ",", Column: 81},
		{Type: token.Number{}, Literal: number7, Lexeme: "097e+34", Column: 88},
		{Type: token.CloseParen{}, Literal: nil, Lexeme: ")", Column: 89},
	}
	for _, tt := range tokens {
		t.Run(tt.Lexeme, func(t *testing.T) {
			got, err := lex.Next()
			if err != nil {
				t.Errorf("Lexer.Next() error = %v, wantErr %v", err, false)
				return
			}
			testLocalToken(t, tt, *got)
		})
	}
}

func testLocalToken(t *testing.T, tt token.Token, got token.Token) {
	if tt.Column != got.Column {
		t.Errorf("Lexer.Next().Column == %v, want %v", got.Column, tt.Column)
	}

	if tt.Type.Type() != got.Type.Type() {
		t.Errorf("Lexer.Next().Type == %v, want %v", got.Type.String(), tt.Type.String())
	}

	switch tt.Literal.(type) {
	case decimal.Decimal:
		if !tt.Literal.(decimal.Decimal).Equal(got.Literal.(decimal.Decimal)) {
			t.Errorf("Lexer.Next().Literal == %v, want %v", got.Literal, tt.Literal)
		}
	}

	if tt.Lexeme != got.Lexeme {
		t.Errorf("Lexer.Next().Lexeme == %v, want %v", got.Lexeme, tt.Lexeme)
	}
}
