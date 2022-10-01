/*
    Construct recursive-descent parsers starting with the following grammars
    a) S -> +SS | -SS | a

*/

package main

import (
    "fmt"
)

const (
    validString = "++aa-aa"
    invalidString = "++aa-aa+++"
)

var (
    str string
    lookahead byte
    complete = false
    lookaheadPointer = 0
)

func S() {
    if (complete) {
        return
    } else {
        fmt.Printf("current lookahead: %v\n", lookahead)
        switch (lookahead) {
            case ('+'):
                fmt.Println("matched '+'")

                match('+')
                S()
                S()
            case ('-'):
                fmt.Println("matched '-'")

                match('-')
                S()
                S()
            case ('a'):
                fmt.Println("matched 'a'")

                match('a')
            default:
                fmt.Println("syntax error!")
        }
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
    // invalid string example
    str = invalidString
    lookahead = str[0]
    fmt.Printf("invalid string: %s len: %v\n", str, len(str))

    // parse string
    S()
    checkParsingOutcome() 

    resetGlobalVariables()

    // valid string example
    str = validString
    lookahead = str[0]
    fmt.Printf("valid string: %s len: %v\n", str, len(str))

    // parse string
    S()
    checkParsingOutcome() 
}

func resetGlobalVariables() {
    complete = false
    lookaheadPointer = 0
}

func checkParsingOutcome() {
    // result
    if complete {
        fmt.Println("successfully parsed string")
    } else {
        fmt.Println("failed to parse string")
    }
}


