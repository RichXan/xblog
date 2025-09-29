package designpattern

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	director := &Director{}

	// 构建汽车
	carBuilder := &CarBuilder{}
	director.Build(carBuilder)
	car := carBuilder.GetVehicle()
	fmt.Println(car)

	// 构建卡车
	truckBuilder := &TruckBuilder{}
	director.Build(truckBuilder)
	truck := truckBuilder.GetVehicle()
	fmt.Println(truck)
}
