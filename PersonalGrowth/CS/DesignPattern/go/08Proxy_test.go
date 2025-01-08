package designpattern

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := NewProxy()
	result := proxy.Do()
	fmt.Println(result)
}
