package designpattern

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	// 数学表达式示例: (5 + 3) - 2
	five := NewNumberExpression(5)
	three := NewNumberExpression(3)
	add := NewAddExpression(five, three)
	two := NewNumberExpression(2)
    subtract := NewSubtractExpression(add, two)
    
    result := subtract.Interpret()
    fmt.Printf("(5 + 3) - 2 = %d\n", result)
    
}

func TestInterpreterBoolean(t *testing.T) {
    // 布尔表达式示例: (true AND false) OR (true AND true)
    trueExp := NewVariableExpression("true", true)
    falseExp := NewVariableExpression("false", false)
    
    and1 := NewAndExpression(trueExp, falseExp)
    and2 := NewAndExpression(trueExp, trueExp)
    or := NewOrExpression(and1, and2)
    
    result := or.Interpret()
    fmt.Printf("(true AND false) OR (true AND true) = %v\n", result)
    
    // NOT表达式示例
    not := NewNotExpression(trueExp)
	fmt.Printf("NOT true = %v\n", not.Interpret())
}
