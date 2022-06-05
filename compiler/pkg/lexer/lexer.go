package lexer

import (
	"bufio"
	"io"
)

const (
	bom = 0xFEFF // byte order mark
	eof = -1     // end of file
)

type Lexer struct {
	Start   int
	Current int
	Line    int
	Column  int
	Tokens  []Token
	reader  *bufio.Reader

	ch         rune
	src        []byte
	insertSemi bool
	ErrorCount int
}

// NewLexer returns a pointer to a Lexer.
// reader is the reader of the file we are lexing
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		Line:   1,
		Column: 0,
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) next() {
	if l.Current < len(l.src) {
		l.Current++
	} else {
		if l.ch == '\n' {
			//...
		}
		l.ch = eof
	}
}

func (l *Lexer) Scan() Tok {
	return Tok{}
}

func (l *Lexer) peek() byte {
	if l.Current < len(l.src) {
		return l.src[l.Current]
	}
	return 0
}
