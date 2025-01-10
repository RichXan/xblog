package designpattern

import "fmt"

// Strategy 定义策略接口
type Strategy interface {
	Execute(data interface{}) interface{}
}

// ConcreteStrategyA 具体策略A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(data interface{}) interface{} {
	return "Executing strategy A with " + fmt.Sprint(data)
}

// ConcreteStrategyB 具体策略B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(data interface{}) interface{} {
	return "Executing strategy B with " + fmt.Sprint(data)
}

// Context 上下文
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) interface{} {
	return c.strategy.Execute(data)
}

// 实际应用示例：支付系统
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardPayment 信用卡支付
type CreditCardPayment struct {
	cardNumber string
	cvv        string
}

func NewCreditCardPayment(cardNumber, cvv string) *CreditCardPayment {
	return &CreditCardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, p.cardNumber)
}

// PayPalPayment PayPal支付
type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email: email}
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, p.email)
}

// AlipayPayment 支付宝支付
type AlipayPayment struct {
	id string
}

func NewAlipayPayment(id string) *AlipayPayment {
	return &AlipayPayment{id: id}
}

func (p *AlipayPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Alipay account %s", amount, p.id)
}

// PaymentContext 支付上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (c *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *PaymentContext) ProcessPayment(amount float64) string {
	return c.strategy.Pay(amount)
}
