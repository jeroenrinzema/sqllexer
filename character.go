package lexer

import (
	"strings"
)

// CharacterState interperts the next runes as characters until a breakout
// character has been read. A check is performed to check whether the current
// characters are a token. If the characters do not match any token is a token
// identifier emitted.
func CharacterState(s *Scanner) StateFn {
	for {
		r := s.Next()
		if r == EOF {
			break
		}

		if strings.ContainsRune(Breakout, r) {
			s.Rewind()
			break
		}
	}

	// TODO(Jeroen): handle functions

	lower := strings.ToLower(s.Current())
	token, is := Keywords[lower]
	if is {
		s.Emit(token)
		return IdleState
	}

	s.Emit(TokenIdentifier)
	return IdleState
}
