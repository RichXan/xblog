package designpattern

import "fmt"

type FileSystemNode interface {
	GetName() string
	GetSize() int
	IsDirectory() bool
	Print(prefix string)
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string   { return f.name }
func (f *File) GetSize() int      { return f.size }
func (f *File) IsDirectory() bool { return false }
func (f *File) Print(prefix string) {
	fmt.Printf("%s- %s (%d bytes)\n", prefix, f.name, f.size)
}

type Directory struct {
	name     string
	children []FileSystemNode
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make([]FileSystemNode, 0),
	}
}

func (d *Directory) Add(child FileSystemNode) {
	d.children = append(d.children, child)
}

func (d *Directory) GetName() string   { return d.name }
func (d *Directory) IsDirectory() bool { return true }

func (d *Directory) GetSize() int {
	total := 0
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}

func (d *Directory) Print(prefix string) {
	fmt.Printf("%s+ %s\n", prefix, d.name)
	for _, child := range d.children {
		child.Print(prefix + "  ")
	}
}
