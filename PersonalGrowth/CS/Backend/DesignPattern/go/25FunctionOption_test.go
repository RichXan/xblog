package designpattern

import (
	"fmt"
	"testing"
	"time"
)

func TestFunctionOption(t *testing.T) {
	client := NewClient(WithName("localhost"), WithTimeout(time.Duration(10)))
	fmt.Println(client)
}
