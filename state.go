package lexer

import (
	"strings"
	"unicode"
)

// IdleState represents the start state used as a starting point.
func IdleState(s *Scanner) StateFn {
	for {
		s.Take(Whitespace)
		s.Ignore()

		r := s.Next()
		if r == EOF || r == Semicoln {
			return nil
		}

		switch {
		case unicode.IsLetter(r):
			s.Rewind()
			return CharacterState
		case unicode.IsDigit(r) || (strings.ContainsRune(MathematicalNotion, r) && unicode.IsDigit(s.Peek())):
			s.Rewind()
			return DigitState
		case r == SingleQuote:
			s.Rewind()
			return SingleQuoteState
		case r == DoubleQuote:
			s.Rewind()
			return DoubleQuoteState
		case strings.ContainsRune(CompareOperators, r):
			s.Rewind()
			return CompareOperator
		case strings.ContainsRune(MathmaticalOperators, r):
			s.Rewind()
			return MathmaticalOperator
		case r == Comma:
			s.Emit(Comma)
		case r == ParenthesesOpening:
			s.Emit(TokenLParentheses)
		case r == ParenthesesClosing:
			s.Emit(TokenRParentheses)
		}
	}
}
