package factory

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}

	if compute(factory, 1, 2) != 3 {
		t.Fatal("加法错误")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("减法错误")
	}
}
