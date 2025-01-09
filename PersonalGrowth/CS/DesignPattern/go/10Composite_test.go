package designpattern

import "testing"

func TestComposite(t *testing.T) {
	root := NewDirectory("root")
	root.Add(NewFile("file1.txt", 100))
	root.Add(NewFile("file2.txt", 200))
	root.Add(NewDirectory("subdir"))
	root.Print("")
}
