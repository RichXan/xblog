package designpattern

// Builder 是生成器接口
type Builder interface {
	AddWheel()
	AddEngine()
	AddDoors()
	GetVehicle() Vehicle
	CreateVehicle()
}

// Director 是指挥者
type Director struct{}

// Build 使用生成器构建
func (d *Director) Build(builder Builder) {
	builder.CreateVehicle()
	builder.AddWheel()
	builder.AddEngine()
	builder.AddDoors()
}

// CarBuilder 是汽车生成器
type CarBuilder struct {
	car Vehicle
}

func (b *CarBuilder) AddWheel() {
	b.car.AddPart("Car Wheel")
}

func (b *CarBuilder) AddEngine() {
	b.car.AddPart("Car Engine")
}

func (b *CarBuilder) AddDoors() {
	b.car.AddPart("Car Doors")
}

func (b *CarBuilder) GetVehicle() Vehicle {
	return b.car
}

func (b *CarBuilder) CreateVehicle() {
	b.car = &Car{}
}

// TruckBuilder 是卡车生成器
type TruckBuilder struct {
	truck Vehicle
}

func (b *TruckBuilder) AddWheel() {
	b.truck.AddPart("Truck Engine")
}

func (b *TruckBuilder) AddEngine() {
	b.truck.AddPart("Truck Engine")
}

func (b *TruckBuilder) AddDoors() {
	b.truck.AddPart("Truck Doors")
}

func (b *TruckBuilder) GetVehicle() Vehicle {
	return b.truck
}

func (b *TruckBuilder) CreateVehicle() {
	b.truck = &Truck{}
}

// Vehicle 是车辆接口
type Vehicle interface {
	AddPart(part string)
}

// Car 是具体的汽车
type Car struct {
	parts []string
}

func (c *Car) AddPart(part string) {
	c.parts = append(c.parts, part)
}

// Truck 是具体的卡车
type Truck struct {
	parts []string
}

func (t *Truck) AddPart(part string) {
	t.parts = append(t.parts, part)
}
