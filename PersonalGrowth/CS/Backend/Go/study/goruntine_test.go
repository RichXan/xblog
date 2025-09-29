package main

import (
	"fmt"
	"sync"
	"testing"
)

var a string

func f() {
	fmt.Println(a)
}

func TestSum(t *testing.T) {
	go f()
	a = "hello, world"
	if a != "hello, world" {
		t.Errorf("Expected hello, world but got %s", a)
	}
}

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]string
}

func (s *SafeMap) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func TestSafeMap(t *testing.T) {
	safeMap := &SafeMap{
		m: make(map[string]string),
	}
	safeMap.Set("key", "value1")
	fmt.Println(safeMap.Get("key"))
}
