package designpattern

import "testing"

func TestBridge(t *testing.T) {
	api1 := &DrawingAPI1{}
	api2 := &DrawingAPI2{}

	circle1 := NewCircleShape(1, 2, 3, api1)
	circle2 := NewCircleShape(5, 7, 11, api2)

	circle1.Draw()
	circle2.Draw()
}
