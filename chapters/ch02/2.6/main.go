package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "num + 50000 + val1 + 20"
	buf := bytes.NewBufferString(str)

	// create Lexer
	lex := NewLexer(buf)

	for i := 1; i <= 7; i++ {
		token := lex.Scan()
		fmt.Printf("tag: %d value: %s\n", token.tag(), token.value())
	}
}
