package lexer

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	//bom                       = 0xFEFF // byte order mark
	TokenEof                  = -1
	TokenIdentifier TokenType = iota
	TokenKeyword
	TokenString
	TokenInt
	TokenFloat
	TokenArrow
	TokenDoubleColon
	TokenOpenBrace
	TokenCloseBrace
	TokenOpenParen
	TokenCloseParen
	TokenOpenBracket
	TokenCloseBracket
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input        string
	position     int
	nextPosition int
	currentChar  rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = rune(l.input[l.nextPosition])
	}
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.currentChar) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar) || l.currentChar == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	var builder strings.Builder

	for l.currentChar != '"' && l.currentChar != 0 {
		if l.currentChar == '\\' {
			l.readChar()
			switch l.currentChar {
			case '"':
				builder.WriteRune('"')
			case 'n':
				builder.WriteRune('\n')
			case 't':
				builder.WriteRune('\t')
			case 'r':
				builder.WriteRune('\r')
			case '\\':
				builder.WriteRune('\\')
				builder.WriteRune(l.currentChar)
			}
		} else {
			builder.WriteRune(l.currentChar)
		}
		l.readChar()
	}

	return builder.String()
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.currentChar) {
		l.readChar()
	}
	if l.currentChar == '.' && unicode.IsDigit(l.peekChar()) {
		l.readChar()
		for unicode.IsDigit(l.currentChar) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() rune {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.nextPosition])
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.currentChar {
	case 0:
		token.Type = TokenEof
	case '{':
		token = newToken(TokenOpenBrace, l.currentChar)
	case '}':
		token = newToken(TokenCloseBrace, l.currentChar)
	case '(':
		token = newToken(TokenOpenParen, l.currentChar)
	case ')':
		token = newToken(TokenCloseParen, l.currentChar)
	case '[':
		token = newToken(TokenOpenBracket, l.currentChar)
	case ']':
		token = newToken(TokenCloseBracket, l.currentChar)
	case '"':
		token.Type = TokenString
		token.Value = l.readString()
	default:
		if unicode.IsLetter(l.currentChar) || l.currentChar == '_' {
			identifier := l.readIdentifier()
			token.Type = lookUpIdentifier(identifier)
			token.Value = identifier
			return token
		} else if unicode.IsDigit(l.currentChar) {
			token.Type = TokenInt
			token.Value = l.readNumber()
			return token
		} else {
			token = newToken(TokenIdentifier, l.currentChar)
		}
	}
	l.readChar()
	return token

}

func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

func lookUpIdentifier(identifier string) TokenType {
	keywords := map[string]TokenType{
		"import":  TokenKeyword,
		"type":    TokenKeyword,
		"trait":   TokenKeyword,
		"fn":      TokenKeyword,
		"while":   TokenKeyword,
		"var":     TokenKeyword,
		"if":      TokenKeyword,
		"else":    TokenKeyword,
		"return":  TokenKeyword,
		"package": TokenKeyword,
	}
	if tokenType, ok := keywords[strings.ToLower(identifier)]; ok {
		return tokenType
	}
	return TokenIdentifier
}
