package designpattern

import (
	"fmt"
	"sync"
)

func NewSingleton() *Singleton {
	return &Singleton{}
}

type Singleton struct {
	Name string
}

var (
	once      sync.Once
	singleton *Singleton
)

func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func (s *Singleton) GetName() string {
	return s.Name
}

func (s *Singleton) SetName(name string) {
	s.Name = name
}

func (s *Singleton) PrintName() {
	fmt.Println(s.Name)
}
