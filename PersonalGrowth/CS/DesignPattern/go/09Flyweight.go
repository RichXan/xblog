package designpattern

import "fmt"

// 文字编辑器中的字符
type Character struct {
	char     rune
	font     string
	size     int
	isBold   bool
	isItalic bool
}

func (c *Character) Print() {
	fmt.Printf("Printing character %c with font %s, size %d, bold %v, italic %v\n", c.char, c.font, c.size, c.isBold, c.isItalic)
}

type CharacterFactory struct {
	characters map[string]*Character
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		characters: make(map[string]*Character),
	}
}

func (f *CharacterFactory) GetCharacter(char rune, font string, size int, bold, italic bool) *Character {
	key := fmt.Sprintf("%c-%s-%d-%v-%v", char, font, size, bold, italic)
	if c, ok := f.characters[key]; ok {
		return c
	}

	c := &Character{
		char:     char,
		font:     font,
		size:     size,
		isBold:   bold,
		isItalic: italic,
	}
	f.characters[key] = c
	return c
}
