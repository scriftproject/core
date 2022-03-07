package core

import (
	"bufio"
	"os"

	"github.com/scriftproject/core/helpers"
)

func ReadSource(filepath string) *Source {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		helpers.ErrLog.Fatalf(": %s\n", err)
	}

	return &Source{Reader: bufio.NewReader(f)}
}

func (s *Source) Lex() {
	// var out []*Token
	l := newLexer(s)
	for {
		token := l.nextToken()

		if token.Kind == EOF {
			break
		}

		helpers.InfoLog.Printf(":%d:%d: %v\t%s", token.Pos.Line, token.Pos.Column, TokenString[token.Kind], token.Value)
		// out = append(out, token)
	}
}
