/*
    simple postfix solver implementation for basic arithmetic operations
    {'+', '-', '*', '/'}

    run: go run ex2.8.go
*/

package main

import (
	"fmt"
	"strconv"
)

const (
	testString = "11+2-"
)

func main() {
	fmt.Println(postfixSolve(testString))
}

func postfixSolve(str string) int {
	var index = 0 

	fmt.Printf("string: %s\n", str)

	for (len(str) != 1) {
		fmt.Printf("string len: %v\n", len(str))
		fmt.Printf("string index: %v\n", index)

		c := string(str[index])
		switch char := c; char {
			case "+":
				sum := charToInt(str[index-2]) + charToInt(str[index-1])
				str = strconv.Itoa(sum) + str[index+1:]

				fmt.Printf("sum: %v\n", sum)
				fmt.Printf("new str: %v\n", str)
				index = 0
			case "-":
				diff := charToInt(str[index-2]) - charToInt(str[index-1])
				str = string(diff) + str[index+1:]

				fmt.Printf("diff: %v\n", diff)
				fmt.Printf("new str: %v\n", str)
				index = 0 
			case "*":
				prod := charToInt(str[index-2]) * charToInt(str[index-1])
				str = string(prod) + str[index+1:]

				fmt.Printf("prod: %v\n", prod)
				fmt.Printf("new str: %v\n", str)
				index = 0 
			case "/":
				quot := charToInt(str[index-2]) / charToInt(str[index-1])
				str = string(quot) + str[index+1:]

				fmt.Printf("quot: %v\n", quot)
				fmt.Printf("new str: %v\n", str)
				index = 0
			default:
				index++	
		}
	}
	fmt.Printf("final: %s\n", str)
	answer, _ := strconv.Atoi(str)
	return answer
}

func charToInt(b byte) int {
	integer, _ := strconv.Atoi(string(b))
	return integer
}
