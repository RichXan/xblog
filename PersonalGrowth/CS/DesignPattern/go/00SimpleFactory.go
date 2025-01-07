package designpattern

import "fmt"

type Factory interface {
	Produce(product string) string
}

type Huawei struct{}

func (*Huawei) Produce(product string) string {
	return fmt.Sprintf("hi %s", product)
}

type Apple struct{}

func (*Apple) Produce(product string) string {
	return fmt.Sprintf("hi %s", product)
}

func NewFactory(factoryType int) Factory {
	if factoryType == 1 {
		return &Huawei{}
	} else if factoryType == 2 {
		return &Apple{}
	}
	return nil
}
