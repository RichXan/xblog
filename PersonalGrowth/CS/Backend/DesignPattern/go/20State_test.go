package designpattern

import (
	"fmt"
	"testing"
)

func TestOrderState(t *testing.T) {
	// 订单状态示例
	order := NewOrder()
	fmt.Println("Order state:", order.GetStateName())

	order.Process() // New -> Paid
	fmt.Println("Order state:", order.GetStateName())

	order.Process() // Paid -> Shipped
	fmt.Println("Order state:", order.GetStateName())

	order.Process() // Shipped -> Delivered
	fmt.Println("Order state:", order.GetStateName())
}
