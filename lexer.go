package core

import (
	"fmt"
	"strings"

	"github.com/scriftproject/core/helpers"
)

func newLexer(s *Source) *Lexer {
	return &Lexer{
		Source: s,
		C:      '\u0000',
		Pos: &Position{
			Line:   1,
			Column: -1,
		},
	}
}

func (l *Lexer) posReset() {
	l.Pos.Column = 0
	l.Pos.Line++
}

func (l *Lexer) next() rune {
	r, _, err := l.Source.Reader.ReadRune()
	if err != nil {
		return '\u0000'
	}

	l.C = r

	l.Pos.Column++

	return r
}

func (l *Lexer) back() {
	err := l.Source.Reader.UnreadRune()
	if err != nil {
		helpers.ErrLog.Fatalf(":%d:%d: Cannot read", l.Pos.Line, l.Pos.Column)
	}

	l.Pos.Column--
}

func (l *Lexer) isDigit() bool {
	return '0' <= l.C && l.C <= '9'
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.C && l.C <= 'z' || 'A' <= l.C && l.C <= 'Z'
}

func (l *Lexer) isChar(r rune) bool {
	return l.C == r
}

func isKeyword(s string) bool {
	for _, k := range Keywords {
		if strings.EqualFold(k, s) {
			return true
		}
	}
	return false
}

func (l *Lexer) nextToken() *Token {
	switch l.next() {
	case '\n':
		var count int
		start := l.Pos
		for l.isChar('\n') {
			count++
			l.posReset()
			l.next()
		}

		l.back()

		return &Token{
			Pos:   start,
			Kind:  NL,
			Value: fmt.Sprintf("%d", count),
		}
	case '\u0000':
		return &Token{
			Kind:  EOF,
			Pos:   l.Pos,
			Value: "\u0000",
		}
	}

	// helpers.InfoLog.Printf(":%d:%d: %s", l.Pos.Line, l.Pos.Column, string(l.C))
	if l.isLetter() || l.isChar('_') {
		var lexme string
		start := l.Pos
		for l.isLetter() || l.isChar('_') || l.isDigit() {
			lexme += string(l.C)
			l.next()
		}
		l.back()

		if isKeyword(lexme) {
			return &Token{
				Pos:   start,
				Kind:  KEYWORD,
				Value: lexme,
			}
		}

		return &Token{
			Pos:   start,
			Kind:  IDENT,
			Value: lexme,
		}
	}

	if l.isChar('"') {
		var lexme string
		start := l.Pos
		l.next()
		for !l.isChar('"') {
			lexme += string(l.C)
			l.next()
		}

		return &Token{
			Pos:   start,
			Kind:  STRING,
			Value: lexme,
		}
	}

	if l.isChar('\u0020') {
		var count int
		start := l.Pos
		for l.isChar('\u0020') {
			count++
			l.next()
		}

		l.back()

		return &Token{
			Pos:   start,
			Kind:  WS,
			Value: fmt.Sprintf("%d", count),
		}
	}

	return &Token{
		Pos:   l.Pos,
		Kind:  BAD,
		Value: string(l.C),
	}
}
