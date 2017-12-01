package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
)

func TestMapResolver_Resolve(t *testing.T) {
	resolver := expressions.NewMapResolver(map[string]float64{
		"t": float64(1),
		"c": float64(2.34),
	})
	if v := resolver.Resolve("t"); v != 1 {
		t.Errorf("1 expected %f", v)
	}
	if v := resolver.Resolve("c"); v != 2.34 {
		t.Errorf("2.34 expected %f", v)
	}
}
