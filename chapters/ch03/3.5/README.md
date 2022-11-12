# The Lexical Analyzer Generator Lex

This is a tool that allows one to specify a Lexical Analyzer by specifying regular expressions
to describe patterns for tokens.

The input notation is referred to as the `Lex Language`
The tool itself is called the `Lex Compiler`

Behind the scenes, the Lex Compiler transforms the input patterns into a transition diagram and generates
code in a file called `lex.yy.c`, that simluare this transition diagram.

# Running the exercises

1. Compiling from Lex to C
```
$ flex <filename>.l
```
this will write a `lex.yy.c` file to the current directory

2. Creating executable from C
```
$ cc lex.yy.c
```
this will write an `a.out` executable to the current directory

3. Running against test case
```
$ a.out <test-filename>.txt
```
this will output a file called `<test-filename>-result.txt` to the current directory







