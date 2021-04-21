package scanner

import (
	"unicode"
	
	"github.com/nkanaev/numb/pkg/token"
)

type Scanner struct {
	src []rune
	cur int

	Token token.Token
	Value string
}

func New(line string) *Scanner {
	return &Scanner{
		src: []rune(line),
		Token: token.Illegal,
	}
}

func (s *Scanner) char() rune {
	if s.cur >= len(s.src) {
		return 0
	}
	return s.src[s.cur]
}

func (s *Scanner) nextChar() {
	s.cur += 1
}

func (s *Scanner) next() {
	for ; unicode.IsSpace(s.char()); s.nextChar() {
	}
	ch := s.char()
	switch {
	case unicode.IsDigit(ch):
		start := s.cur
		for ; unicode.IsDigit(s.char()); s.nextChar() {
		}
		s.Token = token.NUM
		s.Value = string(s.src[start:s.cur])
	case ch == '(':
		s.Token = token.LPAREN
		s.nextChar()
	case ch == ')':
		s.Token = token.RPAREN
		s.nextChar()
	case ch == '*' || ch == '/' || ch == '+' || ch == '-':
		val := string(ch)
		s.Token = token.TokenString[val]
		s.Value = val
		s.nextChar()
	default:
		s.Token = token.Illegal
	}
}

func (s *Scanner) Scan() bool {
	s.next()
	return s.Token != token.Illegal
}
