package designpattern

import "fmt"

// 实际应用示例：订单状态系统
type OrderState interface {
	Handle(order *Order)
	GetName() string
}

type Order struct {
	state OrderState
}

func NewOrder() *Order {
	return &Order{state: &NewOrderState{}}
}

func (o *Order) SetState(state OrderState) {
	o.state = state
}

func (o *Order) Process() {
	o.state.Handle(o)
}

func (o *Order) GetStateName() string {
	return o.state.GetName()
}

// NewOrderState 新订单状态
type NewOrderState struct{}

func (s *NewOrderState) Handle(order *Order) {
	fmt.Println("Processing new order, transitioning to paid")
	order.SetState(&PaidOrderState{})
}

func (s *NewOrderState) GetName() string {
	return "New Order"
}

// PaidOrderState 已支付状态
type PaidOrderState struct{}

func (s *PaidOrderState) Handle(order *Order) {
	fmt.Println("Processing paid order, transitioning to shipped")
	order.SetState(&ShippedOrderState{})
}

func (s *PaidOrderState) GetName() string {
	return "Paid"
}

// ShippedOrderState 已发货状态
type ShippedOrderState struct{}

func (s *ShippedOrderState) Handle(order *Order) {
	fmt.Println("Processing shipped order, transitioning to delivered")
	order.SetState(&DeliveredOrderState{})
}

func (s *ShippedOrderState) GetName() string {
	return "Shipped"
}

// DeliveredOrderState 已送达状态
type DeliveredOrderState struct{}

func (s *DeliveredOrderState) Handle(order *Order) {
	fmt.Println("Order has been delivered, no further transitions")
}

func (s *DeliveredOrderState) GetName() string {
	return "Delivered"
}
