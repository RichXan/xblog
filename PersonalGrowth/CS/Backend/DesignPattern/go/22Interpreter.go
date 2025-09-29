package designpattern

// Expression 定义解释器接口
type Expression interface {
    Interpret() int
}

// NumberExpression 数字表达式
type NumberExpression struct {
    number int
}

func NewNumberExpression(number int) *NumberExpression {
    return &NumberExpression{number: number}
}

func (n *NumberExpression) Interpret() int {
    return n.number
}

// AddExpression 加法表达式
type AddExpression struct {
    left  Expression
    right Expression
}

func NewAddExpression(left, right Expression) *AddExpression {
    return &AddExpression{
        left:  left,
        right: right,
    }
}

func (a *AddExpression) Interpret() int {
    return a.left.Interpret() + a.right.Interpret()
}

// SubtractExpression 减法表达式
type SubtractExpression struct {
    left  Expression
    right Expression
}

func NewSubtractExpression(left, right Expression) *SubtractExpression {
    return &SubtractExpression{
        left:  left,
        right: right,
    }
}

func (s *SubtractExpression) Interpret() int {
    return s.left.Interpret() - s.right.Interpret()
}

// 实际应用示例：简单的布尔表达式解释器
type BooleanExpression interface {
    Interpret() bool
}

// VariableExpression 变量表达式
type VariableExpression struct {
    name  string
    value bool
}

func NewVariableExpression(name string, value bool) *VariableExpression {
    return &VariableExpression{
        name:  name,
        value: value,
    }
}

func (v *VariableExpression) Interpret() bool {
    return v.value
}

// AndExpression AND表达式
type AndExpression struct {
    left  BooleanExpression
    right BooleanExpression
}

func NewAndExpression(left, right BooleanExpression) *AndExpression {
    return &AndExpression{
        left:  left,
        right: right,
    }
}

func (a *AndExpression) Interpret() bool {
    return a.left.Interpret() && a.right.Interpret()
}

// OrExpression OR表达式
type OrExpression struct {
    left  BooleanExpression
    right BooleanExpression
}

func NewOrExpression(left, right BooleanExpression) *OrExpression {
    return &OrExpression{
        left:  left,
        right: right,
    }
}

func (o *OrExpression) Interpret() bool {
    return o.left.Interpret() || o.right.Interpret()
}

// NotExpression NOT表达式
type NotExpression struct {
    expression BooleanExpression
}

func NewNotExpression(expression BooleanExpression) *NotExpression {
    return &NotExpression{expression: expression}
}

func (n *NotExpression) Interpret() bool {
    return !n.expression.Interpret()
}