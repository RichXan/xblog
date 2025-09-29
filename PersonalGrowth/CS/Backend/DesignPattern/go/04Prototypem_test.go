package designpattern

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	cache := NewShapeCache()
	cache.LoadCache()

	shape1 := cache.GetShape(CircleId)
	shape2 := cache.GetShape(SquareId)
	shape3 := cache.GetShape(RectangleId)

	cshape1 := shape1.Clone()
	cshape2 := shape2.Clone()
	cshape3 := shape3.Clone()

	fmt.Println(cshape1.GetType())
	fmt.Println(cshape2.GetType())
	fmt.Println(cshape3.GetType())
}
