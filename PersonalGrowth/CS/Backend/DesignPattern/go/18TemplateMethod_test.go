package designpattern

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	// 基本示例
	classA := NewConcreteClassA()
	classA.TemplateMethod()

	classB := NewConcreteClassB()
	classB.TemplateMethod()
}

func TestPDFMiner(t *testing.T) {
	pdfMiner := NewPDFMiner()
	pdfMiner.Mine()
}

func TestCSVMiner(t *testing.T) {
	csvMiner := NewCSVMiner()
	csvMiner.Mine()
}
