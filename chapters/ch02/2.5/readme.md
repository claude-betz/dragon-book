# simple syntax directed translator

implementation of a recursive-descent directed translator that translates infix to postfix.
The infix expressions are restricted to sequences of digits separated by '+' and '-'
symbols.

## postfix expressions

postfix expressions can be recursively defined as:

1. if $E$ is a constant
$$ postfix(E) = E $$
2. if $E$ is an expression of the form: $E_1 op E_2$
$$ postfix(E) = postfix(E_1) postfix(E_2) op $$
3. if $E$ is an expression of the form: $(E)$
$$ P(E) = postfix(E) $$

## context-free grammar

expr -> expr + term
    | expr - term
    | term

term -> 0
    | 1
    | 2
    ...
    | 9

## syntax directed translation scheme

actions translating infix to postfix

```md
expr -> expr + term {print('+')}
    |   expr - term {print('-')}
    |   term

term -> 0
    |   1 {print('1')}
    |   2 {print('2')}
    ...
    |   9 {print('3')}
```

note: we don't concern ourselves with brackets since the grammar doesn't allow these.

## eliminating left recursion

```md
A -> Aα | Aβ | γ

can be re-written as:

A -> γR
R -> αR | βR | ε
```

### syntax directed translation scheme without left-recursion

```md
A = expr
α = + term {print('+')}
β = - term {print('-')}
γ = term
```

```md
expr -> term rest
rest -> + term {print('+')} rest
    |   - term {print('-')} rest
    |   ε

term -> 0
    |   1 {print('1')}
    |   2 {print('2')}
    ...
    |   9 {print('3')}

```
