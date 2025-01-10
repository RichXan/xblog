package designpattern

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	// 基本示例
	objectStructure := &ObjectStructure{}
	objectStructure.Attach(NewConcreteElementA("A1"))
	objectStructure.Attach(NewConcreteElementB("B1"))

	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}

	objectStructure.Accept(visitor1)
	objectStructure.Accept(visitor2)

}

func TestVisitorFile(t *testing.T) {
	// 文件系统示例
	root := NewDirectoryVisitor("root")
	docs := NewDirectoryVisitor("docs")
	root.Add(docs)

	file1 := NewFileVisitor("file1.txt", 100)
	file2 := NewFileVisitor("file2.txt", 200)
	docs.Add(file1)
	docs.Add(file2)

	sizeVisitor := &SizeVisitor{}
	root.Accept(sizeVisitor)
	fmt.Printf("Total size: %d bytes\n", sizeVisitor.GetTotalSize())
}
