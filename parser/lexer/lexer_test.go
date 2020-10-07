package lexer

import (
	"reflect"
	"testing"

	"github.com/5anthosh/eval/parser/token"
)

func TestLexerNormalExpression(t *testing.T) {
	expression := "34534 + 345.34 - 222 / 43435 * 745.234"
	lex := FromString(expression)
	tokens := []token.Token{
		{Type: token.Number, Literal: float64(34534), Lexeme: "34534", Column: 5},
		{Type: token.Plus, Literal: nil, Lexeme: "+", Column: 7},
		{Type: token.Number, Literal: float64(345.34), Lexeme: "345.34", Column: 14},
		{Type: token.Minus, Literal: nil, Lexeme: "-", Column: 16},
		{Type: token.Number, Literal: float64(222), Lexeme: "222", Column: 20},
		{Type: token.CommonSlash, Literal: nil, Lexeme: "/", Column: 22},
		{Type: token.Number, Literal: float64(43435), Lexeme: "43435", Column: 28},
		{Type: token.Star, Literal: nil, Lexeme: "*", Column: 30},
		{Type: token.Number, Literal: float64(745.234), Lexeme: "745.234", Column: 38},
	}
	for _, tt := range tokens {
		t.Run(tt.Lexeme, func(t *testing.T) {
			got, err := lex.Next()
			if err != nil {
				t.Errorf("Lexer.Next() error = %v, wantErr %v", err, false)
				return
			}
			if !reflect.DeepEqual(got, tt) {
				t.Errorf("Lexer.Next() = %v, want %v", got, tt)
			}
		})
	}
}
