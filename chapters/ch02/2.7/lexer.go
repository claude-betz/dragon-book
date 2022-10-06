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

	// terminals
	tagId   = 256
	bracket = 257

	// reserved key-words
	tagInt  = 258
	tagChar = 259
	tagBool = 260
	tagTerm = 261
)

type Token interface {
	tag() int
	value() string
}

type char struct {
	Tag   int
	Value byte
}

func NewChar(character byte) *char {
	return &char{
		Tag:   tagTerm,
		Value: character,
	}
}

func (c char) tag() int {
	return c.Tag
}

func (c char) value() string {
	return string(c.Value)
}

type word struct {
	Tag    int
	Lexeme string
}

func newWord(str string) *word {
	return &word{
		Tag:    tagId,
		Lexeme: str,
	}
}

func (w *word) tag() int {
	return w.Tag
}

func (w *word) value() string {
	return w.Lexeme
}

var (
	complete = false
)

type Lexer struct {
	LineNum int
	buffer  *bytes.Buffer
	peek    byte
	words   map[string]Token
}

func (l *Lexer) nextInputChar() byte {
	c, err := l.buffer.ReadByte()
	if err != nil {
		// fmt.Println("end of buffer")
		complete = true
	}

	return c
}

func NewLexer(buffer *bytes.Buffer) *Lexer {
	// initialise lexer
	words := make(map[string]Token)
	lexer := &Lexer{
		LineNum: 0,
		buffer:  buffer,
		peek:    space,
		words:   words,
	}

	// populate map with reserved words
	lexer.reserve(
		&word{
			Tag:    tagBool,
			Lexeme: "bool",
		},
	)

	lexer.reserve(
		&word{
			Tag:    tagChar,
			Lexeme: "char",
		},
	)

	lexer.reserve(
		&word{
			Tag:    tagInt,
			Lexeme: "int",
		},
	)

	return lexer
}

func (l *Lexer) reserve(w *word) {
	l.words[w.Lexeme] = w
}

func (l *Lexer) Scan() Token {
	if complete == false {
		// skip white-space
		l.skipWhiteSpace()

		// read word
		word := l.readLexeme()
		if word != nil {
			return word
		}

		// read character
		return l.readCharacter()
	}

	return nil
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

func (l *Lexer) readLexeme() Token {
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

		// unread byte we read in loop
		l.buffer.UnreadByte()

		// check if lexeme in buf already exists in words
		lexeme := buf.String()
		storedWord := l.words[lexeme]
		if storedWord != nil {
			// if exists return word
			return storedWord
		} else {
			// add lexeme to words
			newWord := newWord(lexeme)
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
