package designpattern

import (
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
