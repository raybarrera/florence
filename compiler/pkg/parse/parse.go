package parse

import "fmt"

type scanner struct {
	Start                 int
	Current               int
	Line                  int
	EncounteredParseError bool
	Source                string
	Tokens                []Token
}

func run() {
	s := &scanner{}
	var tokens = s.scan()
	for i := range tokens {
		fmt.Printf("Token, %v", i)
	}
}

/// Parser implementation
func (s *scanner) scan() []Token {
	t := []Token{}
	for s.Current <= 10 {
		s.Start = s.Current
		s.scanTokens()
	}
	t = append(t, Token{})
	return t
}

func (s *scanner) scanTokens() {
	c := s.advance()
	switch string(c) {
	case LeftParen:
		addToken(LeftParen, s)
		break
	}
}

func (s *scanner) advance() rune {
	return []rune(s.Source)[s.Current+1]
}

func addToken(t TokenType, s *scanner) {
	addTokenLiteral(t, "", s)
}

func addTokenLiteral(t TokenType, literal string, s *scanner) {
	text := substr(s.Source, s.Start, s.Current)
	s.Tokens = append(s.Tokens, Token{
		Type:    t,
		Lexeme:  text,
		Literal: literal,
		Line:    s.Line,
	})
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
