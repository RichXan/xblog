package designpattern

import "testing"

func TestAbstractFactory(t *testing.T) {
	factory := &HypeFactoryImpl{}
	huaweiFactory := factory.CreateFactory(FactoryHuawei)
	xiaomiFactory := factory.CreateFactory(FactoryXiaomi)

	huaweiCellphone := huaweiFactory.CreateCellphone()
	xiaomiCellphone := xiaomiFactory.CreateCellphone()

	huaweiCellphone.Call()
	xiaomiCellphone.Call()
}
