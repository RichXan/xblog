package designpattern

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	// 基本示例
	context := NewContext(&ConcreteStrategyA{})
	result := context.ExecuteStrategy("data")
	fmt.Println("result1:", result)

	context.SetStrategy(&ConcreteStrategyB{})
	result = context.ExecuteStrategy("data")
	fmt.Println("result2:", result)
}

func TestPayment(t *testing.T) {
	// 支付示例
	creditCard := NewCreditCardPayment("1234-5678-9012-3456", "123")
	paypal := NewPayPalPayment("user@example.com")
	alipay := NewAlipayPayment("user123")

	payment := NewPaymentContext(creditCard)
	fmt.Println(payment.ProcessPayment(100.50))

	payment.SetStrategy(paypal)
	fmt.Println(payment.ProcessPayment(50.75))

	payment.SetStrategy(alipay)
	fmt.Println(payment.ProcessPayment(200.00))
}
