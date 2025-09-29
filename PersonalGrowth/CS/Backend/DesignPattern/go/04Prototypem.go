package designpattern

const (
	CircleId    = "circle"
	SquareId    = "square"
	RectangleId = "rectangle"
)

// Shape 是形状接口
type Shape interface {
	GetType() string
	GetId() string
	SetId(id string)
	Clone() Shape
}

// BaseShape 提供基础实现
type BaseShape struct {
	id  string
	typ string
}

func (s *BaseShape) GetType() string {
	return s.typ
}

func (s *BaseShape) GetId() string {
	return s.id
}

func (s *BaseShape) SetId(id string) {
	s.id = id
}

// Circle 实现
type Circle struct {
	BaseShape
}

func NewCircle() *Circle {
	return &Circle{
		BaseShape: BaseShape{typ: "Circle"},
	}
}

func (c *Circle) Clone() Shape {
	return &Circle{
		BaseShape: BaseShape{
			id:  c.id,
			typ: c.typ,
		},
	}
}

// Rectangle 实现
type Rectangle struct {
	BaseShape
}

func NewRectangle() *Rectangle {
	return &Rectangle{
		BaseShape: BaseShape{typ: "Rectangle"},
	}
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		BaseShape: BaseShape{
			id:  r.id,
			typ: r.typ,
		},
	}
}

// Square 实现
type Square struct {
	BaseShape
}

func NewSquare() *Square {
	return &Square{
		BaseShape: BaseShape{typ: "Square"},
	}
}

func (s *Square) Clone() Shape {
	return &Square{
		BaseShape: BaseShape{
			id:  s.id,
			typ: s.typ,
		},
	}
}

// ShapeCache 形状缓存管理器
type ShapeCache struct {
	shapeMap map[string]Shape
}

func NewShapeCache() *ShapeCache {
	return &ShapeCache{
		shapeMap: make(map[string]Shape),
	}
}

func (c *ShapeCache) GetShape(id string) Shape {
	cachedShape := c.shapeMap[id]
	return cachedShape.Clone()
}

func (c *ShapeCache) LoadCache() {
	circle := NewCircle()
	circle.SetId(CircleId)
	c.shapeMap[circle.GetId()] = circle

	square := NewSquare()
	square.SetId(SquareId)
	c.shapeMap[square.GetId()] = square

	rectangle := NewRectangle()
	rectangle.SetId(RectangleId)
	c.shapeMap[rectangle.GetId()] = rectangle
}
