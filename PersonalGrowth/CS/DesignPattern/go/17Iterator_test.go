package designpattern

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	// 基本示例
	container := NewConcreteContainer()
	container.Add("Item 1")
	container.Add("Item 2")
	container.Add("Item 3")

	iterator := container.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}

func TestFileSystemIterator(t *testing.T) {
	// 文件系统示例
	root := NewIteratorDirectory("root")
	docs := NewIteratorDirectory("docs")
	root.Add(docs)

	file1 := NewIteratorFile("file1.txt")
	file2 := NewIteratorFile("file2.txt")
	docs.Add(file1)
	docs.Add(file2)

	iterator := NewFileSystemIterator(root.children)
	for iterator.HasNext() {
		if item, ok := iterator.Next().(FileSystemItem); ok {
			fmt.Println(item.GetName())
		}
	}
}
