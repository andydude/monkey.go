# monkey.go
Implementation of Monkey in Go (using ANTLR4)

## The book

The book is [Writing an Interpreter in Go by Thorsten Ball](https://interpreterbook.com/).

## Differences

There are differences between this repo and the book. This variant has the parser written in ANTLR4, which is then used to automatically generate an `LL` parser in Go, roughly with the command:

- `antlr4 -Dlanguage=Go ...`

You can find shell scripts in the `monkey/parser` directory, along with the important files:

- `monkey/parser/MonkeyLexer.g4` this replaces the old `lexer.go`
- `monkey/parser/MonkeyParser.g4` this replaces the old `parser.go`
- `monkey/parser/visitor.go` this takes the automatically generated parse tree, and converts it into the object model defined in `ast.go`. This allows the alternate parser to pass all of the test cases in the book.

## ANTLR4 Rant

I have since learned that there is a blog post called [Parsing with ANTLR 4 and Go](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/), which is a gentle introduction to using the `antlr4-go` runtime. However, having used ANTLR4 on other projects, I noticed that there is a tendancy to use Listeners instead of Visitors in almost every example in the documentation. The code linked to in the article above includes hundreds of examples of Listeners implemented in Go, *but it includes zero examples of how to use Visitors*!

The **Listener** model simply does not support the process of AST generation. Firstly, Listeners assume that you want to implement custom behavior before and after, and so provide `enterX` and `exitX` methods. Secondly, you have to implement a stack, or member method of pushing and popping, and figure out how to return values in a way that doesn't return values, and just doesn't map well to the process of generating an AST from a string of tokens. The **Visitor** model, on the other hand, does support AST generation. Visitors assume that you're going to return a value that becomes the input of visitors higher up in the AST tree. This maps very well to AST generation, and so that's why I decided to use Visitors in this implementation.

Even the official documentation for ANTLR4 in other languages tends to document Listeners, and leave Visitors as an exercise for the reader. Hopefully, this repo will be one giant example of how to use ANTLR4 Visitors in Go.
