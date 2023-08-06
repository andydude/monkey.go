// SPDX-FileCopyrightText: 2023 Andrew Robbins
// SPDX-License-Identifier: MIT
package parser

import (
	"monkey/ast"
	antlr "github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	l *Lexer
	p *MonkeyParser
	errors []string
}

func NewParser(l *Lexer) *Parser {
	tstream := antlr.NewCommonTokenStream(
		l.l, antlr.TokenDefaultChannel)
	p := NewMonkeyParser(tstream)
	return &Parser{l, p, []string{}}
}

func Parse(input string) *ast.Program {
	return NewParser(NewLexer(input)).ParseProgram()
}

func (p *Parser) ParseProgram() *ast.Program {
	ast_visitor := NewAstVisitor()
	tree := p.p.Program().Accept(ast_visitor)
	var program, ok = tree.(*ast.Program)
	if !ok {
		panic("wtf")
	}
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}
