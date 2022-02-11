package parse

import "fmt"

type scanner struct {
	Start                 int
	Current               int
	Line                  int
	EncounteredParseError bool
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
		getTokens()
	}
	t = append(t, Token{})
	return t
}

func getTokens() {

}
