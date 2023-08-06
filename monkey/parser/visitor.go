// SPDX-FileCopyrightText: 2023 Andrew Robbins
// SPDX-License-Identifier: MIT
package parser

import (
	"fmt"
	"strconv"
	"monkey/ast"
	"monkey/token"
	antlr "github.com/antlr4-go/antlr/v4"
)

type AstVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func NewAstVisitor() *AstVisitor {
	return &AstVisitor{}
}

func token2token(t antlr.Token) token.Token {
	return token.Token{
		Type: token.TokenType(t.GetTokenType()),
		Literal: t.GetText()}
}

func (v *AstVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	var stmts []ast.Statement
	var stmt ast.Statement
	
	for _, sctx := range ctx.AllStatement() {
		stmt = v.VisitStatement(
			sctx.(*StatementContext)).(ast.Statement)
		stmts = append(stmts, stmt)
	}

	var result any = &ast.Program{
		Statements: stmts}
	return result
}

func (v *AstVisitor) VisitStatement(ctx *StatementContext) interface{} {
	if ctx.BlockStatement() != nil {
		return v.VisitBlockStatement(
			ctx.BlockStatement(
			).(*BlockStatementContext))
	} else if ctx.LetStatement() != nil {
		return v.VisitLetStatement(
			ctx.LetStatement(
			).(*LetStatementContext))
	} else if ctx.ReturnStatement() != nil {
		return v.VisitReturnStatement(
			ctx.ReturnStatement(
			).(*ReturnStatementContext))
	} else if ctx.ExpressionStatement() != nil {
		return v.VisitExpressionStatement(
			ctx.ExpressionStatement(
			).(*ExpressionStatementContext))
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	var token = ctx.LBRACE().GetSymbol()
	var stmts []ast.Statement
	var stmt ast.Statement

	for _, sctx := range ctx.AllStatement() {
		stmt = v.VisitStatement(
			sctx.(*StatementContext)).(ast.Statement)
		stmts = append(stmts, stmt)
	}

	var result any = &ast.BlockStatement{
		Token: token2token(token),
		Statements: stmts}
	return result
}

func (v *AstVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	var token = ctx.LET().GetSymbol()
	var ident = v.VisitIdentifier(
		ctx.Identifier(
		).(*IdentifierContext)).(*ast.Identifier)
	var expr = v.VisitExpression(
		ctx.Expression(
		).(*ExpressionContext)).(ast.Expression)
	var result any = &ast.LetStatement{
		Token: token2token(token),
		Name: ident,
		Value: expr}
	return result
}

func (v *AstVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	var token = ctx.RETURN().GetSymbol()
	var expr = v.VisitExpression(
		ctx.Expression(
		).(*ExpressionContext)).(ast.Expression)
	var result any = &ast.ReturnStatement{
		Token: token2token(token),
		ReturnValue: expr}
	return result
}

func (v *AstVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	var token = ctx.SEMICOLON().GetSymbol()
	var expr = v.VisitExpression(
		ctx.Expression(
		).(*ExpressionContext)).(ast.Expression)
	var result any = &ast.ExpressionStatement{
		Token: token2token(token),
		Expression: expr}
	return result
}

func (v *AstVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitBinaryExpression(
		ctx.BinaryExpression(
		).(*BinaryExpressionContext))
}

func (v *AstVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	var token antlr.Token
	var oper string = ""
	if ctx.PrefixExpression() != nil {
		return v.VisitPrefixExpression(
			ctx.PrefixExpression(
			).(*PrefixExpressionContext))
	} else if ctx.EqualOperator() != nil {
		token = ctx.EqualOperator().GetStart()
		oper, _ = v.VisitEqualOperator(
			ctx.EqualOperator(
			).(*EqualOperatorContext)).(string)
	} else if ctx.RelOperator() != nil {
		token = ctx.RelOperator().GetStart()
		oper, _ = v.VisitRelOperator(
			ctx.RelOperator(
			).(*RelOperatorContext)).(string)
	} else if ctx.FactorOperator() != nil {
		token = ctx.FactorOperator().GetStart()
		oper, _ = v.VisitFactorOperator(
			ctx.FactorOperator(
			).(*FactorOperatorContext)).(string)
	} else if ctx.TermOperator() != nil {
		token = ctx.TermOperator().GetStart()
		oper, _ = v.VisitTermOperator(
			ctx.TermOperator(
			).(*TermOperatorContext)).(string)
	} else {
		panic("wtf")
	}

	var exprs = ctx.AllBinaryExpression()
	var left = v.VisitBinaryExpression(exprs[0].(*BinaryExpressionContext)).(ast.Expression)
	var right = v.VisitBinaryExpression(exprs[1].(*BinaryExpressionContext)).(ast.Expression)
	
	var result any = &ast.InfixExpression{
		Token: token2token(token),
		Left: left,
		Operator: oper,
		Right: right}
	return result
}

func (v *AstVisitor) VisitEqualOperator(ctx *EqualOperatorContext) interface{} {
	if ctx.EQ_EQ() != nil {
		return ctx.EQ_EQ().GetSymbol().GetText()
	} else if ctx.BANG_EQ() != nil {
		return ctx.BANG_EQ().GetSymbol().GetText()
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitRelOperator(ctx *RelOperatorContext) interface{} {
	if ctx.LT() != nil {
		return ctx.LT().GetSymbol().GetText()
	} else if ctx.GT() != nil {
		return ctx.GT().GetSymbol().GetText()
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitFactorOperator(ctx *FactorOperatorContext) interface{} {
	if ctx.STAR() != nil {
		return ctx.STAR().GetSymbol().GetText()
	} else if ctx.SLASH() != nil {
		return ctx.SLASH().GetSymbol().GetText()
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitTermOperator(ctx *TermOperatorContext) interface{} {
	if ctx.PLUS() != nil {
		return ctx.PLUS().GetSymbol().GetText()
	} else if ctx.MINUS() != nil {
		return ctx.MINUS().GetSymbol().GetText()
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	var result ast.Expression = (
		v.VisitPostfixExpression(
			ctx.PostfixExpression(
			).(*PostfixExpressionContext))).(ast.Expression)

	var nexprs = len(ctx.AllPrefixOperator())
	for i := nexprs - 1; i >= 0; i-- {
		var ectx = ctx.PrefixOperator(i)
		var token = ectx.GetStart()
		var oper string = v.VisitPrefixOperator(
			ectx.(*PrefixOperatorContext)).(string)
		result = &ast.PrefixExpression{
			Token: token2token(token),
			Operator: oper,
			Right: result}
	}

	return result
}

func (v *AstVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	var result ast.Expression = (
		v.VisitPrimaryExpression(
			ctx.PrimaryExpression(
			).(*PrimaryExpressionContext))).(ast.Expression)
	
	for _, ectx := range ctx.AllPostfixOperator() {
		var token = ectx.GetStart()
		if token.GetText() == "(" {
			var args = v.VisitPostfixOperator(
				ectx.(*PostfixOperatorContext)).(*ast.CallExpression)
			result = &ast.CallExpression{
				Token: token2token(token),
				Function: result,
				Arguments: args.Arguments}
		} else if token.GetText() == "[" {
			var expr ast.Expression = (
				v.VisitPostfixOperator(
					ectx.(*PostfixOperatorContext))).(ast.Expression)
			result = &ast.IndexExpression{
				Token: token2token(token),
				Left: result,
				Index: expr}
		} else {
			panic("wtf")
		}
	}

	return result
}

func (v *AstVisitor) VisitPrefixOperator(ctx *PrefixOperatorContext) interface{} {
	return ctx.GetStart().GetText()
}

func (v *AstVisitor) VisitPostfixOperator(ctx *PostfixOperatorContext) interface{} {
	if ctx.LPAREN() != nil {
		if ctx.ExpressionList() != nil {
			return v.VisitExpressionList(
				ctx.ExpressionList(
				).(*ExpressionListContext))
		} else {
			illegal_token := token.Token{
				Type: token.ILLEGAL,
				Literal: ""}
			var result = &ast.CallExpression{
				Token: illegal_token,
				Function: nil,
				Arguments: []ast.Expression{}}
			return result
		}
	} else if ctx.LBRACKET() != nil {
		return v.VisitExpression(
			ctx.Expression(
			).(*ExpressionContext))
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	if ctx.GroupExpression() != nil {
		return v.VisitGroupExpression(
			ctx.GroupExpression(
			).(*GroupExpressionContext))
	} else if ctx.FunctionLiteral() != nil {
		return v.VisitFunctionLiteral(
			ctx.FunctionLiteral(
			).(*FunctionLiteralContext))
	} else if ctx.IfExpression() != nil {
		return v.VisitIfExpression(
			ctx.IfExpression(
			).(*IfExpressionContext))
	} else if ctx.Identifier() != nil {
		return v.VisitIdentifier(
			ctx.Identifier(
			).(*IdentifierContext))
	} else if ctx.LiteralExpression() != nil {
		return v.VisitLiteralExpression(
			ctx.LiteralExpression(
			).(*LiteralExpressionContext))
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitGroupExpression(ctx *GroupExpressionContext) interface{} {
	return v.VisitExpression(
		ctx.Expression(
		).(*ExpressionContext))
}

func (v *AstVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	var token = ctx.IF().GetSymbol()
	var expr = v.VisitExpression(
		ctx.Expression().(*ExpressionContext)).(ast.Expression)
	var blocks = ctx.AllBlockStatement()
	if len(blocks) == 2 {
		var trueBlock = v.VisitBlockStatement(blocks[0].(*BlockStatementContext)).(*ast.BlockStatement)
		var falseBlock = v.VisitBlockStatement(blocks[1].(*BlockStatementContext)).(*ast.BlockStatement)
		var result any = &ast.IfExpression{
			Token: token2token(token),
			Condition: expr,
			Consequence: trueBlock,
			Alternative: falseBlock}
		return result
	} else if len(blocks) == 1 {
		var trueBlock = v.VisitBlockStatement(blocks[0].(*BlockStatementContext)).(*ast.BlockStatement)
		var result any = &ast.IfExpression{
			Token: token2token(token),
			Condition: expr,
			Consequence: trueBlock,
			Alternative: nil}
		return result
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	if ctx.BooleanLiteral() != nil {
		return v.VisitBooleanLiteral(
			ctx.BooleanLiteral(
			).(*BooleanLiteralContext))
	} else if ctx.IntegerLiteral() != nil {
		return v.VisitIntegerLiteral(
			ctx.IntegerLiteral(
			).(*IntegerLiteralContext))
	} else if ctx.StringLiteral() != nil {
		return v.VisitStringLiteral(
			ctx.StringLiteral(
			).(*StringLiteralContext))
	} else if ctx.ArrayListLiteral() != nil {
		return v.VisitArrayListLiteral(
			ctx.ArrayListLiteral(
			).(*ArrayListLiteralContext))
	} else if ctx.HashMapLiteral() != nil {
		return v.VisitHashMapLiteral(
			ctx.HashMapLiteral(
			).(*HashMapLiteralContext))
	} else {
		panic("wtf")
	}
}

func (v *AstVisitor) VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{} {
	var token2 = ctx.FUNCTION().GetSymbol()
	//var ictx = ctx.Identifier().(*IdentifierContext)
	//var ident *ast.Identifier = nil
	//if ictx != nil {
	//	ident = v.VisitIdentifier(ictx).(*ast.Identifier)
	//}

	illegal_token := token.Token{
		Type: token.ILLEGAL,
		Literal: ""}
	var sig = &ast.FunctionLiteral{
		Token: illegal_token,
		Parameters: []*ast.Identifier{},
		Body: nil}

	if ctx.IdentifierList() != nil {
		sig = (v.VisitIdentifierList(
			ctx.IdentifierList(
			).(*IdentifierListContext)).(*ast.FunctionLiteral))
	}
	
	var body = (v.VisitBlockStatement(
		ctx.BlockStatement(
		).(*BlockStatementContext)).(*ast.BlockStatement))
	var result = &ast.FunctionLiteral{
		Token: token2token(token2),
		//Name: ident,
		Parameters: sig.Parameters,
		Body: body}

	return result
}

func (v *AstVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	var token = ctx.GetStart()
	var value bool

	if ctx.TRUE() != nil {
		value = true
	} else if ctx.FALSE() != nil {
		value = false
	}
	
	var result any = &ast.Boolean{
		Token: token2token(token),
		Value: value}
	return result
}

func (v *AstVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	var token = ctx.INTEGER().GetSymbol()
	value, _ := strconv.ParseInt(token.GetText(), 10, 64)
	var result any = &ast.IntegerLiteral{
		Token: token2token(token),
		Value: value}
	return result
}

func (v *AstVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	var token = ctx.STRING().GetSymbol()
	var value = token.GetText()
	value = value[1:len(value)-1]
	token.SetText(value)
	var result any = &ast.StringLiteral{
		Token: token2token(token),
		Value: value}
	return result
}

func (v *AstVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	var token = ctx.IDENT().GetSymbol()
	var name = token.GetText()
	var result any = &ast.Identifier{
		Token: token2token(token),
		Value: name}
	return result
}

func (v *AstVisitor) VisitArrayListLiteral(ctx *ArrayListLiteralContext) interface{} {
	illegal_token := token.Token{
		Type: token.ILLEGAL,
		Literal: ""}
	var cexpr = &ast.CallExpression{
		Token: illegal_token,
		Arguments: []ast.Expression{}}
	if ctx.ExpressionList() != nil {
		cexpr = v.VisitExpressionList(
			ctx.ExpressionList(
			).(*ExpressionListContext)).(*ast.CallExpression)
	}
 	start_token := ctx.LBRACKET().GetSymbol()
	var result = &ast.ArrayLiteral{
		Token: token2token(start_token),
		Elements: cexpr.Arguments}
	return result
}

func (v *AstVisitor) VisitHashMapLiteral(ctx *HashMapLiteralContext) interface{} {
	start_token := ctx.LBRACE().GetSymbol()
	var pairs = make(map[ast.Expression]ast.Expression)
	var hexpr = &ast.HashLiteral{
		Token: token2token(start_token),
		Pairs: pairs}
	if ctx.HashMemberList() != nil {
		hexpr = v.VisitHashMemberList(ctx.HashMemberList().(*HashMemberListContext)).(*ast.HashLiteral)
	}
	var result = &ast.HashLiteral{
		Token: token2token(start_token),
		Pairs: hexpr.Pairs}
	return result
}

func (v *AstVisitor) VisitHashMemberList(ctx *HashMemberListContext) interface{} {
	var pairs = make(map[ast.Expression]ast.Expression)
	for _, pctx := range ctx.AllHashMember() {
		alit, ok := v.VisitHashMember(
			pctx.(*HashMemberContext)).(*ast.ArrayLiteral)
		if !ok {
			fmt.Printf("EXpected ArrayLiteral, got=%q", alit);
		}
		pairs[alit.Elements[0]] = alit.Elements[1]
	}
	illegal_token := token.Token{
		Type: token.ILLEGAL,
		Literal: ""}
 	var result = &ast.HashLiteral{
		Token: illegal_token,
		Pairs: pairs}
	return result
}

func (v *AstVisitor) VisitHashMember(ctx *HashMemberContext) interface{} {
	var token = ctx.COLON().GetSymbol()
	var result = &ast.ArrayLiteral{
		Token: token2token(token),
		Elements: []ast.Expression{
			v.VisitExpression(ctx.Expression(0).(*ExpressionContext)).(ast.Expression),
			v.VisitExpression(ctx.Expression(1).(*ExpressionContext)).(ast.Expression)}}

	return result
}

func (v *AstVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	var args []ast.Expression = []ast.Expression{}
	for _, ectx := range ctx.AllExpression() {
		ectx2, ok := ectx.(*ExpressionContext)
		if !ok {
			break
		}
		args = append(args, v.VisitExpression(ectx2).(ast.Expression))
	}

	// The purpose of this is to return an
	// array list of Expression objects,
	// but it impossible for those to implement
	// an interface, because only structs can implement
	// interfaces. So we wrap it up in an AST so
	// we can use the Arguments somewhere else.
	illegal_token := token.Token{
		Type: token.ILLEGAL,
		Literal: ""}
	var result = &ast.CallExpression{
		Token: illegal_token,
		Function: nil,
		Arguments: args}
	
	return result
}

func (v *AstVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	var params []*ast.Identifier = []*ast.Identifier{}
	for _, ectx := range ctx.AllIdentifier() {
		params = append(params, v.VisitIdentifier(
			ectx.(*IdentifierContext)).(*ast.Identifier))
	}
	
	// The purpose of this is to return an
	// array list of Identifier objects,
	// but it impossible for those to implement
	// an interface, because only structs can implement
	// interfaces. So we wrap it up in an AST so
	// we can use the Parameters somewhere else.
	illegal_token := token.Token{
		Type: token.ILLEGAL,
		Literal: ""}
	var result = &ast.FunctionLiteral{
		Token: illegal_token,
		Parameters: params,
		Body: nil}
	
	return result
}
