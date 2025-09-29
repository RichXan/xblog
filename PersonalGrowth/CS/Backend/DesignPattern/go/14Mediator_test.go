package designpattern

import "testing"

// 基本示例
func TestMediator(t *testing.T) {
	mediator := NewConcreteMediator()

	colleague1 := NewConcreteColleague("Colleague1", mediator)
	_ = NewConcreteColleague("Colleague2", mediator)
	_ = NewConcreteColleague("Colleague3", mediator)

	colleague1.Send("Hello from Colleague1, I am Colleague1.")
}

// 聊天室示例
func TestChatRoom(t *testing.T) {
	chatRoom := NewChatRoom()

	alice := NewUser("Alice", chatRoom)
	_ = NewUser("Bob", chatRoom)
	_ = NewUser("Charlie", chatRoom)

	alice.Send("Hi everyone!")
}
