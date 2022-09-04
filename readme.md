Compiler course on edx from Stanford and "Writing interpreters in Go" by Thorsten Ball

# languages
* compilers - takes input (program) - produces executable. This is another program, that takes data and produces output. Offline
* interpreters - takes input (program) and data and produces direct output. Online

# compilation process

### lexical analysis
recognize words - syntax of language. Divide program text into words - `tokens`

`if x==y then z=1; else z=2;`
tokens:
* if
* then
* else
* ;
* z=1, z=2
* x==y
* spaces

### parsing
understand words. Identify role of token in the text. Group in higher level constructs. Building a tree of `if-then-else` statement:
* x==y - predicate
* == - relation
* = - assignment

### semantic analysis
understanding meaning of high level structure. **This is hard!** Compilers are only catching inconsistencies, we are very limited. We don't know what the program is meant to do. e.g.

* strict rules to avoid ambiguity - variable shadowing
* type checking

### optimization
auto modification of program so that they run faster/use less memory.
e.g. for integers we can optimize:

```
x=y*0
x=0
```
### code generation
producing an exe (c->machine code) or a program in other language (js->ts, c->asm)

# Lexical Analysis (LA)
tokenize text input (classify program substring) and communicate tokens to parser. Done by regular expressions (implementation - finite automaton)

Output: list of tokens - pairs with `class` and `string corresponding`(lexeme)

* recognise substrings corresponding to tokens (lexemes)
* identify token class of each lexeme

classes: identifiers, keywords, `(`, `)`, numbers, strings, whitespace (grouped, 3 spaces are 1 token - whitespace) etc.

example of token - LA output:
* <ID, "foo"> - identifier
* <OP, "="> - assignment

```
if (i==j) z = 0;
else z = 1;
```
find whitespace, keywords, identifiers, numbers, equal operator, (, ), ;, =

Fortran - whitespaces are insignificant. Lexer removes all whitespaces and it still works. Sometimes it requires `lookaheads` - peek to next chars to classify the lexeme. Other examples from regular languages:
* single char scan (vars `i`, `j`) 
* `else` - `els` is valid identifier

PL1 - keywords are not reserved. We can have vars `if` etc. Lookahead is 
needed to get the context. Lexer is harder to write

## finite automata
on state s1, on input a go to state s2. If end state reached - accept. Otherwise - reject (end state in not finite or acceptable state or stuck).

each move consumes input (moves ptr in string). There are also epsilon moves - does not consume inputs.

deterministic finite automata. DFA
* one transition per input
* no e-moves

non deterministic finite automate - NFA
* can have multiple transition points per 1 input
* can have e-moves

implementation (DFA) - 2d array
* rows - states
* columns - inputs
* cells - next states

in case of sparse tables - graph like - map of vector of moves for each possible state. allows to share rows (by setting ptrs)


implementation (NFA) - 2d array
* rows - states
* columns - inputs. Additional column - epsilon
* cells - set of next states

each iteration needs to process in a loop - for each el in next state set

# Parsing

building parsed tree (syntax tree) from list of tokens.
Parser distinguishes valid/invalid list of tokens. Often recursive structure.

`Context-free grammars` (CFG) are a natural notation for that structure.
* set o terminal symbols, 
* set of non terminal symbols - these evaluate to another, like statements, expressions etc., 
* start symbol 
* set of productions (mapping one symbol to multiple symbols)

derivation - sequence of productions. can be drawn as a tree

Abstract syntax tree - parse tree without some details.

BNF (Backus-Naur Form) notation for defining CFG - `<expression> ::= <term> [<addop> <term>]*`

alternative to BNF and CFGs - `pratt's parsing`

nodes
* expression (right side of assignment)
* binary operations (+,-,*,/)
* statements - list of expressions
* assignments
* if-else
* if 
* predicate (boolean expression)
* loop
* function

expression - produces a value, statement does not.

Program is just a series of statements

expressions:
```
5
5+1
true
add(1,3)
x // might be also x identifier in declaration:  var x
func(x,y) { return x+y; }
```

statements:
```
var x = 5;
return 5;
var add = func(x,y) { return x+y; }
```

### Parser Implementation

storing a tree:
* as a literal tree with sum types as nodes (rust enums, haskell datatypes, jsons with loosely defined nodes)
* as oop interfaces (evalExpression and evalStatemnets, both void methods)

# Evaluation

walking through the AST and interpreting each tree node