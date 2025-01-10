package designpattern

// Memento 备忘录
type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) GetState() string {
	return m.state
}

// Originator 发起人
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SaveToMemento() *Memento {
	return NewMemento(o.state)
}

func (o *Originator) RestoreFromMemento(memento *Memento) {
	o.state = memento.GetState()
}

// Caretaker 管理者
type Caretaker struct {
	mementos []*Memento
}

func NewCaretaker() *Caretaker {
	return &Caretaker{
		mementos: make([]*Memento, 0),
	}
}

func (c *Caretaker) AddMemento(memento *Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index >= 0 && index < len(c.mementos) {
		return c.mementos[index]
	}
	return nil
}

// 实际应用示例：文本编辑器
type TextEditorMemento struct {
	content        string
	cursorPosition int
}

func NewTextEditorMemento(content string, cursorPosition int) *TextEditorMemento {
	return &TextEditorMemento{
		content:        content,
		cursorPosition: cursorPosition,
	}
}

type MementoTextEditor struct {
	content        string
	cursorPosition int
	history        *TextEditorHistory
}

func NewMementoTextEditor() *MementoTextEditor {
	return &MementoTextEditor{
		history: NewTextEditorHistory(),
	}
}

func (e *MementoTextEditor) Write(text string) {
	e.history.SaveState(e.content, e.cursorPosition)
	if e.cursorPosition > len(e.content) {
		e.content += text
	} else {
		e.content = e.content[:e.cursorPosition] + text + e.content[e.cursorPosition:]
	}
	e.cursorPosition += len(text)
}

func (e *MementoTextEditor) MoveCursor(position int) {
	if position >= 0 && position <= len(e.content) {
		e.cursorPosition = position
	}
}

func (e *MementoTextEditor) Undo() bool {
	if state := e.history.Undo(); state != nil {
		e.content = state.content
		e.cursorPosition = state.cursorPosition
		return true
	}
	return false
}

func (e *MementoTextEditor) GetContent() string {
	return e.content
}

func (e *MementoTextEditor) GetCursorPosition() int {
	return e.cursorPosition
}

type TextEditorHistory struct {
	states       []*TextEditorMemento
	currentIndex int
}

func NewTextEditorHistory() *TextEditorHistory {
	return &TextEditorHistory{
		states:       make([]*TextEditorMemento, 0),
		currentIndex: -1,
	}
}

func (h *TextEditorHistory) SaveState(content string, cursorPosition int) {
	// 移除当前状态之后的所有状态
	if h.currentIndex < len(h.states)-1 {
		h.states = h.states[:h.currentIndex+1]
	}
	h.states = append(h.states, NewTextEditorMemento(content, cursorPosition))
	h.currentIndex++
}

func (h *TextEditorHistory) Undo() *TextEditorMemento {
	if h.currentIndex > 0 {
		h.currentIndex--
		return h.states[h.currentIndex]
	}
	return nil
}
