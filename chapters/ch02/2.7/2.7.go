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

func NewTranslator(lexer *Lexer) *Translator {
	t := &Translator{
		lexer: lexer,
	}

	t.lookahead = lexer.Scan()
	return t
}

func (t *Translator) program() {
	t.top = nil
	t.block()
}

func (t *Translator) block() {
	t.matchCharacter('{')

	saved := t.top
	t.top = NewEnv(t.top)
	fmt.Print("{ ")

	t.decl()

	t.stmts()

	t.matchCharacter('}')

	t.top = saved
	fmt.Print("} ")
}

func (t *Translator) decl() {
	// match type
	if t.lookahead.tag() == tagBool || t.lookahead.tag() == tagInt || t.lookahead.tag() == tagChar {
		// store type
		s := NewSymbol(t.lookahead)

		// get id
		t.lookahead = t.lexer.Scan()

		// store type against id
		t.top.put(t.lookahead.value(), s)

		// match id
		t.matchId()

	} else {
		return
	}
	t.matchCharacter(';')

	t.decl()
}

func (t *Translator) stmts() {
	if string(t.lookahead.value()) == "}" {
		return
	}

	t.stmt()

	t.stmts()
}

func (t *Translator) stmt() {
	// block
	if string(t.lookahead.value()) == "{" {
		t.block()
	} else {
		// factor
		t.id()

		t.matchCharacter(';')
	}
}

func (t *Translator) id() {
	s := t.top.get(t.lookahead.value())

	fmt.Print(t.lookahead.value())
	fmt.Print(":")
	fmt.Print(s.sType)
	fmt.Print("; ")

	t.lookahead = t.lexer.Scan()
}

func (t *Translator) matchId() {
	if t.lookahead.tag() == tagId {
		t.lookahead = t.lexer.Scan()

		if t.lookahead == nil {
			return
		}
	} else {
		fmt.Println("expected lexeme")
	}
}

func (t *Translator) matchCharacter(char byte) {
	if t.lookahead.tag() == tagTerm && string(char) == t.lookahead.value() {
		t.lookahead = t.lexer.Scan()

		if t.lookahead == nil {
			return
		}
	} else {
		fmt.Println("syntax error matching character")
	}
}
