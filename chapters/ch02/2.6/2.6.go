package main

import (
	"bytes"
	"fmt"
)

const (
	// white-space characters
	space   = 32
	tab     = 9
	newline = 10

	// tags - Public to be used outside package
	// terminals
	TagNum = 256
	TagId  = 257

	// reserved key-words
	TagTrue  = 258
	TagFalse = 259
)

// Interface satisfied by different types of tokens: num, word
type token interface {
	tag() int
	value() string
}

type char struct {
	Tag   int
	Value byte
}

func NewChar(character byte) *char {
	return &char{
		Tag:   TagId,
		Value: character,
	}
}

func (c char) tag() int {
	return c.Tag
}

func (c char) value() string {
	return string(c.Value)
}

type num struct {
	Tag   int
	Value int
}

func NewNum(value int) *num {
	return &num{
		Tag:   TagNum,
		Value: value,
	}
}

func (n num) tag() int {
	return n.Tag
}

func (n num) value() string {
	return fmt.Sprint(n.Value)
}

type word struct {
	Tag    int
	Lexeme string
}

func (w word) tag() int {
	return w.Tag
}

func (w word) value() string {
	return w.Lexeme
}

func NewWord(str string) *word {
	return &word{
		Tag:    TagId,
		Lexeme: str,
	}
}

type Lexer struct {
	LineNum int
	buffer  *bytes.Buffer
	peek    byte
	words   map[string]token
}

func NewLexer(buffer *bytes.Buffer) *Lexer {
	// initialise lexer
	words := make(map[string]token)
	lexer := &Lexer{
		LineNum: 0,
		buffer:  buffer,
		peek:    space,
		words:   words,
	}

	// populate map with reserved words
	lexer.reserve(
		&word{
			Tag:    TagTrue,
			Lexeme: "true",
		},
	)

	lexer.reserve(
		&word{
			Tag:    TagFalse,
			Lexeme: "false",
		},
	)

	return lexer
}

func (l *Lexer) nextInputChar() byte {
	c, err := l.buffer.ReadByte()
	if err != nil {
		fmt.Println("end of buffer")
	}

	return c
}

func (l *Lexer) Scan() token {
	// skip white-space
	l.skipWhiteSpace()

	// read number
	num := l.readNumber()
	if num != nil {
		return num
	}

	// read word
	word := l.readLexeme()
	if word != nil {
		return word
	}

	// read character
	return l.readCharacter()
}

func (l *Lexer) reserve(w *word) {
	l.words[w.Lexeme] = w
}

func (l *Lexer) skipWhiteSpace() {
	for {
		l.peek = l.nextInputChar()

		if l.peek == space || l.peek == tab {
			// do nothing
		} else if l.peek == newline {
			l.LineNum++
		} else {
			// l.peek holds a non-white space character
			break
		}
	}
}

func (l *Lexer) readNumber() *num {
	if isDigit(l.peek) {
		value := 0
		for {
			if isDigit(l.peek) {
				value = value*10 + int(l.peek-'0')
				l.peek = l.nextInputChar()
			} else {
				// set peek to space
				l.peek = space

				// return num token
				return NewNum(value)
			}
		}
	}

	return nil
}

func (l *Lexer) readLexeme() *word {
	if isChar(l.peek) {
		// new buffer to build string -> might want to use string.Builder for efficiency.
		buf := bytes.Buffer{}
		buf.Grow(20) // initial size 20 bytes

		// build lexeme until we hit non char | digit value
		for {
			// first value of peek guaranteed to be char. Subsequent values of lexeme can be char or digit
			if isChar(l.peek) || isDigit(l.peek) {
				// add to buf
				err := buf.WriteByte(l.peek)
				if err != nil {
					fmt.Println("buffer too large.", err)

				}
				l.peek = l.nextInputChar()
			} else {
				break
			}
		}

		// check if lexeme in buf already exists in words
		lexeme := buf.String()
		storedWord := l.words[lexeme]
		if storedWord != nil {
			// if exists return word
			return &word{
				Tag:    TagId,
				Lexeme: storedWord.value(),
			}
		} else {
			// add lexeme to words
			newWord := NewWord(lexeme)
			l.words[lexeme] = newWord

			// set peek to space
			l.peek = space

			// return newly added word
			return newWord
		}
	}

	return nil
}

func (l *Lexer) readCharacter() *char {
	// assume this is a character
	char := NewChar(l.peek)

	// set peek to space
	l.peek = space

	// return
	return char
}

func isDigit(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func isChar(char byte) bool {
	if char >= 65 && char <= 90 || char >= 97 && char <= 122 {
		return true
	}
	return false
}
