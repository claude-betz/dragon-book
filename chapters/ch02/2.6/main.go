package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "num + val1 + 20"
	buf := bytes.NewBufferString(str)

	// create Lexer
	lex := NewLexer(buf)

	token := lex.Scan()

	for i := 1; i <= 5; i++ {
		fmt.Printf("tag: %d value: %s\n", token.tag(), token.value())
	}
}
