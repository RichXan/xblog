package designpattern

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	factory := NewFactory(1)
	fmt.Println(factory.Produce("phone"))
}

