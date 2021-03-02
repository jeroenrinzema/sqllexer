package lexer

// MathmaticalOperator identifies the next rune as a mathmatical operator.
// If the read rune is a unknown operator is an error returned.
func MathmaticalOperator(s *Scanner) StateFn {
	r := s.Next()

	switch r {
	case Plus:
		s.Emit(TokenPlus)
		return IdleState
	case Minus:
		s.Emit(TokenMinus)
		return IdleState
	case Divide:
		s.Emit(TokenDivide)
		return IdleState
	case Asterisk:
		s.Emit(TokenMultiply)
		return IdleState
	}

	s.Errorf("unknown mathmatical operator %s", string(r))
	return nil
}

// CompareOperator reads the next rune(s) and attempts to identify them as a compare
// operator. If the next rune(s) are a unknown compare operator is an error returned.
func CompareOperator(s *Scanner) StateFn {
	c := s.Next()
	n := s.Next()

	// double character operators
	switch {
	case (c == ExclamationMark && n == Equals) || (c == AngleBracketOpening && n == AngleBracketClosing): // != || <>
		s.Emit(TokenNEquals)
		return IdleState
	case c == AngleBracketOpening && n == Equals: // <=
		s.Emit(TokenLte)
		return IdleState
	case c == AngleBracketClosing && n == Equals: // >=
		s.Emit(TokenGte)
		return IdleState
	}

	// single character operators
	s.Rewind()

	switch {
	case c == Equals: // =
		s.Emit(TokenEquals)
		return IdleState
	case c == AngleBracketOpening: // <
		s.Emit(TokenLt)
		return IdleState
	case c == AngleBracketClosing: // >
		s.Emit(TokenGt)
		return IdleState
	}

	s.Errorf("unknown compare operator %s", string(c))
	return nil
}
