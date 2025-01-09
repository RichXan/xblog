package designpattern

import (
	"log"
	"net/http"
	"time"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

func HowAreYou(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("how are you"))
}

// log装饰器
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		// 执行被装饰的handler
		next.ServeHTTP(w, req)
		log.Printf("spend time: %v", time.Since(now))
	}
	return http.HandlerFunc(fn)
}

// 咖啡接口
type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 1.0
}

func (c *SimpleCoffee) Description() string {
	return "Simple coffee"
}

type MilkDecorator struct {
	coffee Coffee
}

func NewMilkDecorator(c Coffee) *MilkDecorator {
	return &MilkDecorator{coffee: c}
}

func (d *MilkDecorator) Cost() float64 {
	return d.coffee.Cost() + 0.5
}

func (d *MilkDecorator) Description() string {
	return d.coffee.Description() + ", milk"
}

type SugarDecorator struct {
	coffee Coffee
}

func NewSugarDecorator(c Coffee) *SugarDecorator {
	return &SugarDecorator{coffee: c}
}

func (d *SugarDecorator) Cost() float64 {
	return d.coffee.Cost() + 0.2
}

func (d *SugarDecorator) Description() string {
	return d.coffee.Description() + ", sugar"
}
