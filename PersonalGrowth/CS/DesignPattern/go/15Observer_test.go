package designpattern

import (
	"testing"
)

func TestObserver(t *testing.T) {
	// 基本示例
	subject := NewConcreteSubject()

	observer1 := NewConcreteObserver("Observer 1")
	observer2 := NewConcreteObserver("Observer 2")

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.SetState("New State!")
}

func TestNewsAgency(t *testing.T) {
	// 新闻订阅示例
	newsAgency := NewNewsAgency()

	emailSub := NewEmailSubscriber("user@example.com")
	smsSub := NewSMSSubscriber("+1234567890")

	newsAgency.Subscribe(emailSub)
	newsAgency.Subscribe(smsSub)

	newsAgency.PublishNews("Breaking News: Go 2.0 Released!")
}
