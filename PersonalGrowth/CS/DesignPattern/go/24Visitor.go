package designpattern

import "fmt"

// Element 定义元素接口
type Element interface {
	Accept(visitor Visitor)
}

// Visitor 定义访问者接口
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteElementA 具体元素A
type ConcreteElementA struct {
	name string
}

func NewConcreteElementA(name string) *ConcreteElementA {
	return &ConcreteElementA{name: name}
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB 具体元素B
type ConcreteElementB struct {
	name string
}

func NewConcreteElementB(name string) *ConcreteElementB {
	return &ConcreteElementB{name: name}
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// ConcreteVisitor1 具体访问者1
type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Visitor1 visited ElementA: %s\n", element.name)
}

func (v *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Visitor1 visited ElementB: %s\n", element.name)
}

// ConcreteVisitor2 具体访问者2
type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Visitor2 visited ElementA: %s\n", element.name)
}

func (v *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Visitor2 visited ElementB: %s\n", element.name)
}

// ObjectStructure 对象结构
type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

// 实际应用示例：文件系统访问
type FileSystemElement interface {
	Accept(visitor FileSystemVisitor)
	GetName() string
	GetSize() int
}

type FileSystemVisitor interface {
	VisitFile(file *FileVisitor)
	VisitDirectory(directory *DirectoryVisitor)
}

type FileVisitor struct {
	name string
	size int
}

func NewFileVisitor(name string, size int) *FileVisitor {
	return &FileVisitor{name: name, size: size}
}

func (f *FileVisitor) Accept(visitor FileSystemVisitor) {
	visitor.VisitFile(f)
}

func (f *FileVisitor) GetName() string { return f.name }
func (f *FileVisitor) GetSize() int    { return f.size }

type DirectoryVisitor struct {
	name     string
	children []FileSystemElement
}

func NewDirectoryVisitor(name string) *DirectoryVisitor {
	return &DirectoryVisitor{name: name}
}

func (d *DirectoryVisitor) Accept(visitor FileSystemVisitor) {
	visitor.VisitDirectory(d)
	for _, child := range d.children {
		child.Accept(visitor)
	}
}

func (d *DirectoryVisitor) Add(element FileSystemElement) {
	d.children = append(d.children, element)
}

func (d *DirectoryVisitor) GetName() string { return d.name }
func (d *DirectoryVisitor) GetSize() int {
	total := 0
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}

// SizeVisitor 计算大小的访问者
type SizeVisitor struct {
	totalSize int
}

func (v *SizeVisitor) VisitFile(file *FileVisitor) {
	v.totalSize += file.GetSize()
}

func (v *SizeVisitor) VisitDirectory(directory *DirectoryVisitor) {
	// 目录本身不占用空间
}

func (v *SizeVisitor) GetTotalSize() int {
	return v.totalSize
}
