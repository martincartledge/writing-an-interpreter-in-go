# Notes from Thorsten Ball's book, Writing an Interpreter in Go


## Lexing

### Lexical analysis

Source Code -> Tokens -> Abstract Syntax Tree

The first transformation from source code to tokens is called *lexical analysis* or *lexing*

This is done by a lexer (or tokenizer or scanner)

Tokens are small easily categorizable data structures that are fed to the parser

The parser turns the tokens into an Abstract Syntax Tree

Example input to a lexer:

```
"let x = 5  + 5;"
```

Example output: 

```
[
  LET,
  IDENTIFIER("x"),
  EQUAL_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```

In this case, whitespace does not show up as tokens, some languages, like in Python , whitespaces matter

A production ready lexer might also attach the line number, column number, and file name to a token. This makes reading error messages much easier

Example:

```
"error: expected semicolon token. line 42, column, 33, program.monkey"
```

### Defining tokens

For the `Token` data structure, the two fields needed are `type` (represents the type) and `literal` (contains the value of the token)

