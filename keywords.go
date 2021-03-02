package lexer

// Reserved SQL keywords used
const (
	SelectKeyword = "select"
	FromKeyword   = "from"
	AsKeyword     = "as"
	TableKeyword  = "table"
	WhereKeyword  = "where"
	AndKeyword    = "and"
	CreateKeyword = "create"
	InsertKeyword = "insert"
	IntoKeyword   = "into"
	ValuesKeyword = "values"
	LeftKeyword   = "left"
	RightKeyword  = "right"
	IntKeyword    = "int"
	TextKeyword   = "text"
)

// Keywords represents a list of keywords and their token types
var Keywords = map[string]TokenType{
	SelectKeyword: TokenKeywordSelect,
	FromKeyword:   TokenKeywordFrom,
	AsKeyword:     TokenKeywordAs,
	TableKeyword:  TokenKeywordTable,
	WhereKeyword:  TokenKeywordWhere,
	AndKeyword:    TokenKeywordAnd,
	CreateKeyword: TokenKeywordCreate,
	InsertKeyword: TokenKeywordInsert,
	IntoKeyword:   TokenKeywordInto,
	ValuesKeyword: TokenKeywordValues,
	LeftKeyword:   TokenKeywordLeft,
	RightKeyword:  TokenKeywordRight,
	IntKeyword:    TokenKeywordInt,
	TextKeyword:   TokenKeywordText,
}
