package main

import (
	"bytes"
	"fmt"
)

/*
	{ int x; char y; { bool y; x; y; } x; y; } -> { { x:int; y:bool; } x:int; y:char; }
*/

func main() {
	str := "{ int x; char y; { bool y; x; y; } x; y; }"
	fmt.Println(str)

	buf := bytes.NewBufferString(str)
	lexer := NewLexer(buf)

	for i := 1; i < 22; i++ {
		token := lexer.Scan()
		fmt.Printf("tag: %d value: %s\n", token.tag(), token.value())
	}
}
