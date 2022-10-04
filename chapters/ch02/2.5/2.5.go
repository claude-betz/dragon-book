/*
	Simple syntax directed translator

	Golang implementation of a syntax directed translator that translates infix 
	to postfix.
*/

package main

import (
	"fmt"
)

const (
	validString = "1+2-3"
)

var (
	str string
	lookahead byte
	lookaheadPointer = 0
	complete = false
)

func expr() {
	matchTerm()
	rest() 
}

func rest() {
	if (complete) {
        return
    } else {
		switch (lookahead) {
			case ('+'):
				match('+')
				matchTerm()
				fmt.Printf("+")
				rest()
			case ('-'):
				match('-')
				matchTerm()
				fmt.Printf("-")
				rest()
			default:
				fmt.Println("[rest] syntax error")
		}
	}
}

func matchTerm() {
	switch (lookahead) {
		case ('0'):
			match('0')
			fmt.Printf("0")
		case ('1'):
			match('1')
			fmt.Printf("1")
		case ('2'):
			match('2')
			fmt.Printf("2")
		case ('3'):
			match('3')
			fmt.Printf("3")
		case ('4'):
			match('4')
			fmt.Printf("4")
		case ('5'):
			match('5')
			fmt.Printf("5")
		case ('6'):
			match('6')
			fmt.Printf("6")
		case ('7'):
			match('7')
			fmt.Printf("7")
		case ('8'):
			match('8')
			fmt.Printf("8")
		case ('9'):
			match('9')
			fmt.Printf("9")
		default:
			fmt.Println("syntax error.")
	}
}

func match(terminal byte) {
	if terminal == lookahead {
		lookaheadPointer++
		if lookaheadPointer < len(str) {
			lookahead = str[lookaheadPointer]
		} else {
            // termination condition
            complete = true            
        }
	} else {
        fmt.Println("syntax error")
    }
}

func main() {
	str = validString
	lookahead = str[0]

	expr()
}
