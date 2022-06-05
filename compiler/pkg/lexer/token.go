package lexer

import (
	"fmt"
	"strconv"
)

type TokenType string

type Tok struct {
	Type    TokenType
	Lexeme  string
	Literal string
	Line    int
}

func (t Tok) String() string {
	return fmt.Sprintf("Type %v, Literal %v, Line %d", t.Type, t.Literal, t.Line)
}

type Token int

const (
	ILLEGAL Token = iota
	EOF           // End of File
	COMMENT

	IDENT  // main
	INT    // 1234
	FLOAT  // 1.25
	CHAR   // 'a'
	STRING // "abc"

	// Infix Operators
	ADD      // +
	SUB      // -
	MUL      // *
	DIV      // /
	MOD      // %
	AND      // &
	OR       // |
	XOR      // ^
	SHL      // <<
	SHR      // >>
	LAND     // &&
	LOR      // ||
	LARROW   // <-
	RARROW   // ->
	EQL      // ==
	LESS     // <
	GREATER  // <
	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	NOT        // !
	ASSIGN     // =
	INC        // ++
	DEC        // --
	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	DIV_ASSIGN // /=
	MOD_ASSIGN // %=
	SCOPE      // ::

	delimiterBegin
	LPAREN   // (
	LBRACKET // [
	LBRACE   // {
	COMMA    // ,
	PERIOD   // .

	RPAREN    // )
	RBRACKET  // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	delimiterEnd

	keywordBegin
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	STRUCT
	TRAIT
	CASE
	SWITCH
	IMPORT
	FN
	PACKAGE
	PROGRAM
	IT
	keywordEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	MOD: "%",
}

func (tok Token) String() string {
	s := ""
	if tok >= 0 && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
