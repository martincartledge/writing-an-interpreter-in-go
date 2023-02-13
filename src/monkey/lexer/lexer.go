package lexer

import "monkey/token"

// Lexer is the struct that will hold the input string and the current position of the Lexer
type Lexer struct {
  input        string
  position     int  // current position in input (points to current char)
  readPosition int  // current reading position in input (after current char)
  ch           byte // current char under examination
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}
// readChar reads the next character and advances the position of the input string
// another note: lexer only supports ASCII characters
// to support unicode, we would need to use the rune type instead of byte
// and change the way we read the input string because they would be multiple bytes wide
func (l *Lexer) readChar() {
  // check if we have reached the end of the input string
  if l.readPosition >= len(l.input) {
    // if so, set the character to 0 (ASCII code for NUL)
    // this signifies we have not read anything yet or we have reached the end of the file
    l.ch = 0
  } else {
    // if not, set the character to the next character in the input string
    l.ch = l.input[l.readPosition]
  }
  // update the values of the postion and readPosition
  l.position = l.readPosition
  l.readPosition += 1
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  switch l.ch { // switch on the current character
  case '=':
    tok = newToken(token.ASSIGN, l.ch)
  case ';':
    tok = newToken(token.SEMICOLON, l.ch)
  case '(':
    tok = newToken(token.LPAREN, l.ch)
  case ')':
    tok = newToken(token.RPAREN, l.ch)
  case ',':
    tok = newToken(token.COMMA, l.ch)
  case '+':
    tok = newToken(token.PLUS, l.ch)
  case '{':
    tok = newToken(token.LBRACE, l.ch)
  case '}':
    tok = newToken(token.RBRACE, l.ch)
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
  }
  l.readChar()
  // look at the current character and return the appropriate token
  return tok
}

// newToken is a helper function to create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

