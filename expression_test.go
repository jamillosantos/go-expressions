package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"math"
	"fmt"
	"errors"
)

type ExpressionFail struct {
}

func (*ExpressionFail) Solve(ctx expressions.Context) (interface{}, error) {
	return nil, errors.New("failed")
}

func TestExpressions(t *testing.T) {
	g := Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Expressions", func() {
		g.Describe("ExpressionValue", func() {
			g.It("should create an expression", func() {
				v := expressions.NewExpressionValue(1234)
				Expect(v).NotTo(BeNil())
			})

			g.It("should solve an expression", func() {
				v := expressions.NewExpressionValue(1234)
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(nil)).To(Equal(1234))
			})
		})

		g.Describe("ExpressionField", func() {
			g.It("should create an expression", func() {
				v := expressions.NewExpressionField("field1")
				Expect(v).NotTo(BeNil())
			})

			g.It("should solve an expression", func() {
				v := expressions.NewExpressionField("field")
				resolver := expressions.NewMapResolver(map[string]float64{
					"field": 1.2345,
				})
				ctx := expressions.NewContext(resolver, nil)
				Expect(v).NotTo(BeNil())
				Expect(resolver).NotTo(BeNil())
				Expect(ctx).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(1.2345)))
			})
		})

		g.Describe("ExpressionMultiple", func() {
			g.It("should create an expression", func() {
				v := expressions.NewExpressionMultiple()
				Expect(v).NotTo(BeNil())
			})

			g.It("should solve an addition of integers", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(int(1)))
				v.Add("+", expressions.NewExpressionValue(int(2)))
				Expect(v).NotTo(BeNil())
				r, err := v.Solve(ctx)
				Expect(err).To(BeNil())
				Expect(r).To(Equal(float64(3)))
			})

			g.It("should solve an addition of floats", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(float64(1)))
				v.Add("+", expressions.NewExpressionValue(float64(2)))
				Expect(v).NotTo(BeNil())
				r, err := v.Solve(ctx)
				Expect(err).To(BeNil())
				Expect(r).To(Equal(float64(3)))
			})

			g.It("should fail addition due to wrong data type (1st param)", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue("string cannot be added"))
				v.Add("+", expressions.NewExpressionValue(1))
				Expect(v).NotTo(BeNil())
				_, err := v.Solve(ctx)
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
			})

			g.It("should fail addition due to wrong data type (2nd param)", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(1))
				v.Add("+", expressions.NewExpressionValue("string cannot be added"))
				Expect(v).NotTo(BeNil())
				_, err := v.Solve(ctx)
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
			})

			g.It("should fail addition due to parameters failing (1st param)", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", &ExpressionFail{})
				v.Add("+", expressions.NewExpressionValue("string cannot be added"))
				Expect(v).NotTo(BeNil())
				_, err := v.Solve(ctx)
				Expect(err).NotTo(BeNil())
			})

			g.It("should fail addition due to parameters failing (2nd param)", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(1))
				v.Add("+", &ExpressionFail{})
				Expect(v).NotTo(BeNil())
				_, err := v.Solve(ctx)
				Expect(err).NotTo(BeNil())
			})

			g.It("should solve an subtraction", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(2))
				v.Add("-", expressions.NewExpressionValue(1))
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(1)))
			})

			g.It("should solve an multiplication", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(2))
				v.Add("*", expressions.NewExpressionValue(4))
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(8)))
			})

			g.It("should solve an division", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(16))
				v.Add("/", expressions.NewExpressionValue(2))
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(8)))
			})

			g.It("should solve an exponentiation", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(4))
				v.Add("^", expressions.NewExpressionValue(2))
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(16)))
			})
		})

		g.Describe("ExpressionMultiplePart", func() {
			g.It("should create an expression", func() {
				expr := expressions.NewExpressionMultiplePart("", nil)
				Expect(expr).NotTo(BeNil())
			})

			g.It("should solve an expression with empty operator", func() {
				expr := expressions.NewExpressionMultiplePart("", expressions.NewExpressionValue(1))
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, nil))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(1))
			})

			g.It("should fail due to expression solving failure injection", func() {
				expr := expressions.NewExpressionMultiplePart("", &ExpressionFail{})
				Expect(expr).NotTo(BeNil())
				_, err := expr.Solve(expressions.NewContext(nil, nil))
				Expect(err).NotTo(BeNil())
			})

			g.It("should fail due to invalid oeprator", func() {
				expr := expressions.NewExpressionMultiplePart("invalid operator", expressions.NewExpressionValue(4))
				Expect(expr).NotTo(BeNil())
				_, err := expr.Solve(expressions.NewContext(nil, nil))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("The operator"))
				Expect(fmt.Sprint(err)).To(ContainSubstring("invalid operator"))
				Expect(fmt.Sprint(err)).To(ContainSubstring("is not supported"))
			})

			g.Describe("Addition", func() {
				g.It("should solve an expression with integer params", func() {
					expr := expressions.NewExpressionMultiplePart("+", expressions.NewExpressionValue(1))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(4.5)))
				})

				g.It("should solve an expression with float params", func() {
					expr := expressions.NewExpressionMultiplePart("+", expressions.NewExpressionValue(float64(1.5)))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(5)))
				})

				g.It("should fail with injected expression solving failure", func() {
					expr := expressions.NewExpressionMultiplePart("+", &ExpressionFail{})
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with wrong data type", func() {
					expr := expressions.NewExpressionMultiplePart("+", expressions.NewExpressionValue("wrong data type"))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})
			})

			g.Describe("Subtraction", func() {
				g.It("should solve an expression with integer params", func() {
					expr := expressions.NewExpressionMultiplePart("-", expressions.NewExpressionValue(1))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(2.5)))
				})

				g.It("should solve an expression with float params", func() {
					expr := expressions.NewExpressionMultiplePart("-", expressions.NewExpressionValue(float64(1.5)))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(2)))
				})

				g.It("should fail with injected expression solving failure", func() {
					expr := expressions.NewExpressionMultiplePart("-", &ExpressionFail{})
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with wrong data type", func() {
					expr := expressions.NewExpressionMultiplePart("-", expressions.NewExpressionValue("wrong data type"))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})
			})

			g.Describe("Multiplication", func() {
				g.It("should solve an expression with integer params", func() {
					expr := expressions.NewExpressionMultiplePart("*", expressions.NewExpressionValue(2))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(7)))
				})

				g.It("should solve an expression with float params", func() {
					expr := expressions.NewExpressionMultiplePart("*", expressions.NewExpressionValue(float64(1.5)))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(4.5)))
				})

				g.It("should fail with injected expression solving failure", func() {
					expr := expressions.NewExpressionMultiplePart("*", &ExpressionFail{})
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with wrong data type", func() {
					expr := expressions.NewExpressionMultiplePart("*", expressions.NewExpressionValue("wrong data type"))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})
			})

			g.Describe("Division", func() {
				g.It("should solve an expression with integer params", func() {
					expr := expressions.NewExpressionMultiplePart("/", expressions.NewExpressionValue(2))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.6)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(1.8)))
				})

				g.It("should solve an expression with float params", func() {
					expr := expressions.NewExpressionMultiplePart("/", expressions.NewExpressionValue(float64(1.8)))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.6)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(2)))
				})

				g.It("should fail with injected expression solving failure", func() {
					expr := expressions.NewExpressionMultiplePart("/", &ExpressionFail{})
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with wrong data type", func() {
					expr := expressions.NewExpressionMultiplePart("/", expressions.NewExpressionValue("wrong data type"))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})
			})

			g.Describe("Power", func() {
				g.It("should solve an expression with integer params", func() {
					expr := expressions.NewExpressionMultiplePart("^", expressions.NewExpressionValue(2))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(4)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(Equal(float64(16)))
				})

				g.It("should solve an expression with float params", func() {
					expr := expressions.NewExpressionMultiplePart("^", expressions.NewExpressionValue(float64(1.8)))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.6)
					v, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(v).To(BeNumerically("~", 10.031006259, 0.0000001))
				})

				g.It("should fail with injected expression solving failure", func() {
					expr := expressions.NewExpressionMultiplePart("^", &ExpressionFail{})
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with wrong data type", func() {
					expr := expressions.NewExpressionMultiplePart("^", expressions.NewExpressionValue("wrong data type"))
					Expect(expr).NotTo(BeNil())
					ctx := expressions.NewContext(nil, nil)
					ctx.SetAccumulated(3.5)
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})
			})
		})

		g.Describe("ExpressionBinary", func() {
			g.It("should create an expression", func() {
				v := expressions.NewExpressionBinary(expressions.NewExpressionValue(4), "==", expressions.NewExpressionValue(4))
				Expect(v).NotTo(BeNil())
			})

			g.It("should fail due to invalid operator", func() {
				expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), "invalid operator", expressions.NewExpressionValue(3))
				ctx := expressions.NewContext(nil, nil)
				Expect(expr).NotTo(BeNil())
				_, err := expr.Solve(ctx)
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("The operator"))
				Expect(fmt.Sprint(err)).To(ContainSubstring("invalid operator"))
				Expect(fmt.Sprint(err)).To(ContainSubstring("is not supported"))
			})

			g.Describe("< (lower than)", func() {
				g.It("should solve true with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), "<", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), "<", expressions.NewExpressionValue(1))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.3), "<", expressions.NewExpressionValue(3.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3.3), "<", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should fail with injected expression solving error (1st param)", func() {
					expr := expressions.NewExpressionBinary(&ExpressionFail{}, "<", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with injected expression solving error (2nd param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "<", &ExpressionFail{})
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})
			})

			g.Describe("> (greater than)", func() {
				g.It("should solve true with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), ">", expressions.NewExpressionValue(1))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), ">", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3.3), ">", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.3), ">", expressions.NewExpressionValue(3.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should fail due to wrong data type (1st param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue("invalid"), ">", expressions.NewExpressionValue(3.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})

				g.It("should fail due to wrong data type (1st param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), ">", expressions.NewExpressionValue("invalid"))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(ContainSubstring("not a valid type"))
				})

				g.It("should fail with injected expression solving error (1st param)", func() {
					expr := expressions.NewExpressionBinary(&ExpressionFail{}, ">", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with injected expression solving error (2nd param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), ">", &ExpressionFail{})
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})
			})

			g.Describe("<= (lower than or equal)", func() {
				g.It("should solve true with integer parameters (lower)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), "<=", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve true with integer parameters (equal)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), "<=", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), "<=", expressions.NewExpressionValue(1))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with float parameters (lower)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "<=", expressions.NewExpressionValue(3.3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve true with float parameters (equal)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "<=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3.3), "<=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should fail with injected expression solving error (1st param)", func() {
					expr := expressions.NewExpressionBinary(&ExpressionFail{}, "<=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with injected expression solving error (2nd param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "<=", &ExpressionFail{})
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})
			})

			g.Describe(">= (greater than or equal)", func() {
				g.It("should solve true with integer parameters (greater)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), ">=", expressions.NewExpressionValue(1))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve true with integer parameters (equal)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), ">=", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), ">=", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with float parameters (greater)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3.3), ">=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve true with float parameters (equal)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), ">=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), ">=", expressions.NewExpressionValue(3.3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should fail with injected expression solving error (1st param)", func() {
					expr := expressions.NewExpressionBinary(&ExpressionFail{}, ">=", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with injected expression solving error (2nd param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), ">=", &ExpressionFail{})
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})
			})

			g.Describe("== (equal)", func() {
				g.It("should solve true with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(3), "==", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with integer parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1), "==", expressions.NewExpressionValue(3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "==", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with float parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "==", expressions.NewExpressionValue(3.3))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should solve true with string parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue("John Doe"), "==", expressions.NewExpressionValue("John Doe"))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeTrue())
				})

				g.It("should solve false with string parameters", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue("John Doe"), "==", expressions.NewExpressionValue("Not John Doe"))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					r, err := expr.Solve(ctx)
					Expect(err).To(BeNil())
					Expect(r).To(BeFalse())
				})

				g.It("should fail with injected expression solving error (1st param)", func() {
					expr := expressions.NewExpressionBinary(&ExpressionFail{}, "==", expressions.NewExpressionValue(1.5))
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})

				g.It("should fail with injected expression solving error (2nd param)", func() {
					expr := expressions.NewExpressionBinary(expressions.NewExpressionValue(1.5), "==", &ExpressionFail{})
					ctx := expressions.NewContext(nil, nil)
					Expect(expr).NotTo(BeNil())
					_, err := expr.Solve(ctx)
					Expect(err).NotTo(BeNil())
					Expect(fmt.Sprint(err)).To(Equal("failed"))
				})
			})
		})

		g.Describe("ExpressionFunction", func() {
			g.It("should create an expression", func() {
				expr := expressions.NewExpressionFunction("acos", expressions.NewExpressionValue(0.1))
				Expect(expr).NotTo(BeNil())
			})

			g.It("should solve a function", func() {
				expr := expressions.NewExpressionFunction("acos", expressions.NewExpressionValue(0.1))
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).NotTo(BeNil())
				Expect(v).To(Equal(math.Acos(0.1)))
			})
		})

		g.Describe("ExpressionBrackets", func() {
			g.It("should create an expression", func() {
				expr := expressions.NewExpressionBrackets(expressions.NewExpressionValue(0.1))
				Expect(expr).NotTo(BeNil())
			})

			g.It("should solve an integer expression", func() {
				expr := expressions.NewExpressionBrackets(expressions.NewExpressionValue(1))
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).NotTo(BeNil())
				Expect(v).To(Equal(1))
			})

			g.It("should solve an float expression", func() {
				expr := expressions.NewExpressionBrackets(expressions.NewExpressionValue(0.1))
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).NotTo(BeNil())
				Expect(v).To(Equal(0.1))
			})

			g.It("should solve an string expression", func() {
				expr := expressions.NewExpressionBrackets(expressions.NewExpressionValue("John Doe"))
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).NotTo(BeNil())
				Expect(v).To(Equal("John Doe"))
			})
		})
	})
}
