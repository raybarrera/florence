package parse

import (
	"bytes"
	"compiler/pkg/lexer"
	"fmt"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type scanner struct {
	Start                 int
	Current               int
	Line                  int
	Column                int
	EncounteredParseError bool
	Source                string
	Tokens                []lexer.Token

	err        ErrorHandler // error reporting; or nil
	src        []byte
	ch         rune //Current character
	offset     int  //character offset
	rdOffset   int  //reading offset
	lineOffset int  //current line offset
	insertSemi bool //insert semicolon before next newline?
	ErrorCount int  //number of errors encountered by the scanner
}

const (
	bom = 0xFEFF // byte order mark
	eof = -1     // end of file
)

type ErrorHandler func(pos token.Position, msg string)

func (s *scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		if s.ch == '\n' {
			s.lineOffset = s.offset
			//Add line to file?
		}
		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		case r == 0:
			// Error here

		case r >= utf8.RuneSelf:
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				// error here too -- illegal utf8 encoding
			} else if r == bom && s.offset > 0 {
				// error here too -- illegal bom (should only appear at beninning)
			}
			s.rdOffset += w
			s.ch = r
		}
	} else {
		// we've reached the end of the file.
		s.offset = len(s.src)
		if s.ch == '\n' {
			s.lineOffset = s.offset
			// file add line here
		}
		s.ch = eof
	}
}

func (s *scanner) peek() byte {
	if s.rdOffset < len(s.src) {
		return s.src[s.rdOffset]
	}
	return 0
}

type Mode uint

const (
	ScanComments    Mode = 1 << iota // return comments as tokens (of type COMMENT)
	dontInsertSemis                  // do not automatically insert semicolons -- for testing only
)

func (s *scanner) error(offs int, msg string) {
	if s.err != nil {
		// handle the error here. look at line 153 on go scanner.
	}
	s.ErrorCount++
}

func (s *scanner) errorf(offs int, format string, args ...interface{}) {
	s.error(offs, fmt.Sprintf(format, args...))
}

func (s *scanner) scanComment() string {
	// initial '/' already consumed; s.ch == '/' || s.ch == '*'
	offs := s.offset - 1
	next := -1
	numCR := 0

	if s.ch == '/' {
		s.next()
		for s.ch != '\n' && s.ch >= 0 {
			if s.ch == '\r' {
				numCR++
			}
			s.next()
		}
		next = s.offset
		if s.ch == '\n' {
			next++
		}
		goto exit
	}
	s.next()
	for s.ch >= 0 {
		ch := s.ch
		if ch == '\r' {
			numCR++
		}
		s.next()
		if ch == '*' && s.ch == '/' {
			s.next()
			next = s.offset
			goto exit
		}
	}
exit:
	lit := s.src[offs:s.offset]

	// On Windows, a (//-comment) line may end in "\r\n".
	// Remove the final '\r' before analyzing the text for
	// line directives (matching the compiler). Remove any
	// other '\r' afterwards (matching the pre-existing be-
	// havior of the scanner).
	if numCR > 0 && len(lit) >= 2 && lit[1] == '/' && lit[len(lit)-1] == '\r' {
		lit = lit[:len(lit)-1]
		numCR--
	}

	// interpret line directives
	// (//line directives must start at the beginning of the current line)
	if next >= 0 /* implies valid comment */ && (lit[1] == '*' || offs == s.lineOffset) && bytes.HasPrefix(lit[2:], prefix) {
		s.updateLineInfo(next, offs, lit)
	}

	if numCR > 0 {
		lit = stripCR(lit, lit[1] == '*')
	}

	return string(lit)
}

var prefix = []byte("line ")

func (s *scanner) updateLineInfo(next, offs int, text []byte) {
	// extract comment text
	if text[1] == '*' {
		text = text[:len(text)-2]
	}
	text = text[7:]
	offs += 7

	i, n, ok := trailingDigits(text)

	if i == 0 {
		return
	}

	if !ok {
		//error
		return
	}

	var line, col int
	i2, n2, ok2 := trailingDigits(text[:i-1])
	if ok2 {
		i, i2 = i2, i
		line, col = n2, n
		if col == 0 {
			// error
			return
		}
		text = text[:i2-1]
	} else {
		line = n
	}

	if line == 0 {
		// error
		return
	}
	//TODO the files stuff
}

func trailingDigits(text []byte) (int, int, bool) {
	i := bytes.LastIndexByte(text, ':')
	if i < 0 {
		return 0, 0, false // no colon (":") found
	}
	n, err := strconv.ParseUint(string(text[i+1]), 10, 0)
	return i + 1, int(n), err == nil
}

func (s *scanner) findLineEnd() bool {
	defer func(offs int) {
		s.ch = '/'
		s.offset = offs
		s.rdOffset = offs + 1
		s.next()
	}(s.offset - 1)

	for s.ch == '/' || s.ch == '*' {
		if s.ch == '/' {
			return true
		}
		s.next()
		for s.ch >= 0 {
			ch := s.ch
			if ch == '\n' {
				return true
			}
			s.next()
			if ch == '*' && s.ch == '/' {
				s.next()
				break
			}
		}
		s.skipWhitespace()
		if s.ch < 0 || s.ch == '\n' {
			return true
		}
		if s.ch != '/' {
			return false
		}
		s.next()
	}
	return false
}

func (s *scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' && !s.insertSemi || s.ch == '\r' {
		s.next()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func lower(ch rune) rune {
	return ('a' - 'A' | ch)
}

func isDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isHex(ch rune) bool {
	return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f'
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
