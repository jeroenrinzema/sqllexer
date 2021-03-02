package lexer

import "testing"

func TestExpressions(t *testing.T) {
	t.Parallel()

	tests := []string{
		`SELECT (first + last), id FROM table`,
		`SELECT * FROM movements WHERE "type" = 'truck'`,
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			scanner := NewScanner(test, 0)
			go scanner.Run()

			for token := range scanner.Tokens() {
				if token.Type == TokenError {
					t.Fatal(token.Value)
				}
			}
		})
	}
}
