package lexer

// DoubleQuoteState interprets the next runes as a double quoted string.
func DoubleQuoteState(s *Scanner) StateFn {
	r := s.Next()
	if r != DoubleQuote {
		s.Errorf("unexpected character expected a double quote")
		return nil
	}

	err := s.UntilRune(DoubleQuote)
	if err != nil {
		s.Errorf("unclosed double quote")
		return nil
	}

	s.Emit(TokenIdentifier)
	return IdleState
}

// SingleQuoteState interprets the next runes as a single quoted string.
func SingleQuoteState(s *Scanner) StateFn {
	r := s.Next()
	if r != SingleQuote {
		s.Errorf("unexpected character expected a single quote")
		return nil
	}

	err := s.UntilRune(SingleQuote)
	if err != nil {
		s.Errorf("unclosed single quote")
		return nil
	}

	s.Emit(TokenIdentifier)
	return IdleState
}
