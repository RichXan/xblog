package designpattern

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	singleton := GetSingleton()

	singleton.SetName("Singleton")
	singleton.PrintName()

	singleton2 := GetSingleton()
	singleton2.PrintName()

	fmt.Println("singleton == singleton2:", singleton == singleton2)
}
