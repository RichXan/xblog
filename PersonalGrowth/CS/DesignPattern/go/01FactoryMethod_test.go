package designpattern

import (
	"fmt"
	"testing"
)

func TestPlusOperatorFactory(t *testing.T) {
	operator := CreateFactory(OperatorPlus)
	operator.SetA(1)
	operator.SetB(3)
	fmt.Print(operator.Result())
}

func TestMinusOperatorFactory(t *testing.T) {
	operator := CreateFactory(OperatorMinus)
	operator.SetA(1)
	operator.SetB(3)
	fmt.Print(operator.Result())
}
