package designpattern

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDecorator(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", HelloWorld)
	mux.HandleFunc("GET /how", HowAreYou)

	srv := http.Server{
		Addr: ":8080",
		// 执行装饰器
		Handler: Logger(mux),
	}

	srv.ListenAndServe()
}

func TestCoffee(t *testing.T) {
	// 咖啡订单示例
	coffee := &SimpleCoffee{}
	coffeeWithMilk := NewMilkDecorator(coffee)
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)

	fmt.Printf("Cost: %.2f\n", coffeeWithMilkAndSugar.Cost())
	fmt.Println("Description:", coffeeWithMilkAndSugar.Description())
}
