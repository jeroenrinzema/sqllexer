package lexer

// Reserved symbols used to perform operations, group columns
// and other use cases.
const (
	Semicoln            = ';'
	Comma               = ','
	Space               = ' '
	Dot                 = '.'
	NewLine             = '\n'
	Tab                 = '\t'
	ParenthesesOpening  = '('
	ParenthesesClosing  = ')'
	AngleBracketOpening = '<'
	AngleBracketClosing = '>'
	RawQuote            = '`'
	DoubleQuote         = '"'
	SingleQuote         = '\''
	Equals              = '='
	ExclamationMark     = '!'
	Plus                = '+'
	Minus               = '-'
	Divide              = '/'
	Asterisk            = '*'
)

// Whitespace represents all whitespace characters.
// This value is often used to identify if a given run is a whitespace character
const Whitespace = string(NewLine) + string(Tab) + string(Space)

// MathmaticalOperators represents all available operator characters as a string.
// This value is often used to identify whether a operator character has been encountered.
const MathmaticalOperators = string(Plus) + string(Minus) + string(Divide) + string(Asterisk)

// CompareOperators represents all available operator characters as a string.
// This value is often used to identify whether a operator character has been encountered.
const CompareOperators = string(Equals) + string(ExclamationMark) + string(AngleBracketOpening) + string(AngleBracketClosing)

// MathematicalNotion represents characters that represents the mathmatical notion
// of a digit.
const MathematicalNotion = string(Plus) + string(Minus)

// Compare operators used to compare values
const Compare = string(Equals) + string(ExclamationMark)

// Breakout represents values often resulting into a state breakout.
const Breakout = string(Comma) + Whitespace + MathmaticalOperators + Compare + string(ParenthesesClosing) + string(Semicoln)
