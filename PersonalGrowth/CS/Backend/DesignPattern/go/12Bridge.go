package designpattern

import "fmt"

type DrawAPI interface {
	DrawCircle(x, y, radius int)
}

type DrawingAPI1 struct{}

func (d *DrawingAPI1) DrawCircle(x, y, radius int) {
	fmt.Printf("API1.circle at %d:%d radius %d\n", x, y, radius)
}

type DrawingAPI2 struct{}

func (d *DrawingAPI2) DrawCircle(x, y, radius int) {
	fmt.Printf("API2.circle at %d:%d radius %d\n", x, y, radius)
}

type CircleShape struct {
	x, y, radius int
	drawAPI      DrawAPI
}

// type Shape interface {
//     Draw()
// }

func NewCircleShape(x, y, radius int, drawAPI DrawAPI) *CircleShape {
	return &CircleShape{
		x:       x,
		y:       y,
		radius:  radius,
		drawAPI: drawAPI,
	}
}

func (c *CircleShape) Draw() {
	c.drawAPI.DrawCircle(c.x, c.y, c.radius)
}
