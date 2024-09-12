package main

import (
    "fmt"
    "kisumu/lexer" 
)

func main() {
    input := `func main() {
        display("Hello, World!");
        x := 5 + 10;
    }`

    l := lexer.NewLexer(input)
    for {
        tok := l.NextToken()
        fmt.Printf("Token Type: %d, Literal: %q\n", tok.Type, tok.Literal)

        if tok.Type == lexer.EOF {
            break
        }
    }
    // fmt.Println(l)
}