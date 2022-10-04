## Lexical Analysis
Continuing with conversion from infix to postfix notation.

We want to implement a lexical analyser that allows numbers, identifiers and white-space (blanks, tabs, newlines)
to appear in expressions.
We will extend our grammar to allow addition, subtraction, division, multiplication and supports numbers and identifiers
e.g. var1 + 400/var2 - var3*600

### context free grammar
```md
In order of increasing precedence:
    left-associative: '+' '-'
    left-associative: '*' '/'

    NOTE:
        non-terminal: term -> for precedence (1)
        non-terminal: expr -> for precedence (2)
        non-terminal: factor -> can not be torn apart by any operator

    expr -> expr + term
        |   expr - term
        |   term

    term -> term / factor
        |   term * factor
        |   factor

    factor -> num | id
```

### translation scheme

```md
    expr -> expr + term     { print('+') }
        |   expr - term     { print('-') }
        |   term           

    term -> term / factor   { print('/') }
        |   term * factor   { print('*') }
        |   factor

    factor -> (expr)
        |   num             { print(num.value) }
        |   id              { print(id.lexeme) }
```

### Removing white spaces
This is achieved by reading characters as long as we see a blank, tab or newline character.

### Recognising numbers
If the peek variable is a digit we continue reading and building up the number until we stop reading a digit.

### Recognising lexemes
If the peek is a character, we read until we read a something that isn't a character. Note this case is a bit more complex 
because we need to distinguish between keywords (e.g. true) and indetifiers (e.g. var1, var2).

We use a string table to hold character string. This solves two problems:
1. Single representation: insulate the rest of the compiler from the respresentation of strings, the other phases of the compiler can word with references to the string table.

2. Reserved words: These can be implemented by initialising the string table with reserved strings and their tokens. When the lexixal analyser reads a string or lexeme that could form an identifier, it first checks if the lexeme is in the string table. If so it returns the token from the table, otherwise, it returns a token with terminal id.

### Token scan
1. Skip white space.
2. handle numbers.
3. handle reserved words and indentifiers.
4. if code reaches here tread peek as token.
5. set peek to blank and return.