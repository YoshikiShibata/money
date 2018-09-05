// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package lexer

import (
	"unicode"

	"github.com/YoshikiShibata/monkey/token"
)

var runeTokenMap = map[rune]token.TokenType{
	'=': token.ASSIGN,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	',': token.COMMA,
	'+': token.PLUS,
	'{': token.LBRACE,
	'}': token.RBRACE,
	0:   token.EOF,
}

type Lexer struct {
	input        []rune
	position     int // current position
	readPosition int // next input position
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespaces()

	tokenType, ok := runeTokenMap[l.ch]
	if !ok {
		if unicode.IsLetter(l.ch) {
			var tok token.Token
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if unicode.IsDigit(l.ch) {
			return token.Token{Type: token.INT,
				Literal: l.readNumber()}
		}

		return token.Token{Type: token.ILLEGAL,
			Literal: string([]rune{l.ch})}
	}
	defer l.readChar()
	if tokenType == token.EOF {
		return token.Token{Type: tokenType, Literal: ""}
	}

	return token.Token{Type: tokenType, Literal: string([]rune{l.ch})}
}

func (l *Lexer) skipWhitespaces() {
	for {
		switch l.ch {
		case ' ', '\t', '\n', '\r':
			l.readChar()
		default:
			return
		}
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}
