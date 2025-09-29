package designpattern

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	charFactory := NewCharacterFactory()

	char1 := charFactory.GetCharacter('A', "Arial", 12, true, false)
	char2 := charFactory.GetCharacter('A', "Arial", 12, true, false)

	// char1 和 char2 是同一个对象的引用
	fmt.Println(char1 == char2) // true
}
