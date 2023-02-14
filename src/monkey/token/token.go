package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y, ...
  INT   = "INT"   // 1343456

  // Operators
  ASSIGN = "="
  PLUS   = "+"

  COMMA     = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET      = "LET"
)

var keywords = map[string]TokenType{
  "fn":  FUNCTION,
  "let": LET,
}

// checks the keywords map to see if the identifier is a keyword
// if it is, return the keyword token type
// otherwise, return the IDENT token type
func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
