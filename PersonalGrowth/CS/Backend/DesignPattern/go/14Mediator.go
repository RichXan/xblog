package designpattern

import "fmt"

// Mediator 定义中介者接口
type Mediator interface {
	Register(colleague Colleague)
	Send(message string, colleague Colleague)
}

// Colleague 定义同事接口
type Colleague interface {
	Send(message string)
	Receive(message string)
}

// ConcreteMediator 具体中介者
type ConcreteMediator struct {
	colleagues []Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: make([]Colleague, 0),
	}
}

func (m *ConcreteMediator) Register(colleague Colleague) {
	m.colleagues = append(m.colleagues, colleague)
}

func (m *ConcreteMediator) Send(message string, sender Colleague) {
	for _, colleague := range m.colleagues {
		// 不要发送给自己
		if colleague != sender {
			colleague.Receive(message)
		}
	}
}

// ConcreteColleague 具体同事类
type ConcreteColleague struct {
	name     string
	mediator Mediator
}

func NewConcreteColleague(name string, mediator Mediator) *ConcreteColleague {
	colleague := &ConcreteColleague{
		name:     name,
		mediator: mediator,
	}
	mediator.Register(colleague)
	return colleague
}

func (c *ConcreteColleague) Send(message string) {
	fmt.Printf("%s sends: %s\n", c.name, message)
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague) Receive(message string) {
	fmt.Printf("%s receives: %s\n", c.name, message)
}

// 实际应用示例：聊天室系统
type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make(map[string]*User),
	}
}

func (c *ChatRoom) Register(user *User) {
	c.users[user.name] = user
}

func (c *ChatRoom) Send(message string, sender *User) {
	for name, user := range c.users {
		if name != sender.name {
			user.Receive(fmt.Sprintf("From %s: %s", sender.name, message))
		}
	}
}

type User struct {
	name     string
	chatRoom *ChatRoom
}

func NewUser(name string, chatRoom *ChatRoom) *User {
	user := &User{
		name:     name,
		chatRoom: chatRoom,
	}
	chatRoom.Register(user)
	return user
}

func (u *User) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.chatRoom.Send(message, u)
}

func (u *User) Receive(message string) {
	fmt.Printf("%s receives: %s\n", u.name, message)
}
