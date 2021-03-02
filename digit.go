package lexer

import (
	"strings"
	"unicode"
)

// DigitState interprets the next runes as a digit. Digits could start with a
// mathematical notion indicating whether the upcoming value is a positive or
// negative number.
func DigitState(s *Scanner) StateFn {
	r := s.Next()
	mnotion := strings.ContainsRune(MathematicalNotion, r)

	if !unicode.IsDigit(r) && !mnotion {
		s.Errorf("invalid digit %s", string(r))
		return nil
	}

	if mnotion && !unicode.IsDigit(s.Peek()) {
		s.Errorf("invalid mathematical notion")
		return nil
	}

	defer s.Emit(TokenNumber)

	for {
		r := s.Next()
		if r == EOF {
			return nil
		}

		if unicode.IsDigit(r) || r == Dot {
			continue
		}

		s.Rewind()
		return IdleState
	}
}
