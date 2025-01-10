package designpattern

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	// 基本示例
	originator := &Originator{}
	caretaker := NewCaretaker()

	originator.SetState("State 1")
	caretaker.AddMemento(originator.SaveToMemento())

	originator.SetState("State 2")
	caretaker.AddMemento(originator.SaveToMemento())

	originator.SetState("State 3")
	fmt.Println("Current State:", originator.GetState())

	originator.RestoreFromMemento(caretaker.GetMemento(1))
	fmt.Println("Restored to State:", originator.GetState())

}

func TestMementoTextEditor(t *testing.T) {
	// 文本编辑器示例
	editor := NewMementoTextEditor()

	editor.Write("Hello")
	fmt.Printf("Content: %s, Cursor: %d\n", editor.GetContent(), editor.GetCursorPosition())

	editor.Write(" World")
	fmt.Printf("Content: %s, Cursor: %d\n", editor.GetContent(), editor.GetCursorPosition())

	editor.Undo()
	fmt.Printf("After undo - Content: %s, Cursor: %d\n", editor.GetContent(), editor.GetCursorPosition())
}
