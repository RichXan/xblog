# 迭代器模式（Iterator）
迭代器模式提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露其内部的表示。这种模式属于行为型模式。

## 主要解决的问题
- 如何顺序访问集合中的元素
- 如何在不暴露集合内部结构的情况下遍历元素
- 如何支持不同的遍历方式
- 如何统一遍历接口

## 应用实例
1. Java的Iterator接口
2. Python的迭代器协议
3. STL的迭代器
4. 数据库结果集遍历

## 使用场景
1. 集合遍历
   - 数组遍历
   - 链表遍历
   - 树结构遍历
2. 数据流处理
   - 文件读取
   - 网络数据流
   - 数据库查询
3. 复杂结构访问
   - 组合对象遍历
   - 多维结构遍历
   - 图结构遍历
4. 并发迭代
   - 安全遍历
   - 快照迭代
   - 并发集合

## 优缺点
### 优点
1. 封装性好
   - 隐藏实现细节
   - 统一访问接口
2. 职责分离
   - 遍历和实现分离
   - 多种遍历方式
3. 扩展性强
   - 易于增加新的遍历方式
   - 不影响聚合对象

### 缺点
1. 类增加
   - 每个集合对应一个迭代器
   - 系统复杂度增加
2. 性能问题
   - 对简单遍历可能过度设计
   - 迭代器创建开销
3. 额外维护
   - 需要同步更新迭代器
   - 并发修改问题

## 代码实现

```golang
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
        index:    0,
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

type File struct {
    name string
}

func NewFile(name string) *File {
    return &File{name: name}
}

func (f *File) GetName() string {
    return f.name
}

type Directory struct {
    name     string
    children []FileSystemItem
}

func NewDirectory(name string) *Directory {
    return &Directory{
        name:     name,
        children: make([]FileSystemItem, 0),
    }
}

func (d *Directory) Add(item FileSystemItem) {
    d.children = append(d.children, item)
}

func (d *Directory) GetName() string {
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
```

## 使用示例

```golang
func main() {
    // 基本示例
    container := NewConcreteContainer()
    container.Add("Item 1")
    container.Add("Item 2")
    container.Add("Item 3")
    
    iterator := container.CreateIterator()
    for iterator.HasNext() {
        fmt.Println(iterator.Next())
    }
    
    // 文件系统示例
    root := NewDirectory("root")
    docs := NewDirectory("docs")
    root.Add(docs)
    
    file1 := NewFile("file1.txt")
    file2 := NewFile("file2.txt")
    docs.Add(file1)
    docs.Add(file2)
    
    iterator := NewFileSystemIterator(root.children)
    for iterator.HasNext() {
        if item, ok := iterator.Next().(FileSystemItem); ok {
            fmt.Println(item.GetName())
        }
    }
}
```

## 类图
```mermaid
classDiagram
    %% 基本迭代器模式
    class Iterator {
        <<interface>>
        +HasNext() bool
        +Next() interface{}
        +Current() interface{}
    }

    class Container {
        <<interface>>
        +CreateIterator() Iterator
    }

    class ConcreteContainer {
        -items []interface{}
        +Add(interface{})
        +CreateIterator() Iterator
    }

    class ConcreteIterator {
        -container ConcreteContainer
        -index int
        +HasNext() bool
        +Next() interface{}
        +Current() interface{}
    }

    %% 文件系统示例
    class FileSystemItem {
        <<interface>>
        +GetName() string
    }

    class File {
        -name string
        +GetName() string
    }

    class Directory {
        -name string
        -children []FileSystemItem
        +Add(FileSystemItem)
        +GetName() string
    }

    class FileSystemIterator {
        -items []FileSystemItem
        -index int
        +HasNext() bool
        +Next() interface{}
        +Current() interface{}
    }

    %% 关系
    Container <|.. ConcreteContainer
    Iterator <|.. ConcreteIterator
    ConcreteContainer ..> ConcreteIterator : creates
    FileSystemItem <|.. File
    FileSystemItem <|.. Directory
    Directory o-- FileSystemItem
    FileSystemIterator o-- FileSystemItem
```

## 说明
1. 迭代器模式的主要角色：
   - Iterator（迭代器）：定义访问和遍历元素的接口
   - ConcreteIterator（具体迭代器）：实现迭代器接口
   - Aggregate（聚合）：定义创建迭代器的接口
   - ConcreteAggregate（具体聚合）：实现创建迭代器的接口
2. 实现要点：
   - 迭代器接口设计
   - 遍历状态维护
   - 并发修改处理
3. 设计考虑：
   - 是否需要双向迭代
   - 是否需要快照迭代
   - 是否需要并发支持
4. 相关模式：
   - 组合模式：遍历树形结构
   - 工厂方法：创建迭代器
   - 备忘录模式：保存迭代状态
