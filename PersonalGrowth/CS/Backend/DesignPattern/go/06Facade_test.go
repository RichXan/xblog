package designpattern

import "testing"

func TestFacade(t *testing.T) {
	facade := NewComputerFacade()
	facade.Start()
	facade.Shutdown()
}
