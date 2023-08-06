// SPDX-FileCopyrightText: 2023 Andrew Robbins
// SPDX-License-Identifier: MIT
package parser

import (
	"monkey/token"
	antlr "github.com/antlr4-go/antlr/v4"
)

type Lexer struct {
	l *MonkeyLexer
}

func NewLexer(input string) *Lexer {
	istream := antlr.NewInputStream(input)
	lexer := NewMonkeyLexer(istream)
	return &Lexer{lexer}
}

func (l *Lexer) NextToken() token.Token {
	antlrtoken := l.l.NextToken()
	tokentype := antlrtoken.GetTokenType()
	text := antlrtoken.GetText()
	if tokentype == antlr.TokenEOF {
		return token.Token{
			Type: token.EOF,
			Literal: ""}
	}
	if tokentype == token.STRING {
		text = text[1:len(text)-1]
	}
	
	return token.Token{
		Type: token.TokenType(tokentype),
		Literal: text}
}
