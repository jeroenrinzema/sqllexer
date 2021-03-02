package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// ErrUnexpectedEOF is returned when a unexpected EOF is encountered
var ErrUnexpectedEOF = errors.New("unexpected EOF")

const (
	// EOF represents the end of a
	EOF rune = -1
	// DefaultBuffer ensures that the default buffer size (len(source) / 2) is applied.
	DefaultBuffer = -1
)

// NewScanner constructs a new scanner for the given source
func NewScanner(source string, buffer int) *Scanner {
	if buffer < 0 {
		buffer = len(source) / 2
		if buffer <= 0 {
			buffer = 1
		}
	}

	return &Scanner{
		source: source,
		tokens: make(chan Token, buffer),
	}
}

// Scanner holds the state of the scanner.
type Scanner struct {
	source   string     // the string being scanned.
	start    int        // start position of this item.
	position int        // current position in the input.
	tokens   chan Token // channel of scanned tokens
	rewind   runeStack  // rewind contains previously scanned runes
}

// StateFn represents the state of the scanner
// as a function that returns the next state.
type StateFn func(*Scanner) StateFn

// Run lexes the input by executing state functions until
// the state is nil.
func (s *Scanner) Run() {
	for state := IdleState; state != nil; {
		state = state(s)
	}

	close(s.tokens) // No more tokens will be delivered.
}

// Current returns the value being being analyzed at this moment.
func (s *Scanner) Current() string {
	return s.source[s.start:s.position]
}

// Emit will receive a token type and push a new token with the current analyzed
// value into the tokens channel.
func (s *Scanner) Emit(t TokenType) {
	tok := Token{
		Type:  t,
		Value: s.Current(),
	}
	s.tokens <- tok
	s.start = s.position
	s.rewind.clear()
}

// Ignore clears the rewind stack and then sets the current beginning position
// to the current position in the source which effectively ignores the section
// of the source being analyzed.
func (s *Scanner) Ignore() {
	s.rewind.clear()
	s.start = s.position
}

// Peek performs a Next operation immediately followed by a Rewind returning the
// peeked rune.
func (s *Scanner) Peek() rune {
	r := s.Next()
	s.Rewind()

	return r
}

// Rewind will take the last rune read (if any) and rewind back. Rewinds can
// occur more than once per call to Next but you can never rewind past the
// last point a token was emitted.
func (s *Scanner) Rewind() {
	r := s.rewind.pop()
	if r > EOF {
		size := utf8.RuneLen(r)
		s.position -= size
		if s.position < s.start {
			s.position = s.start
		}
	}
}

// Next pulls the next rune from the Lexer and returns it, moving the position
// forward in the source.
func (s *Scanner) Next() rune {
	var (
		r    rune
		size int
	)
	str := s.source[s.position:]
	if len(str) == 0 {
		r, size = EOF, 0
	} else {
		r, size = utf8.DecodeRuneInString(str)
	}
	s.position += size
	s.rewind.push(r)

	return r
}

// Errorf emits an error token with the given string format and arguments. A nil
// state function is returns aborting the scanner.
func (s *Scanner) Errorf(format string, args ...interface{}) StateFn {
	s.tokens <- Token{
		Type:  TokenError,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

// EOF indicates whether the end of file has been reached for the given source
func (s *Scanner) EOF() bool {
	return len(s.source[s.position:]) == 0
}

// Take receives a string containing all acceptable strings and will contine
// over each consecutive character in the source until a token not in the given
// string is encountered. This should be used to quickly pull token parts.
func (s *Scanner) Take(chars string) {
	r := s.Next()
	for strings.ContainsRune(chars, r) && r != EOF {
		r = s.Next()
	}

	if r == EOF {
		return
	}

	s.Rewind() // last next wasn't a match
}

// UntilRune reads the next runes until the breakout character has been read or the EOF
// has been encountered. This should be used to quickly read a unknown amount of
// characters until the breakout character has been read.
func (s *Scanner) UntilRune(b rune) error {
	r := s.Next()
	// TODO(Jeroen): ignore the next character if a backslash is encountered
	for r != b && r != EOF {
		r = s.Next()
	}

	if r == EOF {
		return ErrUnexpectedEOF
	}

	return nil
}

// Tokens returns a read only channel receiving tokens emitted by the scanner
func (s *Scanner) Tokens() <-chan Token {
	return s.tokens
}
