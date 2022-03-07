package core

import (
	"bufio"
)

type Source struct {
	Reader *bufio.Reader
}

type Lexer struct {
	Source *Source
	C      rune
	Pos    *Position
}

type Position struct {
	Line   int
	Column int
}

type TokenKind int

const (
	EOF TokenKind = iota
	BAD
	WS
	NL
	IDENT
	KEYWORD
	STRING
)

var TokenString = []string{
	EOF:     "EOF",
	BAD:     "BAD",
	WS:      "WHITESPACE",
	NL:      "NEWLINE",
	KEYWORD: "KEYWORD",
	IDENT:   "IDENT",
	STRING:  "STRING",
}

var Keywords = []string{
	"DOCUMENT",
	"SECTION",
}

type Token struct {
	Pos   *Position
	Kind  TokenKind
	Value string
}
