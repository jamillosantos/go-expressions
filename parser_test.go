package expressions_test

import (
	"testing"
	"github.com/jamillosantos/go-expressions"
)

func TestCompile_PlusAndMinus(t *testing.T) {
	resolver := expressions.NewMapResolver(map[string]float64{
		"t": float64(4),
	})
	expression, err := expressions.Compile("1 + 2 - 3 + 4 - 5 + 6")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(resolver)); v != float64(5) {
		t.Logf("5 expected %f got", v)
		t.FailNow()
	}
}

func TestCompile_PlusSequence(t *testing.T) {
	resolver := expressions.NewMapResolver(map[string]float64{
		"t": float64(4),
	})
	expression, err := expressions.Compile("1+1+2+2+3+3+4+4")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(resolver)); v != float64(20) {
		t.Logf("20 expected %f got", v)
		t.FailNow()
	}
}

func TestCompile_PlusSequenceWithBrackets(t *testing.T) {
	expression, err := expressions.Compile("1+(1+2)+2+3+(3+4)-4")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(nil)); v != float64(12) {
		t.Logf("12 expected %f got", v)
		t.FailNow()
	}
}

func TestCompile_Multiplication(t *testing.T) {
	expression, err := expressions.Compile("2*5 + 2*7*2")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(nil)); v != float64(38) {
		t.Logf("38 expected %f got", v)
		t.FailNow()
	}
}

func TestCompile_Division(t *testing.T) {
	expression, err := expressions.Compile("64/8 - 10/5")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(nil)); v != float64(6) {
		t.Logf("6 expected %f got", v)
		t.FailNow()
	}
}

func TestCompile_Pow(t *testing.T) {
	expression, err := expressions.Compile("2^3")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(nil)); v != float64(8) {
		t.Logf("8 expected %f got", v)
		t.FailNow()
	}
}

/*
func TestCompile_ManyPow(t *testing.T) {
	expression, err := expressions.Compile("2^2^3")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}
	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(nil)); v != float64(256) {
		t.Logf("256 expected %f got", v)
		t.FailNow()
	}
}
*/

func TestCompile_Variable(t *testing.T) {
	resolver := expressions.NewMapResolver(map[string]float64{
		"x": 4.5,
		"y": 2,
	})
	expression, err := expressions.Compile("(5+x)*y")
	if err != nil {
		t.Logf("Error: %s", err)
		t.Fail()
	}

	if expression == nil {
		t.Log("No expression was found.")
		t.FailNow()
	}
	if v := expression.Solve(expressions.NewContext(resolver)); v != float64(19) {
		t.Logf("19 expected %f got", v)
		t.FailNow()
	}
}
