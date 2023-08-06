// SPDX-FileCopyrightText: 2023 Andrew Robbins
// SPDX-License-Identifier: MIT
package token

type TokenType int

const (
	ILLEGAL   = -1
	EOF       = 0

	// Keywords
	FUNCTION  = 1
	LET       = 2
	TRUE      = 3
	FALSE     = 4
	IF        = 5
	ELSE      = 6
	RETURN    = 7
	IDENT     = 8
	INT       = 9
	STRING    = 10
	EQ        = 11
	NOT_EQ    = 12
	PLUS      = 13
	MINUS     = 14
	BANG      = 15
	ASTERISK  = 16
	SLASH     = 17
	LT 	  = 18
	GT 	  = 19
	ASSIGN    = 20

	// Delimiters
	COMMA     = 21
	SEMICOLON = 22
	COLON     = 23
	LPAREN    = 24
	RPAREN    = 25
	LBRACE    = 26
	RBRACE    = 27
	LBRACKET  = 28
	RBRACKET  = 29
)

type Token struct {
	Type    TokenType
	Literal string
}
