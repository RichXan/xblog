package designpattern

import (
	"fmt"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	// 基本示例
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB).SetNext(handlerC)

	fmt.Println(handlerA.Handle("A"))
	fmt.Println(handlerA.Handle("B"))
	fmt.Println(handlerA.Handle("C"))
	fmt.Println(handlerA.Handle("D"))

}

func TestChainOfResponsibilityLeave(t *testing.T) {
	// 请假审批示例
	teamLeader := &TeamLeader{}
	manager := &Manager{}
	director := &DirectorLeaveHandler{}

	teamLeader.SetNext(manager)
	manager.SetNext(director)

	request1 := NewLeaveRequest("John", 2, "Personal matters")
	request2 := NewLeaveRequest("Alice", 5, "Family vacation")
	request3 := NewLeaveRequest("Bob", 10, "Medical treatment")
	request4 := NewLeaveRequest("Emma", 15, "Long vacation")

	fmt.Println(teamLeader.HandleRequest(request1))
	fmt.Println(teamLeader.HandleRequest(request2))
	fmt.Println(teamLeader.HandleRequest(request3))
	fmt.Println(teamLeader.HandleRequest(request4))
}
