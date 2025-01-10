package designpattern

import "fmt"

// AbstractClass 定义抽象类
type AbstractClass interface {
	TemplateMethod()
	PrimitiveOperation1()
	PrimitiveOperation2()
	Hook()
}

// Template 提供模板方法的基础实现
type Template struct {
	AbstractClass
}

// TemplateMethod 定义算法骨架
func (t *Template) TemplateMethod() {
	t.AbstractClass.PrimitiveOperation1()
	t.AbstractClass.PrimitiveOperation2()
	t.AbstractClass.Hook()
}

// Hook 提供默认实现
func (t *Template) Hook() {
	// 默认为空实现，子类可以选择性重写
}

// ConcreteClassA 具体实现A
type ConcreteClassA struct {
	Template
}

func NewConcreteClassA() *ConcreteClassA {
	concrete := &ConcreteClassA{}
	concrete.AbstractClass = concrete
	return concrete
}

func (c *ConcreteClassA) PrimitiveOperation1() {
	fmt.Println("ConcreteClassA: PrimitiveOperation1")
}

func (c *ConcreteClassA) PrimitiveOperation2() {
	fmt.Println("ConcreteClassA: PrimitiveOperation2")
}

// ConcreteClassB 具体实现B
type ConcreteClassB struct {
	Template
}

func NewConcreteClassB() *ConcreteClassB {
	concrete := &ConcreteClassB{}
	concrete.AbstractClass = concrete
	return concrete
}

func (c *ConcreteClassB) PrimitiveOperation1() {
	fmt.Println("ConcreteClassB: PrimitiveOperation1")
}

func (c *ConcreteClassB) PrimitiveOperation2() {
	fmt.Println("ConcreteClassB: PrimitiveOperation2")
}

func (c *ConcreteClassB) Hook() {
	fmt.Println("ConcreteClassB: Hook")
}

// 实际应用示例：数据处理框架
type DataMiner interface {
	Mine()
	OpenFile()
	ExtractData()
	ParseData()
	AnalyzeData()
	SendReport()
	CloseFile()
}

// DataMinerTemplate 提供基础实现
type DataMinerTemplate struct {
	DataMiner
}

func (d *DataMinerTemplate) Mine() {
	d.DataMiner.OpenFile()
	d.DataMiner.ExtractData()
	d.DataMiner.ParseData()
	d.DataMiner.AnalyzeData()
	d.DataMiner.SendReport()
	d.DataMiner.CloseFile()
}

// PDFMiner PDF文件处理器
type PDFMiner struct {
	DataMinerTemplate
}

func NewPDFMiner() *PDFMiner {
	miner := &PDFMiner{}
	miner.DataMiner = miner
	return miner
}

func (p *PDFMiner) OpenFile() {
	fmt.Println("Opening PDF file")
}

func (p *PDFMiner) ExtractData() {
	fmt.Println("Extracting data from PDF")
}

func (p *PDFMiner) ParseData() {
	fmt.Println("Parsing PDF data")
}

func (p *PDFMiner) AnalyzeData() {
	fmt.Println("Analyzing PDF data")
}

func (p *PDFMiner) SendReport() {
	fmt.Println("Sending PDF report")
}

func (p *PDFMiner) CloseFile() {
	fmt.Println("Closing PDF file")
}

// CSVMiner CSV文件处理器
type CSVMiner struct {
	DataMinerTemplate
}

func NewCSVMiner() *CSVMiner {
	miner := &CSVMiner{}
	miner.DataMiner = miner
	return miner
}

func (c *CSVMiner) OpenFile() {
	fmt.Println("Opening CSV file")
}

func (c *CSVMiner) ExtractData() {
	fmt.Println("Extracting data from CSV")
}

func (c *CSVMiner) ParseData() {
	fmt.Println("Parsing CSV data")
}

func (c *CSVMiner) AnalyzeData() {
	fmt.Println("Analyzing CSV data")
}

func (c *CSVMiner) SendReport() {
	fmt.Println("Sending CSV report")
}

func (c *CSVMiner) CloseFile() {
	fmt.Println("Closing CSV file")
}
