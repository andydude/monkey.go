// -*- mode: antlr; -*-
// SPDX-FileCopyrightText: 2023 Andrew Robbins
// SPDX-License-Identifier: MIT
lexer grammar MonkeyLexer;

FUNCTION
   : 'fn'
   ;

LET
   : 'let'
   ;

TRUE
   : 'true'
   ;

FALSE
   : 'false'
   ;

IF
   : 'if'
   ;

ELSE
   : 'else'
   ;

RETURN
   : 'return'
   ;

fragment ALPHA
   : [A-Za-z_]
   ;

IDENT
   : ALPHA+
   ;

fragment DIGIT
   : [0-9]
   ;

INTEGER
   : DIGIT+
   ;

fragment STRING_CHAR
   : ~ ["] // "
   ;

STRING
   : ["] STRING_CHAR* ["]
   ;

EQ_EQ
   : '=='
   ;

BANG_EQ
   : '!='
   ;

PLUS
   : '+'
   ;

MINUS
   : '-'
   ;

BANG
   : '!'
   ;

STAR
   : '*'
   ;

SLASH
   : '/'
   ;

LT
   : '<'
   ;

GT
   : '>'
   ;

EQ
   : '='
   ;

COMMA
   : ','
   ;

SEMICOLON
   : ';'
   ;

COLON
   : ':'
   ;

LPAREN
   : '('
   ;

RPAREN
   : ')'
   ;

LBRACE
   : '{'
   ;

RBRACE
   : '}'
   ;

LBRACKET
   : '['
   ;

RBRACKET
   : ']'
   ;

WS
   : [ \t\r\n]+ -> skip
   ;

