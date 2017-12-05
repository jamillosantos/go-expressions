package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"math"
)

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

			g.It("should solve an addition", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(1))
				v.Add("+", expressions.NewExpressionValue(1))
				Expect(v).NotTo(BeNil())
				r, err := v.Solve(ctx)
				Expect(err).To(BeNil())
				Expect(r).To(Equal(float64(2)))
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

		g.Describe("ExpressionBinary", func() {
			g.It("should create an expression", func() {
				v := expressions.NewExpressionBinary(expressions.NewExpressionValue(4), "==", expressions.NewExpressionValue(4))
				Expect(v).NotTo(BeNil())
			})

			g.It("should solve an addition", func() {
				v := expressions.NewExpressionMultiple()
				ctx := expressions.NewContext(nil, nil)
				v.Add("", expressions.NewExpressionValue(1))
				v.Add("+", expressions.NewExpressionValue(1))
				Expect(v).NotTo(BeNil())
				Expect(v.Solve(ctx)).To(Equal(float64(2)))
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
	})
}
