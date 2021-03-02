package lexer

import "fmt"

// TokenType identifies the type of lex tokens
type TokenType int

// Tokens which could be emitted inside
const (
	TokenEOF TokenType = iota
	TokenError
	TokenIdentifier
	TokenCall

	TokenPlus     // +
	TokenMinus    // -
	TokenDivide   // /
	TokenMultiply // *

	TokenEquals  // =
	TokenNEquals // != || <>
	TokenGt      // >
	TokenGte     // >=
	TokenLt      // <
	TokenLte     // <=

	TokenString
	TokenNumber
	TokenBoolean

	TokenAsterisk     // *
	TokenLParentheses // (
	TokenRParentheses // )

	TokenKeywordSelect // select
	TokenKeywordAs     // as
	TokenKeywordTable  // table
	TokenKeywordFrom   // from
	TokenKeywordWhere  // where
	TokenKeywordAnd    // and
	TokenKeywordCreate // create
	TokenKeywordInsert // insert
	TokenKeywordInto   // into
	TokenKeywordValues // values
	TokenKeywordJoin   // join
	TokenKeywordLeft   // left
	TokenKeywordRight  // right
	TokenKeywordInt    // int
	TokenKeywordText   // text
)

// Token represents a lexer token
type Token struct {
	Type  TokenType // Type, such as TokenKeywordSelect
	Value string    // Value, such as "SELECT"
}

func (i Token) String() string {
	switch i.Type {
	case TokenEOF:
		return "EOF"
	}

	if len(i.Value) > 10 {
		return fmt.Sprintf("%.10q...", i.Value)
	}

	return fmt.Sprintf("%q", i.Value)
}
