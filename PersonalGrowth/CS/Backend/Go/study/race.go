package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var config *Config

type Config struct {
	Name string
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{Name: "Go"}
	})
	return config
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(GetConfig().Name)
			wg.Done()
		}()
	}
	wg.Wait()
}
