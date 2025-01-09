package designpattern

import "fmt"

// Observer 定义观察者接口
type Observer interface {
	Update(message string)
}

// Subject 定义主题接口
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

// ConcreteSubject 具体主题
type ConcreteSubject struct {
	observers []Observer
	state     string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.Notify(state)
}

// ConcreteObserver 具体观察者
type ConcreteObserver struct {
	name string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("Observer %s received: %s\n", o.name, message)
}

// 实际应用示例：新闻订阅系统
type NewsAgency struct {
	subscribers []NewsSubscriber
	latestNews  string
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		subscribers: make([]NewsSubscriber, 0),
	}
}

func (n *NewsAgency) Subscribe(subscriber NewsSubscriber) {
	n.subscribers = append(n.subscribers, subscriber)
}

func (n *NewsAgency) Unsubscribe(subscriber NewsSubscriber) {
	for i, sub := range n.subscribers {
		if sub == subscriber {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NewsAgency) PublishNews(news string) {
	n.latestNews = news
	for _, subscriber := range n.subscribers {
		subscriber.ReceiveNews(news)
	}
}

type NewsSubscriber interface {
	ReceiveNews(news string)
}

type EmailSubscriber struct {
	email string
}

func NewEmailSubscriber(email string) *EmailSubscriber {
	return &EmailSubscriber{email: email}
}

func (s *EmailSubscriber) ReceiveNews(news string) {
	fmt.Printf("Sending news to %s: %s\n", s.email, news)
}

type SMSSubscriber struct {
	phoneNumber string
}

func NewSMSSubscriber(phoneNumber string) *SMSSubscriber {
	return &SMSSubscriber{phoneNumber: phoneNumber}
}

func (s *SMSSubscriber) ReceiveNews(news string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.phoneNumber, news)
}
