// package designpattern
package main

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
func main() {
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
