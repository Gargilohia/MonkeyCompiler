package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT     = "IDENT" // add, foobar, x, y, ..
	INT       = "INT"   // 123456
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	SLASH     = "/"
	ASTERISK  = "*"
	BANG      = "!"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
