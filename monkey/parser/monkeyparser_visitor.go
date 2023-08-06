// Code generated from MonkeyParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MonkeyParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MonkeyParser.
type MonkeyParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MonkeyParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MonkeyParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#letStatement.
	VisitLetStatement(ctx *LetStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by MonkeyParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#equalOperator.
	VisitEqualOperator(ctx *EqualOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#relOperator.
	VisitRelOperator(ctx *RelOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#termOperator.
	VisitTermOperator(ctx *TermOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#factorOperator.
	VisitFactorOperator(ctx *FactorOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#prefixExpression.
	VisitPrefixExpression(ctx *PrefixExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#prefixOperator.
	VisitPrefixOperator(ctx *PrefixOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#postfixOperator.
	VisitPostfixOperator(ctx *PostfixOperatorContext) interface{}

	// Visit a parse tree produced by MonkeyParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#groupExpression.
	VisitGroupExpression(ctx *GroupExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#ifExpression.
	VisitIfExpression(ctx *IfExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#functionLiteral.
	VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by MonkeyParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayListLiteral.
	VisitArrayListLiteral(ctx *ArrayListLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashMapLiteral.
	VisitHashMapLiteral(ctx *HashMapLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashMemberList.
	VisitHashMemberList(ctx *HashMemberListContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashMember.
	VisitHashMember(ctx *HashMemberContext) interface{}
}
