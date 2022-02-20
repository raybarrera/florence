package parse

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal string
	Line    int
}

func (t Token) Srting() string {
	return fmt.Sprintf("Type %v, Literal %v, Line %d", t.Type, t.Literal, t.Line)
}

const (
	EOF        = "EOF"
	Identifier = "Identifier"
	Type       = "Type"

	//Operators
	Assign      = "="
	Inference   = ":="
	Plus        = "+"
	Minus       = "-"
	Bang        = "!"
	Asterisk    = "*"
	Slash       = "/"
	Equality    = "=="
	NotEqual    = "!="
	LessThan    = "<"
	GreaterThan = ">"
	Scope       = "::"

	//Delimiters
	Comma     = ","
	Semicolon = ";"
	Colon     = ":"

	LeftParen    = "("
	RightParent  = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"

	//Ketwords
	True    = "True"
	False   = "False"
	If      = "If"
	Else    = "Else"
	Return  = "Return"
	Struct  = "Struct"
	Trait   = "Trait"
	Enum    = "Enum"
	Imports = "Imports"
)
