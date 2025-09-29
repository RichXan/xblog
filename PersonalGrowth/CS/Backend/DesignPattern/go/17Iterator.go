package designpattern

// Iterator 定义迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
	Current() interface{}
}

// Container 定义容器接口
type Container interface {
	CreateIterator() Iterator
}

// ConcreteContainer 具体容器
type ConcreteContainer struct {
	items []interface{}
}

func NewConcreteContainer() *ConcreteContainer {
	return &ConcreteContainer{
		items: make([]interface{}, 0),
	}
}

func (c *ConcreteContainer) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *ConcreteContainer) CreateIterator() Iterator {
	return &ConcreteIterator{
		container: c,
		index:     0,
	}
}

// ConcreteIterator 具体迭代器
type ConcreteIterator struct {
	container *ConcreteContainer
	index     int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.container.items)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.container.items[i.index]
		i.index++
		return item
	}
	return nil
}

func (i *ConcreteIterator) Current() interface{} {
	if i.index < len(i.container.items) {
		return i.container.items[i.index]
	}
	return nil
}

// 实际应用示例：文件系统遍历
type FileSystemItem interface {
	GetName() string
}

type IteratorFile struct {
	name string
}

func NewIteratorFile(name string) *IteratorFile {
	return &IteratorFile{name: name}
}

func (f *IteratorFile) GetName() string {
	return f.name
}

type IteratorDirectory struct {
	name     string
	children []FileSystemItem
}

func NewIteratorDirectory(name string) *IteratorDirectory {
	return &IteratorDirectory{
		name:     name,
		children: make([]FileSystemItem, 0),
	}
}

func (d *IteratorDirectory) Add(item FileSystemItem) {
	d.children = append(d.children, item)
}

func (d *IteratorDirectory) GetName() string {
	return d.name
}

type FileSystemIterator struct {
	items []FileSystemItem
	index int
}

func NewFileSystemIterator(items []FileSystemItem) *FileSystemIterator {
	return &FileSystemIterator{
		items: items,
		index: 0,
	}
}

func (i *FileSystemIterator) HasNext() bool {
	return i.index < len(i.items)
}

func (i *FileSystemIterator) Next() interface{} {
	if i.HasNext() {
		item := i.items[i.index]
		i.index++
		return item
	}
	return nil
}

func (i *FileSystemIterator) Current() interface{} {
	if i.index < len(i.items) {
		return i.items[i.index]
	}
	return nil
}
