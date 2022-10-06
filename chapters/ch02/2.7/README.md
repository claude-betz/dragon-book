## Symbol Tables
Symbol tables are data structures used by compilers to hold information about source-program constructs. The information is collected incrementally during the analysis phases of a compiler and used by the synthesis phases to generate the target code.

Entries in the symbol table contain information about an identifier such as its character string (lexeme), it's type, it's position in storage. 
They typically need to support multiple declarations of the same identifier within a program.

## Implementing a translator that relies on a symbol table
Using a dumbed down language in which a program consists of blocks with optional declarations and "statements" consisting of single identifiers.

e.g.
```
{ int x; char y; { bool y; x; y; } x; y; }
```

The task we wish to perform is to print a revised program, in which the declarations have been removed and each "statement" has it's identifier followed by a colon and it's type.

i.e for the above example the translation looks like:
```
{ { x:int; y:bool; } x:int; y:char; }
```

## Symbol table per scope
The "scope of identifier x" refers to the scope of a particular declaration of x. The term scope by itself refers to a portion of the program that is the scope of one or more declarations.

Scopers are important, because the same identifier can be declared for different purposes in different parts of a program. Common names like i and x often have multiple uses. Also in languages that support inheritance, subclasses can redeclare a method name to overridee a method in a superclass.

If blocks can be nested, several declarations of the same identifier can appear within a single block. The following syntax results in nested blocks when `stmts` can generate a block

```
block -> '{' decls stmts '}' 
```

The most closely nested rule for blocks is that an identifier x is in the scope of the most-closely nested declaration of x. 
That is, the declaration of x found by examining blocks inside out, starting with the block in which x appears.

## translation scheme for program

```
program ->  block           {   
                                top = null;
                            }

block ->    '{'             {
                                saved = top; 
                                top = new Env(top);
                                print('{');
                            } 

            decls stmts '}' {
                                top = saved;
                                print('}');
                            }

decls   ->  decls decl
        |   ϵ

decl    ->  type id;        {
                                s = new Symbol();
                                s.type = type.lexeme
                                top.put(id.lexeme, s);
                            }

stmts   ->  stmts stmt
        |   ϵ

stmt    ->  block
        |   factor;         {   
                                print(';')
                            }

factor  ->  id              {
                                s = top.get(id.lexeme);
                                print(id.lexeme);
                                print(':');
                                print(s.type)  
                            }
```

Notice that the non-terminals `decls` and `stmts` are left-recursive which we already know can cause infinite loops when implementing a recursive-descent top down parser.

To remedy this, we rewrite the above translation scheme to eliminate using what we learned in section 2.5

```
decls   ->  ϵ       R
R       ->  decl    R
        |   ϵ
```

```
stmts   ->  ϵ       P
P       ->  stmt    P
        |   ϵ
```

We can simplify by substituting (ignoring the epsilon terminal)
- R into decls
- P into stmts

```
decls   ->  decl    R
        |   ϵ
```

```
stmts   ->  stmt    P
        |   ϵ
```

The translation scheme we will implement in code becomes the following:
```
program ->  block           {   
                                top = null;
                            }

block ->    '{'             {
                                saved = top; 
                                top = new Env(top);
                                print('{');
                            } 

            decls stmts '}' {
                                top = saved;
                                print('}');
                            }

decls   ->  decl    R       {
        |   ϵ                   s = new Symbol();
                                s.type = type.lexeme
                                top.put(id.lexeme, s);
                            }

decl    ->  type id;        {
                                s = new Symbol();
                                s.type = type.lexeme
                                top.put(id.lexeme, s);
                            }

stmts   ->  stmt    P
        |   ϵ

stmt    ->  block
        |   factor;         {   
                                print(';')
                            }

factor  ->  id              {
                                s = top.get(id.lexeme);
                                print(id.lexeme);
                                print(':');
                                print(s.type)  
                            }
```