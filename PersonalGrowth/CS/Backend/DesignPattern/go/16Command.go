package designpattern

import "fmt"

// Command 定义命令接口
type Command interface {
	Execute() string
	Undo() string
}

// Receiver 命令接收者
type Receiver struct {
	name string
}

func NewReceiver(name string) *Receiver {
	return &Receiver{name: name}
}

func (r *Receiver) Action() string {
	return fmt.Sprintf("Receiver %s is handling the request", r.name)
}

func (r *Receiver) UndoAction() string {
	return fmt.Sprintf("Receiver %s is undoing the request", r.name)
}

// ConcreteCommand 具体命令
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() string {
	return c.receiver.Action()
}

func (c *ConcreteCommand) Undo() string {
	return c.receiver.UndoAction()
}

// Invoker 命令调用者
type Invoker struct {
	commands []Command
	history  []Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make([]Command, 0),
		history:  make([]Command, 0),
	}
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() []string {
	var results []string
	for _, cmd := range i.commands {
		results = append(results, cmd.Execute())
		i.history = append(i.history, cmd)
	}
	i.commands = make([]Command, 0)
	return results
}

func (i *Invoker) UndoLastCommand() string {
	if len(i.history) == 0 {
		return "No commands to undo"
	}
	lastIndex := len(i.history) - 1
	lastCommand := i.history[lastIndex]
	i.history = i.history[:lastIndex]
	return lastCommand.Undo()
}

// 实际应用示例：文本编辑器
type TextEditor struct {
	content string
}

func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

func (e *TextEditor) GetContent() string {
	return e.content
}

// InsertCommand 插入文本命令
type InsertCommand struct {
	editor *TextEditor
	text   string
	pos    int
}

func NewInsertCommand(editor *TextEditor, text string, pos int) *InsertCommand {
	return &InsertCommand{
		editor: editor,
		text:   text,
		pos:    pos,
	}
}

func (c *InsertCommand) Execute() string {
	if c.pos > len(c.editor.content) {
		c.editor.content += c.text
	} else {
		c.editor.content = c.editor.content[:c.pos] + c.text + c.editor.content[c.pos:]
	}
	return fmt.Sprintf("Inserted '%s' at position %d", c.text, c.pos)
}

func (c *InsertCommand) Undo() string {
	if c.pos > len(c.editor.content) {
		c.editor.content = c.editor.content[:len(c.editor.content)-len(c.text)]
	} else {
		c.editor.content = c.editor.content[:c.pos] + c.editor.content[c.pos+len(c.text):]
	}
	return fmt.Sprintf("Undid insertion of '%s' at position %d", c.text, c.pos)
}

// DeleteCommand 删除文本命令
type DeleteCommand struct {
	editor      *TextEditor
	text        string
	pos         int
	deletedText string
}

func NewDeleteCommand(editor *TextEditor, pos int, length int) *DeleteCommand {
	return &DeleteCommand{
		editor: editor,
		pos:    pos,
		text:   editor.content[pos : pos+length],
	}
}

func (c *DeleteCommand) Execute() string {
	c.deletedText = c.editor.content[c.pos : c.pos+len(c.text)]
	c.editor.content = c.editor.content[:c.pos] + c.editor.content[c.pos+len(c.text):]
	return fmt.Sprintf("Deleted '%s' at position %d", c.deletedText, c.pos)
}

func (c *DeleteCommand) Undo() string {
	c.editor.content = c.editor.content[:c.pos] + c.deletedText + c.editor.content[c.pos:]
	return fmt.Sprintf("Restored '%s' at position %d", c.deletedText, c.pos)
}
