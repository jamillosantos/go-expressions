// Generated from /home/jsantos/projects/aure-api/src/github.com/jamillosantos/go-expressions/antlr/Expression.g4 by ANTLR 4.7.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpressionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BaseExpressionListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BaseExpressionListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterPowExpression is called when production powExpression is entered.
func (s *BaseExpressionListener) EnterPowExpression(ctx *PowExpressionContext) {}

// ExitPowExpression is called when production powExpression is exited.
func (s *BaseExpressionListener) ExitPowExpression(ctx *PowExpressionContext) {}

// EnterSignedAtom is called when production signedAtom is entered.
func (s *BaseExpressionListener) EnterSignedAtom(ctx *SignedAtomContext) {}

// ExitSignedAtom is called when production signedAtom is exited.
func (s *BaseExpressionListener) ExitSignedAtom(ctx *SignedAtomContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BaseExpressionListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BaseExpressionListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseExpressionListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseExpressionListener) ExitAtom(ctx *AtomContext) {}

// EnterStr is called when production str is entered.
func (s *BaseExpressionListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseExpressionListener) ExitStr(ctx *StrContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BaseExpressionListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BaseExpressionListener) ExitScientific(ctx *ScientificContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseExpressionListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseExpressionListener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseExpressionListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseExpressionListener) ExitVariable(ctx *VariableContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseExpressionListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseExpressionListener) ExitFunction(ctx *FunctionContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaseExpressionListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaseExpressionListener) ExitRelop(ctx *RelopContext) {}
