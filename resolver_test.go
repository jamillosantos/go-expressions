package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMapResolver(t *testing.T) {
	g := Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Resolvers", func() {
		g.Describe("MapResolver", func() {
			g.It("should resolve two variables", func() {
				resolver := expressions.NewMapResolver(map[string]float64{
					"t": float64(1),
					"c": float64(2.34),
				})
				v, err := resolver.Resolve("t")
				Expect(err).To(BeNil())
				Expect(v).To(Equal(float64(1)))
				v, err = resolver.Resolve("c")
				Expect(err).To(BeNil())
				Expect(v).To(Equal(float64(2.34)))
			})
		})
	})
}
