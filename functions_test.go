package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
	"math"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"fmt"
)

func TestFunctions(t *testing.T) {
	g := Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Default Functions", func() {
		g.Describe("'cos' function", func() {
			g.It("should calculate 'cos'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "cos", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Cos(0.1)))
			})

			g.It("should fail 'cos' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "cos")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("cos"))
			})

			g.It("should fail 'cos' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "cos", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("cos"))
			})
		})

		g.Describe("'cosh' function", func() {
			g.It("should calculate 'cosh'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "cosh", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Cosh(0.1)))
			})

			g.It("should fail 'cosh' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "cosh")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("cosh"))
			})

			g.It("should fail 'cosh' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "cosh", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("cosh"))
			})
		})

		g.Describe("'acos' function", func() {
			g.It("should calculate 'acos'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "acos", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Acos(0.1)))
			})

			g.It("should fail 'acos' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "acos")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("acos"))
			})

			g.It("should fail 'acos' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "acos", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("acos"))
			})
		})

		g.Describe("'acosh' function", func() {
			g.It("should calculate 'acosh'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "acosh", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Acosh(2)))
			})

			g.It("should fail 'acosh' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "acosh")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("acosh"))
			})

			g.It("should fail 'acosh' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "acosh", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("acosh"))
			})
		})

		g.Describe("'sin' function", func() {
			g.It("should calculate 'sin'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "sin", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sin(0.1)))
			})

			g.It("should fail 'sin' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sin")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sin"))
			})

			g.It("should fail 'sin' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sin", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sin"))
			})
		})

		g.Describe("'sinh' function", func() {
			g.It("should calculate 'sinh'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "sinh", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sinh(0.1)))
			})

			g.It("should fail 'sinh' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sinh")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sinh"))
			})

			g.It("should fail 'sinh' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sinh", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sinh"))
			})
		})

		g.Describe("'asin' function", func() {
			g.It("should calculate 'asin'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "asin", expressions.NewExpressionValue(0.1))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Asin(0.1)))
			})

			g.It("should fail 'asin' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "asin")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("asin"))
			})

			g.It("should fail 'asin' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "asin", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("asin"))
			})
		})

		g.Describe("'asinh' function", func() {
			g.It("should calculate 'asinh'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "asinh", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Asinh(2)))
			})

			g.It("should fail 'asinh' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "asinh")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("asinh"))
			})

			g.It("should fail 'asinh' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "asinh", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("asinh"))
			})
		})

		g.Describe("'sqrt' function", func() {
			g.It("should calculate 'sqrt'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "sqrt", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sqrt(2)))
			})

			g.It("should calculate 'sqrt'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "sqrt", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Sqrt(2)))
			})

			g.It("should calculate 'sqrt' with the second param", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "sqrt", expressions.NewExpressionValue(8), expressions.NewExpressionValue(3))
				Expect(err).To(BeNil())
				Expect(v).To(BeNumerically("~", float64(2), 0.00001))
			})

			g.It("should fail 'sqrt' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sqrt")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sqrt"))
			})

			g.It("should fail 'sqrt' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "sqrt", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("sqrt"))
			})
		})

		g.Describe("'tan' function", func() {
			g.It("should calculate 'tan'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "tan", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Tan(2)))
			})

			g.It("should fail 'tan' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "tan")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("tan"))
			})

			g.It("should fail 'tan' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "tan", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("tan"))
			})
		})

		g.Describe("'atan' function", func() {
			g.It("should calculate 'atan'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "atan", expressions.NewExpressionValue(2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atan(2)))
			})

			g.It("should fail 'atan' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atan")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atan"))
			})

			g.It("should fail 'atan' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atan", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atan"))
			})
		})

		g.Describe("'atan2' function", func() {
			g.It("should calculate 'atan2'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "atan2", expressions.NewExpressionValue(0.8), expressions.NewExpressionValue(1.2))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atan2(0.8, 1.2)))
			})

			g.It("should fail 'atan2' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atan2", expressions.NewExpressionValue(0.8))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atan2"))
			})

			g.It("should fail 'atan2' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atan2", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atan2"))
			})
		})

		g.Describe("'atanh' function", func() {
			g.It("should calculate 'atanh'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "atanh", expressions.NewExpressionValue(0.8))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Atanh(0.8)))
			})

			g.It("should fail 'atanh' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atanh")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atanh"))
			})

			g.It("should fail 'atanh' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "atanh", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("atanh"))
			})
		})

		g.Describe("'log' function", func() {
			g.It("should calculate 'log'", func() {
				functions := &expressions.DefaultFunctions{}
				v, err := functions.Call(expressions.NewContext(nil, nil), "log", expressions.NewExpressionValue(0.8))
				Expect(err).To(BeNil())
				Expect(v).To(Equal(math.Log(0.8)))
			})

			g.It("should fail 'log' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "log")
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("log"))
			})

			g.It("should fail 'log' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "log", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("log"))
			})
		})

		g.Describe("'if' function", func() {
			g.It("should calculate 'if'", func() {
				// TODO
				// functions := &expressions.DefaultFunctions{}
				// v, err := functions.Call(expressions.NewContext(nil, nil), "if", expressions.NewExpressionValue(0.8))
				// Expect(err).To(BeNil())
				// Expect(v).To(Equal(math.Log(0.8)))
			})

			g.It("should fail 'if' due to lack of parameters", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "if", expressions.NewExpressionValue(0.8), expressions.NewExpressionValue(0.8))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("if"))
			})

			g.It("should fail 'if' due to too many params", func() {
				functions := &expressions.DefaultFunctions{}
				_, err := functions.Call(expressions.NewContext(nil, nil), "if", expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1), expressions.NewExpressionValue(0.1))
				Expect(err).NotTo(BeNil())
				Expect(fmt.Sprint(err)).To(ContainSubstring("if"))
			})
		})
	})
}
