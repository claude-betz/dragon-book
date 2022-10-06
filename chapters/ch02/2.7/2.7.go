/*
	Syntax directed translator

	from:
	{ int x; char y; { bool y; x; y; } x; y; }

	to:
	{ { x:int; y:bool; } x:int; y:char; }
*/

package main

import (
	"fmt"
)

type Symbol struct {
	sType string
}

func NewSymbol(t Token) *Symbol {
	return &Symbol{
		sType: t.value(),
	}
}

type Env struct {
	table map[string]*Symbol
	pEnv  *Env
}

func NewEnv(pEnv *Env) *Env {
	t := make(map[string]*Symbol)
	return &Env{
		table: t,
		pEnv:  pEnv,
	}
}

func (e *Env) put(lexeme string, symbol *Symbol) {
	e.table[lexeme] = symbol
}

func (e *Env) get(lexeme string) *Symbol {
	currEnv := e
	for {
		// search for declaration in tables from inner nesting outwards
		symbol := currEnv.table[lexeme]
		if symbol != nil {
			return symbol
		}

		// if not found in any symbol table, break out
		if currEnv.pEnv == nil {
			break
		}

		currEnv = currEnv.pEnv
	}

	return nil
}

type Translator struct {
	lookahead        Token
	lookaheadPointer int
	complete         bool
	top              *Env
	lexer            *Lexer
}

func (t *Translator) program() {
	t.top = nil
	t.block()
}

func (t *Translator) block() {
	t.matchCharacter('{')

	saved := t.top
	t.top = NewEnv(t.top)
	fmt.Printf("{")

	t.decl()

	t.stmt()

	t.matchCharacter('}')

	t.top = saved
	fmt.Printf("}")
}

func (t *Translator) decl() {
	// match type
	if t.lookahead.tag() == tagBool|tagInt|tagChar {
		// match id
		t.matchId()
	} else {
		fmt.Println("expected type")
	}
	t.matchCharacter(';')

	s := NewSymbol(t.lookahead)
	t.top.put(t.lookahead.value(), s)

	t.decl()
}

func (t *Translator) matchId() {
	if t.lookahead.tag() == tagId {
		// do nothing
	} else {
		fmt.Println("expected lexeme")
	}
}

func (t *Translator) stmt() {
	token := t.lookahead

	// block
	if token.value() == "{" {
		t.block()
	}

	t.matchId()

	s := t.top.get(token.value())
	print(token.value())
	print(":")
	print(s.sType)
}

func (t *Translator) matchCharacter(char byte) {
	if t.lookahead.tag() == tagChar && string(char) == t.lookahead.value() {
		t.lookaheadPointer++
	} else {
		fmt.Println("syntax error matching character")
	}
}
