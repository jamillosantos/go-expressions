package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"math"
)

func TestCompile(t *testing.T) {
	g := Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Compile", func() {
		g.It("should resolve a simple equation with additions and subtractions", func() {
			expr, err := expressions.Compile("1 + 2 - 3 + 4 - 5 + 6")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(5)))
		})

		g.It("should resolve a simple sequence of additions", func() {
			expr, err := expressions.Compile("1+1+2+2+3+3+4+4")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(20)))
		})

		g.It("should resolve a simple sequence of additions with brackets", func() {
			expr, err := expressions.Compile("1+(1+2)+2+3+(3+4)-4")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(12)))
		})

		g.It("should resolve a signed integers", func() {
			expr, err := expressions.Compile("1+(-3)")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(-2)))
		})

		g.It("should resolve a strange signed integers equation", func() {
			expr, err := expressions.Compile("1 + -3 + 4")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(2)))
		})

		g.It("should resolve a multiplication", func() {
			expr, err := expressions.Compile("2*5 + 2*7*2")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(38)))
		})

		g.It("should resolve a multiplication", func() {
			expr, err := expressions.Compile("2*5 + 2*7*2")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(38)))
		})

		g.It("should resolve a division", func() {
			expr, err := expressions.Compile("64/8 - 10/5")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(6)))
		})

		g.It("should resolve a pow", func() {
			expr, err := expressions.Compile("2^3")
			Expect(err).To(BeNil())
			Expect(expr).NotTo(BeNil())
			v, err := expr.Solve(expressions.NewContext(nil, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(8)))
		})

		g.It("should resolve a equation with variables", func() {
			resolver := expressions.NewMapResolver(map[string]float64{
				"x": 4.5,
				"y": 2,
			})
			expr, err := expressions.Compile("(5+x)*y")
			if err != nil {
				t.Logf("Error: %s", err)
				t.Fail()
			}
			v, err := expr.Solve(expressions.NewContext(resolver, nil))
			Expect(err).To(BeNil())
			Expect(v).To(Equal(float64(19)))
		})

		g.Describe("Functions", func() {
			g.It("should resolve a function 'acos'", func() {
				expr, err := expressions.Compile("cos(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Cos(0.1)))
			})

			g.It("should resolve a function 'cosh'", func() {
				expr, err := expressions.Compile("cosh(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Cosh(0.1)))
			})

			g.It("should resolve a function 'acos'", func() {
				expr, err := expressions.Compile("acos(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Acos(0.1)))
			})

			g.It("should resolve a function 'acosh'", func() {
				expr, err := expressions.Compile("acosh(2)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Acosh(2)))
			})

			g.It("should resolve a function 'sin'", func() {
				expr, err := expressions.Compile("sin(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sin(0.1)))
			})

			g.It("should resolve a function 'sinh'", func() {
				expr, err := expressions.Compile("sinh(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sinh(0.1)))
			})

			g.It("should resolve a function 'asin'", func() {
				expr, err := expressions.Compile("asin(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Asin(0.1)))
			})

			g.It("should resolve a function 'asinh'", func() {
				expr, err := expressions.Compile("asinh(2)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Asinh(2)))
			})

			g.It("should resolve a function 'sqrt'", func() {
				expr, err := expressions.Compile("sqrt(16)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(float64(4)))
			})

			g.It("should resolve a function 'sqrt' with the second param", func() {
				expr, err := expressions.Compile("sqrt(8, 3)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(BeNumerically("~", float64(2), 0.0001))
			})

			g.It("should resolve a function 'tan'", func() {
				expr, err := expressions.Compile("tan(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Tan(0.1)))
			})

			g.It("should resolve a function 'atan'", func() {
				expr, err := expressions.Compile("atan(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atan(0.1)))
			})

			g.It("should resolve a function 'atan2'", func() {
				expr, err := expressions.Compile("atan2(0.1, 0.2)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atan2(0.1, 0.2)))
			})

			g.It("should resolve a function 'atanh'", func() {
				expr, err := expressions.Compile("atanh(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atanh(0.1)))
			})

			g.It("should resolve a function 'log'", func() {
				expr, err := expressions.Compile("log(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Log(0.1)))
			})

			g.It("should resolve a function 'if'", func() {
				// TODO!!
				/*
				expr, err := expressions.Compile("if(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				v, err := expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Log(0.1)))
				*/
			})

			g.It("should fail resolving an unknown function 'unexistent'", func() {
				expr, err := expressions.Compile("unexistent(0.1)")
				Expect(err).To(BeNil())
				Expect(expr).NotTo(BeNil())
				_, err = expr.Solve(expressions.NewContext(nil, &expressions.DefaultFunctions{}))
				Expect(err).NotTo(BeNil())
			})
		})
	})
}
