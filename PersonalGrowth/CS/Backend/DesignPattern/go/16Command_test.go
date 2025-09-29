package designpattern

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	// 基本示例
	receiver := NewReceiver("Main")
	command := NewConcreteCommand(receiver)

	// 命令调用者
	invoker := NewInvoker()
	invoker.AddCommand(command)

	results := invoker.ExecuteCommands()
	fmt.Println(results)

	undoResult := invoker.UndoLastCommand()
	fmt.Println(undoResult)

}

func TestTextEditor(t *testing.T) {
	invoker := NewInvoker()
	// 文本编辑器示例
	editor := NewTextEditor()

	insertCmd := NewInsertCommand(editor, "Hello", 0)
	invoker.AddCommand(insertCmd)
	invoker.ExecuteCommands()
	fmt.Println("Content:", editor.GetContent())

	insertCmd2 := NewInsertCommand(editor, " World", 5)
	invoker.AddCommand(insertCmd2)
	invoker.ExecuteCommands()
	fmt.Println("Content:", editor.GetContent())

	deleteCmd := NewDeleteCommand(editor, 5, 6)
	invoker.AddCommand(deleteCmd)
	invoker.ExecuteCommands()
	fmt.Println("Content:", editor.GetContent())

	invoker.UndoLastCommand()
	fmt.Println("Content after undo:", editor.GetContent())
}
