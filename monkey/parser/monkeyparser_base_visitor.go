// Code generated from MonkeyParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MonkeyParser

import "github.com/antlr4-go/antlr/v4"

type BaseMonkeyParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMonkeyParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitEqualOperator(ctx *EqualOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitRelOperator(ctx *RelOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFactorOperator(ctx *FactorOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitTermOperator(ctx *TermOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPrefixOperator(ctx *PrefixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPostfixOperator(ctx *PostfixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitGroupExpression(ctx *GroupExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayListLiteral(ctx *ArrayListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHashMapLiteral(ctx *HashMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHashMapMember(ctx *HashMapMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}
