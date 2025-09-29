package designpattern

type OperatorType string

const (
	OperatorPlus  OperatorType = "plus"
	OperatorMinus OperatorType = "minus"
)

// 方法
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 方法基类
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 加法
type PlusOperator struct {
	*OperatorBase
}

func (o *PlusOperator) Result() int {
	return o.a + o.b
}

// 减法
type MinusOperator struct {
	*OperatorBase
}

func (o *MinusOperator) Result() int {
	return o.a - o.b
}

// 工厂方法
type OperatorFactory interface {
	Create() Operator
}

// 加法工厂类
type PlusOperatorFactory struct{ OperatorFactory }

func (f *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 减法工厂类
type MinusOperatorFactory struct{ OperatorFactory }

func (f *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 创建方法工厂类
func CreateFactory(operator OperatorType) Operator {
	switch operator {
	case OperatorPlus:
		factory := &PlusOperatorFactory{}
		return factory.Create()
	case OperatorMinus:
		factory := &MinusOperatorFactory{}
		return factory.Create()
	}
	return nil
}
