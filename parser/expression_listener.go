// Generated from /home/jsantos/projects/aure-api/src/github.com/jamillosantos/go-expressions/antlr/Expression.g4 by ANTLR 4.7.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterPowExpression is called when entering the powExpression production.
	EnterPowExpression(c *PowExpressionContext)

	// EnterSignedAtom is called when entering the signedAtom production.
	EnterSignedAtom(c *SignedAtomContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitPowExpression is called when exiting the powExpression production.
	ExitPowExpression(c *PowExpressionContext)

	// ExitSignedAtom is called when exiting the signedAtom production.
	ExitSignedAtom(c *SignedAtomContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
