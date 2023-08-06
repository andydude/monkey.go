#!/bin/bash
# SPDX-FileCopyrightText: 2023 Andrew Robbins
# SPDX-License-Identifier: MIT

cat <<EOF
package token

type TokenType int

const (
	ILLEGAL   = -1
	EOF       = 0
EOF

fgrep "$(printf '\t')MonkeyLexer" monkey_lexer.go \
    | fgrep = | sed -e 's/MonkeyLexer//g' \
    | sed -E -e "s/INTEGER   =/INT       =/g" \
    | sed -E -e "s/EQ        =/ASSIGN    =/g" \
    | sed -E -e "s/EQ_EQ     =/EQ        =/g" \
    | sed -E -e "s/BANG_EQ   =/NOT_EQ    =/g" \
    | sed -E -e "s/STAR      =/ASTERISK  =/g" 

cat <<EOF
)

type Token struct {
	Type    TokenType
	Literal string
}
EOF
